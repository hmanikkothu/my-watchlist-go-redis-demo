# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.12 as builder

# Copy local code to the container image.
WORKDIR /

COPY . .
COPY static static

RUN go get github.com/go-redis/redis

# Build the command inside the container.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /main /opt/main
COPY --from=builder static /static

# Run the web service on container startup.
CMD ["/opt/main"]

#docker network create -d bridge hkb
#docker run --rm --network=hkb --name=watchlistredis gcr.io/google_containers/redis:e2e
#docker run --rm --network=hkb -it -p=8080:8080 --name=my-watchlist --env REDIS_HOST=watchlistredis localhost/my-watchlist sh
#docker run --rm --network=hkb -p=8080:8080 --name=my-watchlist --env REDIS_HOST=watchlistredis localhost/my-watchlist