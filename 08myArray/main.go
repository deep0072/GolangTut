package main

import "fmt"

func main() {
	fmt.Println("hello! ")

	// defining of array in go lang [lenght of array]typeofArray

	var fruitList [5]string

	// this the another way to declare the array
	var vegList = [3]string{"onion", "potato", "tomato"}

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "grape"

	fruitList[4] = "fig"

	fmt.Println(fruitList)
	fmt.Println("veg list:", vegList)

	// output:==>>    [Apple Banana grape  fig]

	// as we did not assign any value at index 3 so it will automatically
	//  add the space at that position or index

	// to get the lenght of array we use the len(array)
	fmt.Println("lenght of array is:", len(fruitList))
}
