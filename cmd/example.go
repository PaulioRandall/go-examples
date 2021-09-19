package main

import (
	"fmt"
	"log"

	"github.com/PaulioRandall/go-examples/server"
)

func main() {
	fmt.Println("Starting server on port 8080")
	h := server.CreateHandler()
	log.Fatal(h.ListenAndServe(":8080"))
}
