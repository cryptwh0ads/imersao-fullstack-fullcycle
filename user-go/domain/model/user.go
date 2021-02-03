package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type User struct {
	ID string
	Name string
	Email string
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(email string, name string) (*User, error) {
	user := User{
		Email: email,
		Name: name,
	}
	user.ID = uuid.NewV4().String()
	err := user.isValid()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
