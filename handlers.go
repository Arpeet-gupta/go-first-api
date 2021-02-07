package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Arpeet-gupta/go-first-api/v4/database"
	"github.com/Arpeet-gupta/go-first-api/v4/migrations"
	"github.com/Arpeet-gupta/go-first-api/v4/models"
	"github.com/gorilla/mux"
)

func addAuthor(w http.ResponseWriter, r *http.Request) {
	newAuthor := &models.Author{}

	if err := migrations.Createtable(); err != nil {
		errorHandle(w, "Error when creating table:", err, http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(newAuthor); err != nil {
		errorHandle(w, "Error when decoding json to struct:", err, http.StatusBadRequest)
		return
	}

	response := database.Db.Debug().Create(newAuthor)

	if response.Error != nil {
		errorHandle(w, "Error when inserting row into posts table:", response.Error, http.StatusBadRequest)
		return
	}

	log.Printf("%d Author created ", response.RowsAffected)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "New Author was created")
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	updatedAuthor := &models.Author{}

	if err != nil {
		errorHandle(w, "ID could not converted to integer:", err, http.StatusBadRequest)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("Error: %s", r)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, err)
		}
	}()
	// res := database.Db.Where("id = ?", id).Preload("AllBook").Find(&updatedAuthor)
	res := database.Db.Preload("AllBook").Find(&updatedAuthor, id)

	if res.Error != nil {
		errorHandle(w, "Error:", res.Error, http.StatusBadRequest)
		return
	}

	if err = json.NewDecoder(r.Body).Decode(updatedAuthor); err != nil {
		errorHandle(w, "Error when decoding json to struct:", err, http.StatusBadRequest)
		return
	}

	res = database.Db.Save(updatedAuthor)
	if res.Error != nil {
		errorHandle(w, "Error:", res.Error, http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	log.Printf("%d Author Updated ", res.RowsAffected)
	io.WriteString(w, "Author Updated")
}

func getAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorHandle(w, "ID could not converted to integer:", err, http.StatusBadRequest)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("Error: %s", r)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, err)
		}
	}()
	res := database.Db.Preload("AllBook").Find(&author, id)

	if res.Error != nil {
		errorHandle(w, "Error:", res.Error, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&author)
	if err != nil {
		errorHandle(w, "Error while Encoding response from struct:", err, http.StatusBadRequest)
		return
	}
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	deleteauthor := &models.Author{}
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorHandle(w, "ID could not converted to integer:", err, http.StatusBadRequest)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("Error: %s", r)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, err)
		}
	}()
	res := database.Db.Unscoped().Select("AllBook").Delete(deleteauthor, id)
	if res.Error != nil {
		log.Printf("Error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}
	log.Printf("%d Author Deleted ", res.RowsAffected)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Author Deleted")
}

func getAllAuthors(w http.ResponseWriter, r *http.Request) {
	var allauthors []models.Author

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("Error: %s", r)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, err)
		}
	}()
	res := database.Db.Preload("AllBook").Find(&allauthors)

	if res.Error != nil {
		errorHandle(w, "Error:", res.Error, http.StatusBadRequest)
		return
	}

	log.Printf("Author Count: %d ", res.RowsAffected)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&allauthors)
	if err != nil {
		errorHandle(w, "Error while Encoding response from struct:", err, http.StatusBadRequest)
		return
	}
}
