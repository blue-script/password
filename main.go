package main

import (
	"fmt"
	"github.com/blue-script/password/account"
)

func main() {
	login := promtData("Введите логин")
	password := promtData("Введите пароль")
	url := promtData("Введите URL")

	myAccount, err := account.NewAccountWIthTimeStamp(login, password, url)
	if err != nil {
		fmt.Print(err)
		return
	}

	myAccount.OutputPassword()
	fmt.Println(myAccount)
}

func promtData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
