package context

import (
	"github.com/google/uuid"
)

type User struct {
	id          uuid.UUID
	phoneNumber string
}

func (r *User) Id() uuid.UUID {
	return r.id
}

func (r *User) SetId(id uuid.UUID) *User {
	r.id = id
	return r
}

func (r *User) PhoneNumber() string {
	return r.phoneNumber
}

func (r *User) SetPhoneNumber(phoneNumber string) *User {
	r.phoneNumber = phoneNumber
	return r
}
