package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	connectionString := "user=postgres dbname=loja_alura_go password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err.Error())
	}

	return db
}
