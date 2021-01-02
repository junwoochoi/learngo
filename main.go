package main

import (
	"fmt"

	"github.com/junwoochoi/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary["hello"] = "world"

	result, err := dictionary.Search("hello")

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
