package main

import (
	"fmt"

	"github.com/junwoochoi/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("junu")
	fmt.Println(account)
}
