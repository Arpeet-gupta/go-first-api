package main

import (
	"io"
	"log"
	"net/http"
)

func errorHandle(w http.ResponseWriter, message string, err error, status int) {
	log.Printf("%s: %s", message, err)
	w.WriteHeader(status)
	io.WriteString(w, err.Error())
}
