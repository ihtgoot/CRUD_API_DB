package routs

import (
	"github.com/gorilla/mux"
	"github.com/ihtgoot/CRUD_API_DB/pkg/controller"
)

var RegisterBookStoreRouts = func(router *mux.Router) {
	router.HandleFunc("/", controller.Frontend).Methods("GET")
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookID}", controller.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookID}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookID}", controller.DeleteBook).Methods("DELETE")
}

// GO IDIOMS :
// RegisterBookStoreRouts is a variable it stores and ananymous function
// Same behavior. Same performance. Same visibility
// This pattern is useful when one or more of these are true:
//		You want to replace / override the function (testing, dependency injection)
// 		You want to pass it around like data
//		Package-level dependency injection
// This is exactly what you want for:
//		route registration
//		helpers
//		utilities
//		most application code
// if normal function its immutability by default but this is mutable , i.e,it can be repace in runtime

// It’s like a pluggable function pointer, not like class method overloading.

//Why Go allows this (and Java discourages it) :
// 		Go prefers:
//			composition over inheritance
//			functions over objects
//			explicit wiring
