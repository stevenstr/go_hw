package postgresql

import (
	"database/sql"
	"fmt"
	"strings"
	"sync"

	"github.com/stevenstr/go_hw/pkg/model"
)

type PostgersContactsRepository struct {
	datab *sql.DB
}

func PostgresNewContactsRepository(db *sql.DB) *PostgersContactsRepository {
	return &PostgersContactsRepository{
		datab: db,
	}
}

func (r *PostgersContactsRepository) Save(contact model.Contact) (model.Contact, error) {
	
	query := "INSERT INTO contacts(firstname,lastname,phone,email) VALUES($1,$2,$3,$4) returning id;"
	err := r.datab.QueryRow(query, contact.FirstName, contact.LastName, contact.Phone, contact.Email).Scan(&contact.ID)

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
func (r *PostgersContactsRepository) ListAll() (contacts []model.Contact, error) {
	
	query := "SELECT id, firstname, lastname, email, phone FROM contacts;"
	
	rows, err := r.datab.Query(query)
	
	if err != nil {
		return contacts, err
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
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (r *PostgersContactsRepository) GetByID(id uint) (contact model.Contact, error) {
	
	query := "SELECT id, firstname, lastname, email, phone FROM contacts WHERE id = $1;"
	
	row := r.datab.QueryRow(query, id)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)

	if err != nil {
		if err != sql.ErrNoRows {
			return model.Contact{}, err
		}

		return model.Contact{}, fmt.Errorf("record not found")
	}
	return contact, nil
}

func (r *PostgersContactsRepository) GetByPhone(phone string) (contact model.Contact, error) {
	
	query := "SELECT id, firstname, lastname, email, phone FROM contacts WHERE phone = $1;"
	
	row := r.datab.QueryRow(query, phone)
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)

	if err != nil {
		if err != sql.ErrNoRows {
			return model.Contact{}, err
		}
		return model.Contact{}, fmt.Errorf("record not found")
	}

	return contact, nil
}

func (r *PostgersContactsRepository) GetByEmail(email string) (conatact model.Contact, error) {
	
	query := "SELECT id, firstname, lastname, email, phone FROM contacts WHERE email = $1;"
	
	row := r.datab.QueryRow(query, email)
	
	err = row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Email, &contact.Phone)

	if err != nil {
		if err != sql.ErrNoRows {
			return model.Contact{}, err
		}
		return model.Contact{}, fmt.Errorf("record not found")
	}

	return contact, nil
}

func (r *PostgersContactsRepository) SearchByName(n string) (contacts []model.Contact, error) {
	query := "SELECT id, firstname, lastname, email, phone FROM contacts WHERE firstname = $1;"
	rows, err := r.datab.Query(query, n)
	if err != nil {
		return contacts, err
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
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *PostgersContactsRepository) Delete(id uint) error {
	
	query := "DELETE FROM contacts WHERE id = $1;"
	
	_, err := r.databb.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
