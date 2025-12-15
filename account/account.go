package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?,.")

type Account struct {
	login    string
	password string
	url      string
}

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc *Account) OutputPassword() {
	color.Cyan(acc.login, acc.password, acc.url)
}

func (acc *Account) generatePassword(n int) {
	passwordRune := make([]rune, n)

	for i := range passwordRune {
		randNumber := rand.IntN(len(letterRunes))
		passwordRune[i] = letterRunes[randNumber]
	}

	acc.password = string(passwordRune)
}

func NewAccountWIthTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	acc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if password == "" {
		acc.generatePassword(8)
	}

	return acc, nil
}

// func newAccount(login, password, urlString string) (*account, error) {
// 	if login == "" {
// 		return nil, errors.New("EMPTY_LOGIN")
// 	}
// 	_, err := url.ParseRequestURI(urlString)
// 	if err != nil {
// 		return nil, errors.New("INVALID_URL")
// 	}
// 	acc := &account{
// 		login:    login,
// 		password: password,
// 		url:      urlString,
// 	}
// 	if password == "" {
// 		acc.generatePassword(8)
// 	}

// 	return acc, nil
// }
