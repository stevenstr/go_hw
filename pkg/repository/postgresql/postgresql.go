/*
 *Author: Stefan
 *Date: 12/18/2019
 *Last changes: 12/18/2019 13.55
 *Task: Replace in memory implementation with implementation using
 *		mongoDB or PostgreSQL
 */

package postgresql

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/stevenstr/pkg/model"
)

//ContactsRepositoryPostgreSQL type
type ContactsRepositoryPostgreSQL struct {
	db *sql.DB
}

//NewContactsRepositoryPostgreSQL function
func NewContactsRepositoryPostgreSQL(db *sql.DB) *ContactsRepositoryPostgreSQL {
	return &ContactsRepositoryPostgreSQL{
		db: db,
	}
}

//Save method
func (r *ContactsRepositoryPostgreSQL) Save(contact model.Contact) (model.Contact, error) {
	query := "INSERT INTO contsct(first_name,last_name,phone,email) VALUES($1,$2,$3,$4) returning id;"

	err := r.db.QueryRow(query, contact.FirstName, contact.LastName, contact.Phone, contact.Email).Scan(&contact.ID)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"contacts_phone_idx\"") {
			return model.Contact{}, fmt.Errorf("contact with phone %q already exists", contact.Phone)
		}
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"contacts_email_idx\"") {
			return model.Contact{}, fmt.Errorf("contact with email %q already exists", contact.Email)
		}
		return model.Contact{}, err
	}
	return contact, nil
}

//ListAll method
func (r *ContactsRepositoryPostgreSQL) ListAll() (users []model.Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM contsct;"
	rows, err := r.db.Query(query)

	if err != nil {
		return users, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var contact model.Contact
		err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
		if err != nil {
			return nil, err
		}
		users = append(users, contact)
	}
	return users, nil
}

//GetByID method
func (r *ContactsRepositoryPostgreSQL) GetByID(id uint) (contact model.Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM contsct WHERE id = $1;"
	row := r.db.QueryRow(query, id)

	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)

	if err != nil {
		if err != sql.ErrNoRows {
			return model.Contact{}, err
		}
		return model.Contact{}, fmt.Errorf("record not found")
	}
	return contact, nil
}

//GetByPhone method
func (r *ContactsRepositoryPostgreSQL) GetByPhone(phone string) (contact model.Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM contsct WHERE phone = $1;"
	row := r.db.QueryRow(query, phone)

	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)

	if err != nil {
		if err != sql.ErrNoRows {
			return model.Contact{}, err
		}

		return model.Contact{}, fmt.Errorf("record not found")
	}
	return contact, nil
}

//GetByEmail method
func (r *ContactsRepositoryPostgreSQL) GetByEmail(email string) (contact model.Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM contsct WHERE email = $1;"
	row := r.db.QueryRow(query, email)

	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)

	if err == sql.ErrNoRows {
		return model.Contact{}, fmt.Errorf("record not found")
	} else if err != nil {
		return model.Contact{}, err
	}
	return contact, nil
}

//SearchByName method
func (r *ContactsRepositoryPostgreSQL) SearchByName(n string) (conresult []model.Contact, err error) {
	query := "SELECT id, first_name, last_name, email, phone FROM contsct WHERE first_name = $1;"
	rows, err := r.db.Query(query, n)

	if err != nil {
		return conresult, err
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var contact model.Contact
		err = rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)
		if err != nil {
			return nil, err
		}
		conresult = append(conresult, contact)
	}
	return conresult, nil
}

//Delete method
func (r *ContactsRepositoryPostgreSQL) Delete(id uint) error {
	query := "DELETE FROM contsct WHERE id = $1;"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
