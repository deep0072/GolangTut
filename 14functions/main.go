// we can define function inside the function
//main function which execute automatically
package main

import "fmt"

func main() {

	fmt.Println("welcome to my functions")
	Greet()

	result := Adder(5, 4)
	advance, message := advacnedAdder(4, 5, 6, 6, 67, 8, 9)

	fmt.Println("sum is:", result)

	fmt.Println("advancesum is:", advance)

	fmt.Println("message is:", message)

}

func Greet() { //to execute this function we call it in main function given above
	fmt.Println("hello namstey!")

}

func Adder(val1 int, val2 int) int {
	return val1 + val2
	//when return keyword used inside the function we always mention type of return
}

func advacnedAdder(values ...int) (int, string) { //here "...int" keyword will take multiple int values
	total := 0
	for _, val := range values {

		total += val

	}
	return total, "hi this my power"
}
