package main

import "fmt"

func main() {
	fmt.Println("hello welcome to the methods")

	Deepak := User{"Deepak", 23, "True"}

	Deepak.GetStatus() //here we are calling method

	fmt.Println("age after 2 years:", Deepak.Multiply(2)) //here we are calling second method

}

/*Go methods are similar to Go function with one difference,
i.e, the method contains a receiver argument in it. With the help of the receiver argument,
the method can access the properties of the receiver.
Here, the receiver can be of struct type or non-struct type */

/* method syntax is ===>>>      func(reciver_name Type) method_name(parameter_list)(return_type) {
// Code
}*/

type User struct {
	Name   string
	Age    int
	Status string
}

func (u User) GetStatus() { //here u is reciever name and User is type
	fmt.Println("is user active:", u.Status)
}

func (u User) Multiply(val2 int) int {
	return u.Age + val2
}
