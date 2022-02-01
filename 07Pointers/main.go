package main

import "fmt"

// pointers: btw why pointers do exists
// variable just refrence the memory address in which actual value is stored.
// well when we have a variable and want to pass its value to other function or method
// we are not passing  the actual varibale itself. we are just creating another copy of variable which create irregularities.
// so to avoid this we use Pointers

// so pointer is just direct refrence to memory address

func main() {
	fmt.Println("hi ")

	// var ptr *int //here ptr is kind of variable and astrisk "*" is just to show that ptr is pointer that is holding int values
	// we can create for strings,float also

	// fmt.Println(ptr)

	var myNumber = 23

	var ptr = &myNumber //here "&" sign means refrence.  "ptr" variable is refrencing to the memory address of "myNumber" variable

	fmt.Println("&myNumber:", &myNumber, ptr)

	fmt.Println("value of actual pointer is  ", ptr) //it will print the value of ptr that is holding the address of memory

	fmt.Println("value of actual pointer is  ", *ptr) //it will print the actual value that is hold by the memory. so * sign will give the value that is inside the pointer's memory address

	// lets take an example

	*ptr = *ptr * 2

	fmt.Println("new value is:", myNumber)

	// *ptr ==> give the value that is inside the memory. where memory is refrenced by &myNumber
	// so thats why mynumber value changed. because the *ptr hold the value of myNumber variable which is mulitplied by the 2
	// *ptr does not copy the value of myNumber
}
