package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nayonacademy/golang-oauth2/handlers"
)

func main() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":8000"),
		Handler: handlers.New(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server Closed")
	}
}
