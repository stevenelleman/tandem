package main

import (
	"os"
	"sg/services/public-api/internal/api"
)

func main() {
	host := os.Args[1]
	if host == "" {
		panic("Host not defined")
	}
	store := os.Args[2]
	if store == "" {
		panic("Store not defined")
	}
	api.Run(store)
}