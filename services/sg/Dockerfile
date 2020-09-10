FROM golang:latest
ENV GO111MODULE=on

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .

# Copy local vendor file
COPY vendor ./vendor
COPY . .

RUN go build -mod=vendor -o main main.go
CMD ["./main"]