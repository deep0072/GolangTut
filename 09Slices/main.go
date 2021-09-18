package main

import (
	"fmt"
	"sort"
)

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

	// another way to declare the slices using make([]typeOfSlice, intial number of  slice can have)

	highscore := make([]int, 5)

	highscore[0] = 1
	highscore[1] = 5
	highscore[2] = 5
	highscore[3] = 5

	fmt.Println("these are the slices:", highscore)

	//now to add more data into slice we can use append method

	highscore = append(highscore, 90, 45, 5, 45, 45)

	sort.Ints(highscore) //sort the slice in Ascending order

	fmt.Println(highscore)
}
