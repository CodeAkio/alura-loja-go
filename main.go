package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)

	fmt.Printf("🚀 Running at http://localhost:8000\n")
	http.ListenAndServe(":8000", nil)
}

func connectDb() *sql.DB {
	connectionString := "user=postgres dbname=loja_alura_go password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectDb()

	selectAllProducts, err := db.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectAllProducts.Scan(&id, &name, &description, &quantity, &price)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Quantity = quantity
		product.Price = price

		products = append(products, product)
	}

	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
