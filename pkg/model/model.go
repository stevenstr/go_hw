package model

import "fmt"

//ContactsRepository type
type ContactsRepository interface {
	Save(Contact) (Contact, error)
	ListAll() ([]Contact, error)
	GetByID(uint) (Contact, error)
	GetByPhone(string) (Contact, error)
	GetByEmail(string) (Contact, error)
	SearchByName(string) ([]Contact, error)
	Delete(uint) error
}

//Contact type
type Contact struct {
	ID uint `json:"id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	Phone string `json:"phone"`
	Email string `json:"email"`
}

//String method
func (c Contact) String() string {
	return fmt.Sprintf("Contact:\n\tID - %d\n\tFirst Name - %q\n\tLast Name - %q\n\tPhone - %q\n\tEmail - %q\n", c.ID, c.FirstName, c.LastName, c.Phone, c.Email)
}
