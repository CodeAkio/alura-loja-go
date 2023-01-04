package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)

	fmt.Printf("🚀 Running at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"Camiseta", "Camiseta azul de algodão", 39, 5},
		{"Tenis", "Confortável", 89, 3},
		{"Fone", "Alta definição", 59, 2},
	}

	temp.ExecuteTemplate(w, "Index", products)
}
