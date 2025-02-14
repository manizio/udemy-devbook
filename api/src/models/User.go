package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Nick string `json:"nick,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdat,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {

	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if user.Nick== "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if user.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O email inserido não é válido")
	}
	if step == "cadastro" && user.Password == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "cadastro" {
		hashPass, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPass)
	}

	return nil
}
