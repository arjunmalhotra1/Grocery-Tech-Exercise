package db

import (
	"database/sql"
	"fmt"
)

func NewDb() *sql.DB {
	db, err := sql.Open("MySQL", "")
	if err != nil {
		fmt.Println("Error opening the db connection")
	}
	return db

}
