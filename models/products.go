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

	selectAllProducts, err := db.Query("SELECT * FROM products ORDER BY id ASC")

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

func DeleteProduct(id string) {
	db := db.ConnectDb()

	deleteProduct, err := db.Prepare("DELETE FROM products WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}

	_, err = deleteProduct.Exec(id)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectDb()

	selectedProduct, err := db.Query("SELECT * FROM products WHERE id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}

	for selectedProduct.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectedProduct.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity
	}

	defer db.Close()

	return productToUpdate
}

func UpdateProduct(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDb()

	updateProduct, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}

	_, err = updateProduct.Exec(name, description, price, quantity, id)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
