package main

import (
	"fmt"
)

func main() {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAccount, err := newAccountWIthTimeStamp(login, password, url)
	if err != nil {
		fmt.Print(err)
		return
	}

	myAccount.generatePassword(8)
	myAccount.outputPassword()
	fmt.Println(myAccount)
}

func promtData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
