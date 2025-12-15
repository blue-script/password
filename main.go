package main

import (
	"fmt"

	"github.com/blue-script/password/account"
	"github.com/blue-script/password/files"
)

func main() {
	fmt.Println("Password manager")
Menu:
	for {
		choice := getMenuChoice()

		switch choice {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			removeAccount()
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

func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault := account.NewVault()
	vault.AddAccount(*myAccount)

	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Not successful marshal account")
		return
	}
	files.WriteFile(data, "data.json")
}

func findAccount() {}

func removeAccount() {}

func promptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
