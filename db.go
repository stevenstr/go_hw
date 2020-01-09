/*
 *Author: Stefan
 *Date: 12/18/2019
 *Last changes: 12/18/2019 13.55
 *Task: Querry for db
*/

package main

import (
	"database/sql"
)

//insert finction
func insert(message, accountID string) (sql.Result, error) {
	return db.Exec("INSERT INTO tweets VALUES (default, $1, $2)",
		message, accountID)
}

//readOne function
func readOne(id int) (Tweet, error) {
	var rec Tweet
	row := db.QueryRow("SELECT * FROM tweets WHERE id=$1 ORDER BY id", id)
	return rec, row.Scan(&rec.Id, &rec.Message, &rec.AID)
}

//read function
func read(str string) ([]Tweet, error) {
	var rows *sql.Rows
	var err error
	if str != "" {
		rows, err = db.Query("SELECT * FROM tweets WHERE name LIKE $1 ORDER BY id",
			"%"+str+"%")
	} else {
		rows, err = db.Query("SELECT * FROM tweets ORDER BY id")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rs = make([]Tweet, 0)
	var rec Tweet
	for rows.Next() {
		if err = rows.Scan(&rec.Id, &rec.Message, &rec.Message); err != nil {
			return nil, err
		}
		rs = append(rs, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rs, nil
}
