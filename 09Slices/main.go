package main

import "fmt"

func main() {

	fmt.Println("Welcome to slices")

	// slices just like the array

	//  so to create slices we use []typeofslice. but we do not mention to number that our slice will contain so we can add as many values we want.
	//and then intialise it with using curly brace in the end

	var fruitList = []string{"Apple", "Banana"} //to fill the values in slice. type inside the curly braces

	fmt.Printf("type of fruitList is %T\n", fruitList)
	fmt.Println("fruitList is ", fruitList)

	// to add data in slices we use append method syntax is append(slice_var_name, multiple data wrt its types)

	fruitList = append(fruitList, "mango", "fig", "orange")

	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) //this  slice the our slice using [start:stop]

	fmt.Println(fruitList)

}
