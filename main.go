package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Arpeet-gupta/go-first-api/database"
	"github.com/gorilla/mux"
)

//User is struct to hold JSON's objects
// type User struct {
// 	FullName string `json:"fullname"`
// 	UserName string `json:"username"`
// 	Email    string `json:"email"`
// }

// Post is  struct to hold JSON body of Request
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

var posts []Post
var db *sql.DB
var err error

func main() {
	router := mux.NewRouter()
	db, err = database.CreateDataConnectionPool("ecommerce")
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
		log.Fatal(err.Error())
	}
	defer db.Close()

	log.Printf("Successfully connected to database")

	router.HandleFunc("/posts", addItem).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
