package main

import "fmt"

func main() {

	fmt.Println(" hi welcome to structs ")
	//no inheritance in go lang

	//syntax
	//    ||
	//    \/
	// type Varibale_name struct {
	// 	Varaible1 string
	// 	Variable2 string
	// }

	deepak := User{"Deepak", "deepak@go.dev", true, 23}

	fmt.Println("name is ",deepak.Name)
	fmt.Printf("deepak details are : %+v\n", deepak) // output ==>> {Name:Deepak Email:deepak@go.dev Status:true Age:23}

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
