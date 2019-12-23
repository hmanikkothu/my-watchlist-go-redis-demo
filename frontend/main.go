package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
)

var (
	rdb *redis.Client
)

type item struct {
	ID    string
	Value string
}

const wastlist string = "watchlist"
const httpServerError int = 500

func main() {
	host := os.Getenv("REDIS_HOST")

	rdb = redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: "",
		DB:       0,
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/item", itemHandler)
	http.HandleFunc("/allitems", getAllHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getAllHandler(w http.ResponseWriter, r *http.Request) {
	data, err := rdb.HGetAll(wastlist).Result()
	if err != nil {
		serverError(w, r, fmt.Sprintf("error %s!", err))
	}
	log.Printf("getAllHandler: retrieved items")
	l := len(data)
	items := make([]item, l, l)
	var c int
	for k, v := range data {
		i := item{
			ID:    k,
			Value: v,
		}
		items[c] = i
		c++
	}
	jsonString, _ := json.Marshal(items)
	fmt.Fprintf(w, string(jsonString))
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	var i item
	method := r.Method

	var err error
	if method == "DELETE" {
		id := r.FormValue("id")
		i = item{ID: id}
	} else {
		err = json.NewDecoder(r.Body).Decode(&i)
	}
	if err == nil {
		if method == "GET" {
			getHandler(w, r, i)
		} else if method == "PUT" {
			putHandler(w, r, i)
		} else if method == "DELETE" {
			deleteHandler(w, r, i)
		} else {
			errorHandler(w, r, "method not supported", 405)
		}
	} else {
		serverError(w, r, fmt.Sprintf("json.NewDecoder error: %s!", err))
	}
}

func getHandler(w http.ResponseWriter, r *http.Request, i item) {
	val, err := rdb.HGet(wastlist, i.ID).Result()
	if err == nil || err == redis.Nil {
		i.Value = val
		writeItem(w, r, i)
		log.Printf("got item %s", i)
	} else {
		serverError(w, r, fmt.Sprintf("error: %s!", err))
	}
}

func putHandler(w http.ResponseWriter, r *http.Request, i item) {
	err := rdb.HSet(wastlist, i.ID, i.Value).Err()
	if err != nil {
		serverError(w, r, fmt.Sprintf("error %s!", err))
	}
	log.Printf("put item = %s", i)
	writeItem(w, r, i)
}

func deleteHandler(w http.ResponseWriter, r *http.Request, i item) {
	err := rdb.HDel(wastlist, i.ID).Err()
	if err != nil {
		serverError(w, r, fmt.Sprintf("error %s!", err))
	}
	log.Printf("deleted item = %s", i)
	writeItem(w, r, i)
}

func serverError(w http.ResponseWriter, r *http.Request, msg string) {
	errorHandler(w, r, msg, httpServerError)
}

func errorHandler(w http.ResponseWriter, r *http.Request, msg string, code int) {
	http.Error(w, msg, code)
}

func writeItem(w http.ResponseWriter, r *http.Request, i item) {
	b, err := json.Marshal(i)
	si := string(b)
	//log.Printf("item = %s", si)
	if err != nil {
		msg := fmt.Sprintf("error %s!", err)
		serverError(w, r, msg)
	} else {
		fmt.Fprintf(w, si)
	}
}
