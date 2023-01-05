package controllers

import (
	"alura-loja-go/models"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()
	temp.ExecuteTemplate(w, "Index", products)
}
