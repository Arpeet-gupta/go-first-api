package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func addItem(w http.ResponseWriter, r *http.Request) {
	newItem := &Post{}
	json.NewDecoder(r.Body).Decode(newItem)
	fmt.Printf("%#v", newItem)

	posts = append(posts, *newItem)
	w.Header().Set("Content-Type", "application.json")
	json.NewEncoder(w).Encode(posts)
}
