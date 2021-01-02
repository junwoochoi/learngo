package main

import (
	"fmt"

	"github.com/junwoochoi/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"

	dictionary.Add(word, "First")
	result, _ := dictionary.Search(word)
	fmt.Println(result)

	dictionary.Delete(word)
	fmt.Println("Delete")
	result2, err := dictionary.Search(word)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result2)
}
