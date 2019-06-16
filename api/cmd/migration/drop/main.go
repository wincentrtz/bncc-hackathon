package main

import (
	"log"

	"github.com/wincentrtz/bncc/api/config"
)

func main() {
	db := config.InitDb()
	defer db.Close()
	schema := `DROP TABLE 
		users,
		hotel,
		flight,
		albums`
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}
