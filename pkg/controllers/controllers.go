package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Rexkizzy22/bookstore/pkg/models"
	"github.com/Rexkizzy22/bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}	

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error parsing bookId: ")
	}
	bookDetails, _ := models.GetBookByID(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	NewBook := &models.Book{}
	utils.ParseBody(r, NewBook)
	b := NewBook.CreateBook()
	res, err := json.Marshal(b)
	if err != nil {
		log.Fatal("Failed to marshal book")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error parsing bookId: ")
	}
	bookDetails := models.DeleteBook(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdatedBook := &models.Book{}
	utils.ParseBody(r, UpdatedBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error parsing bookId: ")
	}

	bookDetails, _ := models.GetBookByID(ID)
	if bookDetails.Name != "" {
		bookDetails.Name = UpdatedBook.Name
	}

	if bookDetails.Author != "" {
		bookDetails.Author = UpdatedBook.Author
	}

	if bookDetails.Publication != "" {
		bookDetails.Publication = UpdatedBook.Publication
	}

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}