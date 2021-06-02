package auth

import (
	"fmt"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

var bcryptCost = 13

type User struct {
	ID          int64  `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Contact_Num string `json:"contact"`
	PassHash    []byte `json:"-"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	Name         string `json:"name"`
	Contact_Num  string `json:"contact"`
}

type Updates struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Contact_Num string `json:"contact"`
}

//Validate validates the new user and returns an error if
//any of the validation rules fail, or nil if its valid
func (nu *NewUser) Validate() error {
	email, err := mail.ParseAddress(nu.Email)
	if err != nil {
		return fmt.Errorf("invalid email address %v", email)
	}
	if len(nu.Password) < 6 {
		return fmt.Errorf("password %v is less than 6 characters", nu.Password)
	}
	if nu.Password != nu.PasswordConf {
		return fmt.Errorf("password doesn't match confirmed password")
	}
	if len(nu.Name) <= 0 {
		return fmt.Errorf("provide a name")
	}
	return nil
}

//ToUser converts the NewUser to a User, setting the
//PhotoURL and PassHash fields appropriately
func (nu *NewUser) ToUser() (*User, error) {
	err := nu.Validate()
	if err != nil {
		return nil, err
	}
	user := &User{}
	user.Email = nu.Email
	user.ID = 0
	user.Name = nu.Name
	user.SetPassword(nu.Password)
	return user, nil
}

//SetPassword hashes the password and stores it in the PassHash field
func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		return err
	}
	u.PassHash = hashedPassword
	return nil
}

//Authenticate compares the plaintext password against the stored hash
//and returns an error if they don't match, or nil if they do
func (u *User) Authenticate(password string) error {
	err := bcrypt.CompareHashAndPassword(u.PassHash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

//ApplyUpdates applies the updates to the user. An error
//is returned if the updates are invalid
func (u *User) ApplyUpdates(updates *Updates) error {
	if updates.Name == "" {
		return fmt.Errorf("no new name provided, sticking with existing")
	}
	u.Name = updates.Name
	return nil
}
