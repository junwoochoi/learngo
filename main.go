package main

import (
	"fmt"
)

func repeat(words ...string) {
	fmt.Println(words)
}

func main() {
	repeat("choi", "junwoo", "eddie", "luke")
}
