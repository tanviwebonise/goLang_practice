package main 

import (
	"io"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong!")
}

func contact(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Check contacts here")
}


func main() {
	mux:= http.NewServeMux()
	mux.HandleFunc("/", ping)
	mux.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", mux)
}