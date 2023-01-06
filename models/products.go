package models

import (
	"alura-loja-go/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func FindAllProducts() []Product {
	db := db.ConnectDb()

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

	defer db.Close()

	return products
}

func CreateProduct(name, description string, price float64, quantity int) {
	db := db.ConnectDb()

	insertDataOnDb, err := db.Prepare("INSERT INTO products(name, description, price, quantity) VALUES($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	_, err = insertDataOnDb.Exec(name, description, price, quantity)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
