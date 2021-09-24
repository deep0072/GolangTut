package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to switch case")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("number is ", diceNumber)

	switch diceNumber { //switch "variable"  variable is used on which switch case will be applied

	case 1:
		fmt.Println("value is 1 and u can open")
	case 2:
		fmt.Println("value is 2 and u can move 2 spot")
	case 3:
		fmt.Println("value is 3 and u can move to 3 spot ")
	case 4:
		fmt.Println("value is 4 and u can move to 4 spot ")
	case 5:
		fmt.Println("value is 5 and u can move to 5 spot ")
	case 6:
		fmt.Println("value is 6 and u can move to 6 pot and roll the dice again")
	default: //it is used to handle random number that is not mentioned in switch case
		fmt.Println("well what was that")

	}
}
