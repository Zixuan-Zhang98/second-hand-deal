package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Welcome to the index page!")
	if err != nil {
		log.Fatalln(err)
	}
}
