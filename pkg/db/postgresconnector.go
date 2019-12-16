package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//DBConnector func
func DBConnector(connector string) (*sql.DB, error) {

	db, err := sql.Open("postgres", connector)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")
	return db, nil
}

