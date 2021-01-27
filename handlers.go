package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func addItem(w http.ResponseWriter, r *http.Request) {
	newItem := &Post{}
	json.NewDecoder(r.Body).Decode(newItem)
	fmt.Printf("%#v", newItem)

	posts = append(posts, *newItem)
	w.Header().Set("Content-Type", "application.json")
	json.NewEncoder(w).Encode(posts)
}

func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "ID could not converted to integer")
		return
	}

	if id > len(posts) {
		w.WriteHeader(404)
		io.WriteString(w, "No post found with the specific id id")
		return
	}
	post := posts[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "ID could not converted to integer")
		return
	}

	if id > len(posts) {
		w.WriteHeader(404)
		io.WriteString(w, "No post found with the specific id id")
		return
	}

	updatedPost := &Post{}

	json.NewDecoder(r.Body).Decode(updatedPost)

	posts[id] = *updatedPost

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(updatedPost)
}
