/*
 *Author: Stefan
 *Date: 12/25/2019
 *Last changes: 01/08/2019
 *Task: Gorm is a top^), DB connection and automigration our struct to DB
 */

package models

//import
import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

//db var
var db *gorm.DB

//init function
func init() {

	//Load out .env file to load configs
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	//Connector
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	//Migration
	db = conn
	db.Debug().AutoMigrate(&Account{})
}

//GetDB function
func GetDB() *gorm.DB {
	return db
}
