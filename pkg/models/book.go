package models

import (
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func Init(database *gorm.DB) {
	if database == nil {
		panic("models : db is nil")
	}
	db = database
	db.AutoMigrate((&Book{}))
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	var getBookByID Book
	result := db.Where("ID=?", Id).Find(&getBookByID)
	return &getBookByID, result
}

func DeleteBook(Id int64) Book {
	var books Book
	db.Where("ID=?", Id).Delete(books)
	return books
}

// we dont write query ORM (gorm) takes care of it
