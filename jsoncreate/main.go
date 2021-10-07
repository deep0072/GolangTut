package main

import (
	"encoding/json"
	"fmt"
)

type courses struct { //intialising the struct
	Name     string
	Price    int
	Platform string
	Password string
	Tags     []string
}

func main() {

	fmt.Println("welcome to json data creation")

	jsonEncode()
}

func jsonEncode() {

	DeepCourse := []courses{

		{"Ruby", 560, "Deepcourse.com", "deep123", []string{"web-dev", "ruby"}}, // populate the struct

		{"python", 460, "Deepcourse.com", "deep123", []string{"web-dev", "py"}},

		{"js", 660, "Deepcourse.com", "deep123", []string{"web-dev", "js"}},

		
	}

	finalJson, err := json.Marshal(DeepCourse)
	//It allows you to do encoding of JSON data as defined in the RFC 7159.
	// When working with data we
	//usually want to encode some Go struct into a json string

	// to print json in reading manner we use "json.MarshalIndent"

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))

}
