package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nayonacademy/golang-oauth2/api/handlers"
	"log"
)
var server = handlers.Server{}
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize()
	server.Run(":8000")

}
