package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to uber app")

	reader := bufio.NewReader(os.Stdin) //buffio store the input value in reader

	fmt.Println("enter rating")

	input, _ := reader.ReadString('\n') // this line read the input value that is given by user

	numString, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //string.TrimSpace will terminate the trailing spaces form input string

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added extra one in your rating", numString+1)
	}
}
