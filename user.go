package contextplus

import (
	"github.com/google/uuid"
)

type User struct {
	id                  uuid.UUID
	email               string
	emailVerified       bool
	phoneNumber         string
	phoneNumberVerified bool
}

func (r *User) Id() uuid.UUID {
	return r.id
}

func (r *User) SetId(id uuid.UUID) {
	r.id = id
}

func (r *User) Email() string {
	return r.email
}

func (r *User) SetEmail(email string) {
	r.email = email
}

func (r *User) EmailVerified() bool {
	return r.emailVerified
}

func (r *User) SetEmailVerified(emailVerified bool) {
	r.emailVerified = emailVerified
}

func (r *User) PhoneNumber() string {
	return r.phoneNumber
}

func (r *User) SetPhoneNumber(phoneNumber string) {
	r.phoneNumber = phoneNumber
}

func (r *User) PhoneNumberVerified() bool {
	return r.phoneNumberVerified
}

func (r *User) SetPhoneNumberVerified(phoneNumberVerified bool) {
	r.phoneNumberVerified = phoneNumberVerified
}
