package main

import (
	"fmt"
	"strings"

	"github.com/blue-script/password/account"
	"github.com/blue-script/password/files"
	"github.com/blue-script/password/output"
	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": removeAccount,
}

func main() {
	fmt.Println("Password manager")
	vault := account.NewVault(files.NewJsonDb("data.json"))
	// vault := account.NewVault(cloud.NewCloudDb("data.json"))
Menu:
	for {
		choice := promptData([]string{"1. Создать аккаунт", "2. Найти аккаунт по URL", "3. Найти аккаунт по логину ", "4. Удалить аккаунт", "5. Выход", "Choose variant"})

		menuFunc := menu[choice]
		if menuFunc == nil {
			fmt.Println("Выход из программы.")
			break Menu
		}
		menuFunc(vault)
	}
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError(err)
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска"})
	data := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	if len(data) == 0 {
		output.PrintError("Not found")
	}
	for _, acc := range data {
		acc.Output()
	}
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData([]string{"Введите логин для поиска"})
	data := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	if len(data) == 0 {
		output.PrintError("Not found")
	}
	for _, acc := range data {
		acc.Output()
	}
}

func removeAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для удаления"})

	isSuccess := vault.DeleteAccountByUrl(url)

	if !isSuccess {
		output.PrintError("Delete error")
		return
	}
	color.Green("Successful delete")

}

func promptData[T any](prompt []T) string {
	for index, line := range prompt {
		if index == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

	var res string
	fmt.Scan(&res)
	return res
}
