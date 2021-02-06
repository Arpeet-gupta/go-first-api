package main

import (
	"log"
	"net/http"

	"github.com/Arpeet-gupta/go-first-api/v3/database"
	"github.com/gorilla/mux"
)

// //User is struct to hold JSON's objects
// type User struct {
// 	FullName string `json:"fullname"`
// 	UserName string `json:"username"`
// 	Email    string `json:"email"`
// }

// // Post is  struct to hold JSON body of Request
// type Post struct {
// 	Title  string `json:"title"`
// 	Body   string `json:"body"`
// 	Author User   `json:"author"`
// }

// var posts []Post

func main() {
	router := mux.NewRouter()
	err := database.CreateDataConnectionPool("ecommerce")
	if err != nil {
		log.Printf("Error when getting database connection pool: %s", err)
	}
	defer database.Db.Close()
	log.Printf("Successfully connected to database")

	router.HandleFunc("/posts", addAuthor).Methods("POST")
	router.HandleFunc("/posts", getAllAuthors).Methods("GET")
	router.HandleFunc("/posts/{id}", getAuthor).Methods("GET")
	router.HandleFunc("/posts/{id}", updateAuthor).Methods("PUT")
	router.HandleFunc("/posts/{id}", deleteAuthor).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
