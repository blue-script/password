package main

import (
	"fmt"

	"github.com/blue-script/password/account"
	"github.com/blue-script/password/files"
	"github.com/blue-script/password/output"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("Password manager")
	vault := account.NewVault(files.NewJsonDb("data.json"))
	// vault := account.NewVault(cloud.NewCloudDb("data.json"))
Menu:
	for {
		choice := promptData([]string{"1. Создать аккаунт", "2. Найти аккаунт", "3. Удалить аккаунт", "4. Выход", "Choose variant"})

		switch choice {
		case "1":
			createAccount(vault)
		case "2":
			findAccount(vault)
		case "3":
			removeAccount(vault)
		default:
			fmt.Println("Выход из программы.")
			break Menu
		}
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

func findAccount(vault *account.VaultWithDb) {
	url := promptData([]string{"Введите URL для поиска"})
	data := vault.FindAccountsByURL(url)
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
