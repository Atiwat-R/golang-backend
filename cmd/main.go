package main

import (
	"log"

	"github.com/atiwat-r/golang-backend/cmd/api"
)





func main() {
	server := api.NewAPIServer(":8080", nil)
	err := server.Run()

	if (err != nil) {
		log.Fatal(err)
	}
}