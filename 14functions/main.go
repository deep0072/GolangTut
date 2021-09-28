// we can define function inside the function
//main function which execute automatically
package main

import "fmt"

func main() {

	fmt.Println("welcome to my functions")
	Greet()

	result := Adder(5, 4)

	fmt.Println("sum is:", result)

}

func Greet() { //to execute this function we call it in main function given above
	fmt.Println("hello namstey!")

}



func Adder(val1 int, val2 int) int {
	return val1 + val2
}

//when return keyword used inside the function we always mention type of return
