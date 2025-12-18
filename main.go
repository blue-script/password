package main

import (
	"fmt"

	"github.com/blue-script/password/account"
	"github.com/blue-script/password/files"
	"github.com/blue-script/password/output"
	"github.com/fatih/color"
)

func main() {
	output.PrintError(1)
	output.PrintError("sdfdsfdsf")

	fmt.Println("Password manager")
	vault := account.NewVault(files.NewJsonDb("data.json"))
	// vault := account.NewVault(cloud.NewCloudDb("data.json"))
Menu:
	for {
		choice := getMenuChoice()

		switch choice {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			removeAccount(vault)
		default:
			fmt.Println("Выход из программы.")
			break Menu
		}
	}
}

func getMenuChoice() int8 {
	var choice int8
	fmt.Println("Меню:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
	fmt.Scan(&choice)
	return choice
}

func createAccount(vault *account.VaultWithDb) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError(err)
		return
	}
	vault.AddAccount(*myAccount)
}

func findAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	data := vault.FindAccountsByURL(url)
	if len(data) == 0 {
		output.PrintError("Not found")
	}
	for _, acc := range data {
		acc.Output()
	}
}

func removeAccount(vault *account.VaultWithDb) {
	url := promptData("Введите URL для удаления")

	isSuccess := vault.DeleteAccountByUrl(url)

	if !isSuccess {
		output.PrintError("Delete error")
		return
	}
	color.Green("Successful delete")

}

func promptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
