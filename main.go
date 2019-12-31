package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	Buyer = iota + 1
	Seller
	Admin
)

type User struct {
	Id int

	Name  string
	First string
	Last  string

	Password []byte

	// Enum type
	// 1 stands for buyer, 2 stands for seller, 3 stands for admin
	Role []int

	IsBuyer  bool
	IsSeller bool
}

type Good struct {
	Id          int
	Name        string
	Category    string
	Owner       User
	Description string
	Price       float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

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
