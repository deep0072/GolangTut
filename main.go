package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the userInput"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) //bufio is a package used for buffered IO. Buffering IO is a technique used to temporarily
	//accumulate the results for an IO operation before transmitting it forward
	//it is used to store things in variable

	fmt.Println("enter the rating for my driving:")

	// input, err := reader.ReadString('\n') //here input var is like try and err is like catch.

	input, _ := reader.ReadString('\n')

	fmt.Println("thanks for you rating of", input)
	fmt.Printf("thanks for you rating of  %T ", input)

}
