package main

import (
	"fmt"
	"strings"

	"github.com/blue-script/password/account"
	"github.com/blue-script/password/encrypter"
	"github.com/blue-script/password/files"
	"github.com/blue-script/password/output"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": removeAccount,
}

var menuVariants = []string{
	"1. Создать аккаунт",
	"2. Найти аккаунт по URL",
	"3. Найти аккаунт по логину ",
	"4. Удалить аккаунт",
	"5. Выход",
	"Choose variant",
}

func menuCounter() func() {
	i := 0
	return func() {
		i++
		fmt.Println("Количество вызовов: ", i)
	}
}

func main() {
	fmt.Println("Password manager")
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Not found env file")
	}

	vault := account.NewVault(files.NewJsonDb("data.json"), *encrypter.NewEncrypter())
	// vault := account.NewVault(cloud.NewCloudDb("data.json"))
Menu:
	for {
		choice := promptData(menuVariants...)

		menuFunc := menu[choice]
		if menuFunc == nil {
			fmt.Println("Выход из программы.")
			break Menu
		}
		menuFunc(vault)
	}
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

func findAccountByUrl(vault *account.VaultWithDb) {
	url := promptData("Введите URL для поиска")
	data := vault.FindAccounts(url, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)
	})
	outputResult(&data)
}

func findAccountByLogin(vault *account.VaultWithDb) {
	login := promptData("Введите логин для поиска")
	data := vault.FindAccounts(login, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)
	})
	outputResult(&data)
}

func outputResult(accounts *[]account.Account) {
	if len(*accounts) == 0 {
		output.PrintError("Not found")
	}
	for _, acc := range *accounts {
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

func promptData(prompt ...string) string {
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
