package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

func indexFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>STRONA GŁÓWNA</body></html>")
}

func itemFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body>STRONA ITEM<br>ADRES: ")
	fmt.Fprintf(w, r.RequestURI)
	fmt.Fprintf(w, "<br>METODA: ")
	fmt.Fprintf(w, r.Method)
	fmt.Fprintf(w, "</body></html>")
}

type Student struct {
	Imie     string `json:"imie"`
	Nazwisko string `json:"nazwisko"`
	Indeks   int    `json:"indeks"`
	Email    string `json:"-"`
}

func daneFunc(w http.ResponseWriter, r *http.Request) {

	var studenci []Student = []Student{
		{"Jan", "Kowalski", 12345, "test@test"},
		{"Marek", "Nowak", 30000, "to@tamto"},
		{"Anna", "Zdyb", 23232, "anna@zdyb"},
	}
	tmpl, _ := template.ParseFiles("pages/dane.html")
	tmpl.Execute(w, studenci)
}

func JSONFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // nagłówek JSON
	var studenci []Student = []Student{
		{Imie: "Jan", Nazwisko: "Kowalski", Indeks: 12345, Email: "test@test"},
		{Imie: "Marek", Nazwisko: "Nowak", Indeks: 30000, Email: "to@tamto"},
		{Imie: "Anna", Nazwisko: "Zdyb", Indeks: 23232, Email: "anna@zdyb"},
	}
	data, _ := json.Marshal(studenci)
	w.Write(data)
}

func Zad3Func(w http.ResponseWriter, r *http.Request) {
	los := rand.Intn(3)
	switch los {
	case 0:
		itemFunc(w, r)
	case 1:
		daneFunc(w, r)
	case 2:
		JSONFunc(w, r)
	}
}

func main() {
	http.HandleFunc("/", indexFunc)     // wzorzec /*
	http.HandleFunc("/item/", itemFunc) // wzorzec /item/*
	http.HandleFunc("/strona/", daneFunc)
	http.HandleFunc("/json/", JSONFunc)
	http.HandleFunc("/zad3/", Zad3Func)
	http.ListenAndServe("localhost:8080", nil)
}
