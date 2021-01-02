package main

import (
	"fmt"

	"github.com/junwoochoi/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}

	word := "hello"
	definition := "Greeting"

	err := dictionary.Add(word, definition)

	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search("hello")

	fmt.Println(hello)

	err2 := dictionary.Add(word, definition)

	if err2 != nil {
		fmt.Println(err2)
	}
}
