FROM golang:latest
ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

# Copy local vendor file
COPY vendor ./vendor

# Copy migrations
COPY db/migrations ./migrations

COPY . .

RUN go build -mod=vendor -o main ./api/api.go
CMD ["./main", "172.18.0.4"]
