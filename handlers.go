package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Arpeet-gupta/go-first-api/v3/database"
	"github.com/Arpeet-gupta/go-first-api/v3/migrations"
	"github.com/Arpeet-gupta/go-first-api/v3/models"
	"github.com/gorilla/mux"
)

func addAuthor(w http.ResponseWriter, r *http.Request) {
	newAuthor := &models.Author{}

	err := migrations.Createtable()

	if err != nil {
		log.Printf("Error when creating table: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}

	err = json.NewDecoder(r.Body).Decode(newAuthor)

	if err != nil {
		log.Printf("Error when decoding json to struct: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}

	response := database.Db.Debug().Create(newAuthor)

	if response.Error != nil {
		log.Printf("Error when inserting row into posts table: %s", response.Error)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, response.Error.Error())
	}

	log.Printf("%d Author created ", response.RowsAffected)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "New Author was created")
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "ID could not converted to integer")
	}
	updatedAuthor := &models.Author{}

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("Error: %s", r)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, err)
		}
	}()
	res := database.Db.Where("id = ?", id).Preload("AllBook").Find(&updatedAuthor)

	if res.Error != nil {
		log.Printf("Error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}

	if err = json.NewDecoder(r.Body).Decode(updatedAuthor); err != nil {
		log.Printf("Error when decoding json to struct: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}
	res = database.Db.Save(updatedAuthor)
	if res.Error != nil {
		log.Printf("Error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
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
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "ID could not converted to integer")
		log.Printf("ID %s could not converted to integer", idParam)
	}

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("Error: %s", r)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, err)
		}
	}()
	res := database.Db.Where("id = ?", id).Preload("AllBook").Find(&author)

	if res.Error != nil {
		log.Printf("Error: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&author)
	if err != nil {
		log.Printf("Error while Encoding response from struct %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	deleteauthor := &models.Author{}
	idParam := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "ID could not converted to integer")
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
		log.Printf("Error: %s", res.Error.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, res.Error.Error())
	}

	log.Printf("%d Author Deletd ", res.RowsAffected)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&allauthors)
	if err != nil {
		log.Printf("Error while Encoding response from struct %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
	}
}
