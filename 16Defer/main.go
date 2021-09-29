package main

import "fmt"

func main() {
	defer fmt.Print("hi ") //here it will be execute after the print function
	fmt.Println("welcome ")

	defer fmt.Print("four ")
	defer fmt.Print("three ")
	defer fmt.Print("two ")
	fmt.Println("welcome to defer")
	myDefer()

	//output ===>>>   welcome welcome to defertwo three four hi

	// first those will be execute which has no any defer keyword then after lifo structure will be followed

	//like bottom to top

	// output ==> welcome
	// welcome to defer
	// 876543210two three four hi

}

func myDefer() {
	for i := 0; i < 9; i++ {
		defer fmt.Print(i) //here a queue will be genreate 012345678 and then according to defer counts will ne 876543210
	}
}
