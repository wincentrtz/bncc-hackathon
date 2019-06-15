package main

import (
	"log"

	"github.com/wincentrtz/bncc/config"
)

func createPostTable() {
	db := config.InitDb()
	defer db.Close()
	schema := `CREATE TABLE posts(
			id serial PRIMARY KEY,
			post_parent_id integer,
			post_title VARCHAR,
			post_description VARCHAR,
			created_on TIMESTAMP NOT NULL
		);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createPostStatusTable() {
	db := config.InitDb()
	defer db.Close()
	schema := `CREATE TABLE post_status(
			id serial PRIMARY KEY,
			post_id integer NOT NULL,
			status integer NOT NULL,
			created_on TIMESTAMP NOT NULL,
			FOREIGN KEY (post_id) REFERENCES posts (id)
		);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createUserTable() {
	db := config.InitDb()
	defer db.Close()
	schema := `CREATE TABLE users(
			id serial PRIMARY KEY,
			name VARCHAR NOT NULL,
			email VARCHAR NOT NULL,
			password VARCHAR NOT NULL,
			created_on TIMESTAMP NOT NULL
		);`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	createPostTable()
	createPostStatusTable()
	createUserTable()
}

func insertData() {
	db := config.InitDb()
	defer db.Close()
	schema := `
		INSERT INTO posts
		VALUES 
			(DEFAULT, 1, 'title', 'Hello World', NOW() - INTERVAL 10 DAY),
			(DEFAULT, 2, 'title', 'Nice weather today!', NOW() - INTERVAL 6 DAY),
			(DEFAULT, 1, 'title', 'Wow, I'm in demo :D', NOW() - INTERVAL 6 DAY),
			(DEFAULT, 4, 'title', 'Great app', NOW() - INTERVAL 2 DAY),
			(DEFAULT, 5, 'title', 'Today is Monday', NOW() - INTERVAL 1 DAY),
			(DEFAULT, 6, 'title', 'Wahawhawhawhah funny', NOW() - INTERVAL 1 DAY),
			(DEFAULT, 7, 'title', 'Ulalal', NOW()),
			(DEFAULT, 5, 'title', 'Im hungry', NOW()),
			(DEFAULT, 8, 'title', 'What should I eat', NOW()),
			(DEFAULT, 1, 'title', 'I shared Hello World, check it out', NOW()),
			(DEFAULT, 2, 'title', 'I shared Nice Weather today!, check it out', NOW()),
			(DEFAULT, 1, 'title', 'I shared Hello World, check it out', NOW());
	`

	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}
