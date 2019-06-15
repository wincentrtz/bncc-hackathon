package main

import (
	"log"

	"github.com/wincentrtz/bncc/api/config"
)

func createUserTable() {
	db := config.InitDb()
	defer db.Close()
	schema := `CREATE TABLE users(
			id serial PRIMARY KEY,
			name VARCHAR NOT NULL,
			email VARCHAR NOT NULL,
			password VARCHAR NOT NULL,
			phone VARCHAR NOT NULL,
			address VARCHAR NOT NULL,
			gender VARCHAR NOT NULL,
			ktp VARCHAR NOT NULL,
			created_on TIMESTAMP NOT NULL
		);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	createUserTable()
}
