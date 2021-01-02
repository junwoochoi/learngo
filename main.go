package main

import (
	"fmt"

	"github.com/junwoochoi/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("junu")
	account.Deposite(100)
	fmt.Println(account.Balance())

	err := account.Withdraw(200)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)
}
