package main

import (
	"log"
	"net/http"

	"github.com/Arpeet-gupta/go-first-api/v4/database"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	err := database.CreateDataConnectionPool("ecommerce")
	if err != nil {
		log.Fatal(err)
	}
	sqlDb, err := database.Db.DB()
	log.Printf("Successfully connected to database")
	defer sqlDb.Close()

	router.HandleFunc("/posts", addAuthor).Methods("POST")
	router.HandleFunc("/posts", getAllAuthors).Methods("GET")
	router.HandleFunc("/posts/{id}", getAuthor).Methods("GET")
	router.HandleFunc("/posts/{id}", updateAuthor).Methods("PUT")
	router.HandleFunc("/posts/{id}", deleteAuthor).Methods("DELETE")
	http.ListenAndServe(":8080", router)
}
