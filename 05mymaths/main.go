package main

import (
	"crypto/rand"
	"math/big"

	"fmt"
	// "math/rand"
)

func main() {
	fmt.Println("hi there!")

	var myNymberone int = 2
	var myNumbertwo float64 = 4

	fmt.Println("sum of two number is", myNymberone+int(myNumbertwo))

	// generate random number  using crypot rand and math/rand

	//using math/rand

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random number generation using crypto

	myRandnumber, _ := rand.Int(rand.Reader, big.NewInt(10))

	fmt.Println(myRandnumber)

}
