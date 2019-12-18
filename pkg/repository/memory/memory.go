package memory

import (
	"fmt"
	"strings"
	"sync"

	"github.com/stevenstr/pkg/model"
)

//ContactsRepositoryInMemory  type
type ContactsRepositoryInMemory struct {
	sync.RWMutex
	storage map[uint]model.Contact
	lastID  uint
}

//NewContactsRepositoryInMemory function
func NewContactsRepositoryInMemory() *ContactsRepositoryInMemory {
	return &ContactsRepositoryInMemory{
		storage: make(map[uint]model.Contact),
	}
}

//Save method
func (r *ContactsRepositoryInMemory) Save(contact model.Contact) (model.Contact, error) {
	r.Lock()
	defer r.Unlock()

	for _, c := range r.storage {
		if c.Email == contact.Email {
			return model.Contact{}, fmt.Errorf("contact with email %q already exists", c.Email)
		}

		if c.Phone == contact.Phone {
			return model.Contact{}, fmt.Errorf("contact with phone %q already exists", c.Phone)
		}
	}

	r.lastID++
	contact.ID = r.lastID
	r.storage[contact.ID] = contact

	return contact, nil
}

//ListAll method
func (r *ContactsRepositoryInMemory) ListAll() ([]model.Contact, error) {
	r.RLock()
	defer r.RUnlock()

	result := make([]model.Contact, 0, len(r.storage))
	for _, c := range r.storage {
		result = append(result, c)
	}

	return result, nil
}

//GetByID method
func (r *ContactsRepositoryInMemory) GetByID(id uint) (model.Contact, error) {
	r.RLock()
	defer r.RUnlock()

	contact, ok := r.storage[id]
	if !ok {
		return model.Contact{}, fmt.Errorf("record not found")
	}

	return contact, nil
}

//GetByPhone method
func (r *ContactsRepositoryInMemory) GetByPhone(phone string) (model.Contact, error) {
	r.RLock()
	defer r.RUnlock()

	for _, c := range r.storage {
		if c.Phone == phone {
			return c, nil
		}
	}

	return model.Contact{}, fmt.Errorf("record not found")
}

//GetByEmail method
func (r *ContactsRepositoryInMemory) GetByEmail(email string) (model.Contact, error) {
	r.RLock()
	defer r.RUnlock()

	for _, c := range r.storage {
		if c.Email == email {
			return c, nil
		}
	}

	return model.Contact{}, fmt.Errorf("record not found")
}

//SearchByName method
func (r *ContactsRepositoryInMemory) SearchByName(n string) ([]model.Contact, error) {
	r.RLock()
	defer r.RUnlock()

	result := make([]model.Contact, len(r.storage))
	for _, c := range r.storage {
		if strings.HasPrefix(c.FirstName, n) || strings.HasPrefix(c.LastName, n) {
			result = append(result, c)
		}
	}

	return result, nil
}

//Delete method
func (r *ContactsRepositoryInMemory) Delete(id uint) error {
	r.Lock()
	defer r.Unlock()

	delete(r.storage, id)

	return nil
}
