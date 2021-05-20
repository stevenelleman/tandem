FROM golang:latest
ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

# Copy local vendor file
COPY vendor ./vendor

# Copy migrations
COPY migrations ./migrations

COPY . .

RUN go build -mod=vendor -o main ./api/api.go
CMD ["./main", "public-api", "api-store"]
