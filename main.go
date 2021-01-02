package main

import "fmt"

func main() {
	junu := map[string]string{"name": "junu", "age": "19"}

	for key, value := range junu {
		fmt.Println(key, value)
	}
}
