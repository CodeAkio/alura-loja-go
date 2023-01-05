package main

import (
	"alura-loja-go/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.LoadRoutes()

	fmt.Printf("🚀 Running at http://localhost:8000\n")
	http.ListenAndServe(":8000", nil)
}
