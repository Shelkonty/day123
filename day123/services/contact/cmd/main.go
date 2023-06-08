package main

import (
	"Projects/day/pkg/store/postgres"
	"database/sql"
	"fmt"
)

func main() {
	db, err := postgres.OpenDB("localhost", "", "postgres", "1234", "gop")
	if err != nil {
		fmt.Printf("Fail")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
}
