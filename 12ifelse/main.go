package main

import "fmt"

func main() {
	fmt.Println("if else")

	loginCount := 10
	var result string

	if loginCount < 10 {

		result = "less than 10"

	} else if loginCount > 10 {
		result = "greater than 10"
	} else {
		result = "equal to 10"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("number less than 10")
	} else {
		fmt.Println("greater than 10")
	}
}
