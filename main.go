package main

import (
	"alura-loja-go/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()

	http.ListenAndServe(":8000", nil)
}
