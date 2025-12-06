package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ihtgoot/CRUD_API_DB/pkg/models"
	"github.com/ihtgoot/CRUD_API_DB/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error parsing book Id")
	}
	bookDetails, _ := models.GetBookByID(Id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookID"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("err while parseing")
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookID"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("err while parseing")
	}

	bookDetails, db := models.GetBookByID(Id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Render(w http.ResponseWriter, html string) {
	parseTemplate, err := template.ParseFiles("cmd/main/index.html")
	if err != nil {
		fmt.Println("Error parsing templates:", err)
		http.Error(w, "Internal Server Error 1 ", http.StatusInternalServerError)
		return
	}
	err = parseTemplate.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		fmt.Println("error executing template : ", err)
		http.Error(w, "Internal Server Error 2 ", http.StatusInternalServerError)
		return
	}
}

func Frontend(w http.ResponseWriter, r *http.Request) {
	Render(w, "index.html")
}
