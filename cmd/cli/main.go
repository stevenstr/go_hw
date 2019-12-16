package main

import (
	"fmt"
	"log"

	db "github.com/stevenstr/go_hw/pkg/db"
	"github.com/stevenstr/go_hw/pkg/model"
	repo "github.com/stevenstr/go_hw/pkg/repository/postgresql"
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

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := db.DBConnectDB(*psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	repository := repo.NewContactsRepositoryPostgreSQL(conn)

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
			log.Printf("command not foumd for value %d\n", command)
		}

		printSeparator()
	}
}

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

func Delete(rep model.ContactsRepository) error {
	id := readUint("Please enter an 'ID' field and press Enter")

	if err := rep.Delete(id); err != nil {
		return fmt.Errorf("error in GetByID: %q", err.Error())
	}

	fmt.Printf("Delete:\nRecord with ID %d successfylly deleted\n", id)
	return nil
}

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

func printSeparator() {
	for i := 0; i < 50; i++ {
		fmt.Print("*")
	}

	fmt.Println()
}
