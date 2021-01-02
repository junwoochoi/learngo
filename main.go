package main

import "fmt"

func main() {
	names := []string{"choi", "junwoo", "eddie"}

	names = append(names, "blahblah1")
	names = append(names, "blahblah3")
	names = append(names, "blahblah2")

	fmt.Println(names)
}
