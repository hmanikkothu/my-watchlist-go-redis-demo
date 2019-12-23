# build frontend

### build
```
docker build localhost/watchlist .
```

### run locally
```
docker run -p 8080:8080 localhost/watchlist
```

### run frontend and redis on same docker network
```
docker network create -d bridge hkb
docker run --rm --network=hkb --name=redis-watchlist gcr.io/google_containers/redis:e2e
docker run --rm --network=hkb -p=8080:8080 --name=my-watchlist --env REDIS_HOST=redis-watchlist localhost/my-watchlist
```

### push image to the registry
```
docker build -t <repo>/watchlist:v1 .
docker push <repo>/watchlist:v1
```