package main

import "fmt"

const Val = 6000 //here it is public variable means it can be used any where.
// variable should be start with   capital letter to declare public.

func main() {

	var username string = "hitesh"

	fmt.Println(username)

	fmt.Printf("Variable is type of: %T \n", username) // here PrintF is    used for formatting of string

	// bool variable decleration
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)

	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	// here i cant set value more than 255 if int value type is uint8
	// uint16 can take upto 65535
	// uint32 can take upto 4294967295
	// uint64 can take more than 429496795
	var smallVal uint8 = 255
	fmt.Println(smallVal)

	fmt.Printf("Variable is of type: %T \n", smallVal)

	// decleration of float value using float32
	var smallFloat float32 = 255.456298489649649646
	fmt.Println(smallFloat)

	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default value and some aliases

	// when you declare int variable witout assigning any
	// value it will set value to zero by default
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicite type

	// here variable declared witout defining its type
	// well it is done because in backend lexer define that "type" of variable
	var website = "httpa://depprofile.herokuapp.com"
	fmt.Println(website)

	// no var style using walrus operator.
	// But it is only possible if it is inside the method
	numberOfUser := 5000

	fmt.Println(numberOfUser, Val)
}
