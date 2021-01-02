package main

import "fmt"

type person struct {
	name          string
	age           int
	favoriteFoods []string
}

func main() {
	favoriteFoods := []string{"kimchi", "pizza"}
	junu := person{name: "junu", favoriteFoods: favoriteFoods, age: 29}

	fmt.Println(junu)
}
