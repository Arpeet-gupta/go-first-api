package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Arpeet-gupta/go-first-api/v2/database"
	"github.com/gorilla/mux"
)

func addItem(w http.ResponseWriter, r *http.Request) {
	err := database.Createtable(db)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Printf("Error %s when creating posts table", err)
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, "INSERT INTO posts(title, body, author) VALUES(?, ?, ?)")
	if err != nil {
		io.WriteString(w, err.Error())
		log.Printf("Error %s when preparing SQL statement", err)
	}
	defer stmt.Close()

	newItem := Post{}
	err = json.NewDecoder(r.Body).Decode(&newItem)

	if err != nil {
		io.WriteString(w, err.Error())
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", newItem)

	res, err := stmt.ExecContext(ctx, newItem.Title, newItem.Body, newItem.Author)
	if err != nil {
		io.WriteString(w, err.Error())
		log.Printf("Error %s when inserting row into posts table", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		io.WriteString(w, err.Error())
		log.Printf("Error %s when finding rows affected", err)
	}
	log.Printf("%d posts created ", rows)

	res.LastInsertId()

	// w.Header().Set("Content-Type", "")
	io.WriteString(w, "New Post was created")
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

func patchPost(w http.ResponseWriter, r *http.Request) {
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

	patchPost := &posts[id]
	json.NewDecoder(r.Body).Decode(patchPost)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(patchPost)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
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

	posts = append(posts[:id], posts[id+1:]...)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(posts)
}
