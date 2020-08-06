package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/nayonacademy/golang-oauth2/api/models"
	"log"
	"net/http"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize() {

	var err error
	server.DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Printf("Cannot connect to database\n")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the database\n")
	}
	server.DB.Exec("PRAGMA foreign_keys = ON")

	server.DB.Debug().AutoMigrate(&models.User{}) //database migration
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}