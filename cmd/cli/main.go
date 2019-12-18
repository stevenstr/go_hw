/*
 *Author: Stefan
 *Date: 12/18/2019
 *Last changes: 12/18/2019 13.55
 *Task: Replace in memory implementation with implementation using
 *		mongoDB or PostgreSQL
 */

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/stevenstr/pkg/model"
	pos "github.com/stevenstr/pkg/repository/postgresql"

	//driver
	_ "github.com/lib/pq"
)

const (
	CommandSave = iota + 1
	CommandListAll
	CommandGetByID
	CommandGetByPhone
	CommandGetByEmail
	CommandSearchByName
	CommandDelete
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "postgres"
)

func main() {

	//connector to db using manual : https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	fmt.Println("Successfully connected!")

	//
	repository := pos.NewContactsRepositoryPostgreSQL(db)

	for {
		fmt.Print(menu)
		var command int
		if _, err := fmt.Scanf("%d", &command); err != nil {
			log.Println(err)
		}

		switch command {

		case CommandSave:
			if err := Save(repository); err != nil {
				log.Println(err)
			}
		case CommandListAll:
			if err := ListAll(repository); err != nil {
				log.Println(err)
			}
		case CommandGetByID:
			if err := GetByID(repository); err != nil {
				log.Println(err)
			}
		case CommandGetByPhone:
			if err := GetByPhone(repository); err != nil {
				log.Println(err)
			}
		case CommandGetByEmail:
			if err := GetByEmail(repository); err != nil {
				log.Println(err)
			}
		case CommandSearchByName:
			if err := SearchByName(repository); err != nil {
				log.Println(err)
			}
		case CommandDelete:
			if err := Delete(repository); err != nil {
				log.Println(err)
			}
		default:
			log.Printf("HZ %d\n", command)
		}

		printSeparator()
	}
}

//ListAll function
func ListAll(rep model.ContactsRepository) error {
	records, err := rep.ListAll()
	if err != nil {
		return fmt.Errorf("error in ListAll: %q", err.Error())
	}

	fmt.Println("ListAll:")
	for _, r := range records {
		fmt.Println(r)
	}

	return nil
}

//GetByID function
func GetByID(rep model.ContactsRepository) error {
	id := readUint("Please enter an 'ID' field and press Enter")

	record, err := rep.GetByID(id)
	if err != nil {
		return fmt.Errorf("error in GetByID: %q", err.Error())
	}

	fmt.Println("GetByID")
	fmt.Println(record)

	return nil
}

//GetByPhone function
func GetByPhone(rep model.ContactsRepository) error {
	phone := readString("Please enter an 'Phone' field and press Enter")

	record, err := rep.GetByPhone(phone)
	if err != nil {
		return fmt.Errorf("error in GetByPhone: %q", err.Error())
	}

	fmt.Println("GetByPhone:")
	fmt.Println(record)

	return nil
}

//GetByEmail function
func GetByEmail(rep model.ContactsRepository) error {
	email := readString("Please enter an 'Email' field and press Enter")

	record, err := rep.GetByEmail(email)
	if err != nil {
		return fmt.Errorf("error in GetByEmail: %q", err.Error())
	}

	fmt.Println("GetByEmail:")
	fmt.Println(record)

	return nil
}

//SearchByName function
func SearchByName(rep model.ContactsRepository) error {
	email := readString("Please enter prefix for 'Name' field and press Enter")

	records, err := rep.SearchByName(email)
	if err != nil {
		return fmt.Errorf("error in SearchByName: %q", err.Error())
	}

	fmt.Println("SearchByName:")
	for _, r := range records {
		fmt.Println(r)
	}

	return nil
}

//Delete function
func Delete(rep model.ContactsRepository) error {
	id := readUint("Please enter an 'ID' field and press Enter")

	if err := rep.Delete(id); err != nil {
		return fmt.Errorf("error in GetByID: %q", err.Error())
	}

	fmt.Printf("Delete:\nRecord with ID %d successfylly deleted\n", id)
	return nil
}

//Save function
func Save(rep model.ContactsRepository) error {
	contact := model.Contact{
		FirstName: readString("Please enter an 'FirstName' field and press Enter"),
		LastName:  readString("Please enter an 'LastName' field and press Enter"),

		Phone: readString("Please enter an 'Phone' field and press Enter"),
		Email: readString("Please enter an 'Email' field and press Enter"),
	}

	result, err := rep.Save(contact)
	if err != nil {
		return err
	}

	fmt.Println("Save Contact:")
	fmt.Println(result)

	return nil
}

//Just simple menu
const menu = `
Please enter operation number:
  * 1 - Save
  * 2 - ListAll
  * 3 - GetByID
  * 4 - GetByPhone
  * 5 - GetByEmail
  * 6 - SearchByName
  * 7 - Delete 
  * Control + C - to exit 
`

//readStriing function
func readString(message string) string {
	var r string

	for r == "" {
		fmt.Println(message)
		if _, err := fmt.Scanf("%s", &r); err != nil {
			fmt.Printf("Error in process of reading string from console\n\t%q\n please try again\n", err.Error())
			printSeparator()
		}

	}
	return r
}

//readInt function
func readUint(message string) uint {
	var r uint

	for {
		fmt.Println(message)
		_, err := fmt.Scanf("%d", &r)
		if err == nil {
			break
		}

		fmt.Printf("Error in process of reading string from console\n\t%q\n please try again\n", err.Error())
		printSeparator()
	}
	return r
}

//printSeparator function
func printSeparator() {
	for i := 0; i < 50; i++ {
		fmt.Print("*")
	}

	fmt.Println()
}
