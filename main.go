package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//User is struct to hold JSON's objects
type User struct {
	FullName string `json:"fullname"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

// Post is  struct to hold JSON body of Request
type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author User   `json:"author"`
}

var posts []Post

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", addItem).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", patchPost).Methods("PATCH")
	http.ListenAndServe(":8080", router)
}
