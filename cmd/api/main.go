package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Rexkizzy22/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/Rexkizzy22/bookstore/pkg/config"
	"github.com/Rexkizzy22/bookstore/pkg/models"
)

func init() {
	config.Connect()
	db := config.GetDB()
	db.AutoMigrate(&models.Book{})

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := mux.NewRouter()
	routes.RegisterBookRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server listening on http://localhost:3001")
	log.Fatal(http.ListenAndServe(":3001", r))
}