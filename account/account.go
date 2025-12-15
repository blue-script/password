package account

import (
	"encoding/json"
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?,.")

type Account struct {
	Login     string `json:"login" xml:"test"`
	Password  string `json:"password"`
	Url       string `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) ToBytes() ([]byte, error) {
	file, err := json.Marshal(acc)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (acc *Account) generatePassword(n int) {
	passwordRune := make([]rune, n)

	for i := range passwordRune {
		randNumber := rand.IntN(len(letterRunes))
		passwordRune[i] = letterRunes[randNumber]
	}

	acc.Password = string(passwordRune)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	acc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       urlString,
		Login:     login,
		Password:  password,
	}
	if password == "" {
		acc.generatePassword(8)
	}

	return acc, nil
}
