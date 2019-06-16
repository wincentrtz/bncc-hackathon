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

func createHotelTable() {
	db := config.InitDb()
	defer db.Close()
	schema := `CREATE TABLE hotel(
			id serial PRIMARY KEY,
			name VARCHAR NOT NULL,
			phone VARCHAR NOT NULL,
			address VARCHAR NOT NULL,
			city VARCHAR NOT NULL,
			price INT NOT NULL,
			created_on TIMESTAMP NOT NULL
		);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createFlightTable() {
	db := config.InitDb()
	defer db.Close()
	schema := `CREATE TABLE flight(
			id serial PRIMARY KEY,
			date TIMESTAMP,
			time VARCHAR,
			departure VARCHAR,
			destination VARCHAR,
			price INT,
			created_on TIMESTAMP NOT NULL
		);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createAlbumTable() {
	db := config.InitDb()
	defer db.Close()
	schema := `CREATE TABLE albums(
			id serial PRIMARY KEY,
			name VARCHAR,
			description VARCHAR,
			paid INTEGER,
			created_on TIMESTAMP NOT NULL
		);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	createUserTable()
	createHotelTable()
	createFlightTable()
	createAlbumTable()
}
