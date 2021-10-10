package main

import (
	"encoding/json"
	"fmt"
)

type courses struct { //intialising the struct
	Name     string `json:"courseName"` //here we are creatin assumed name(alias) of keys that will be showed in output
	Price    int
	Platform string
	Password string `json:"-"` // here "-" will hide the password field
	Tags     []string
}

func main() {

	fmt.Println("welcome to json data creation")

	jsonEncode()
	DecodeJson()
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
	if err != nil {
		panic(err)
	}

	// to print json in reading manner we use "json.MarshalIndent"

	finalJson1, err := json.MarshalIndent(DeepCourse, "", "\t")
	//here second parameter is prefix \t is tab will be given in json output

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
	fmt.Println(string(finalJson1))

}

func DecodeJson() {

	JsonDataFromweb := []byte(`
	{
		"courseName": "Ruby",

		"Price": 560,

		"Platform": "Deepcourse.com",
		"Password" : "1234556",

		"Tags": ["web-dev",	"ruby"]
	
				
	}
	
	
	
	`)

	var DeepCourse courses

	//now check json is valid

	checkValid := json.Valid(JsonDataFromweb)

	if checkValid {
		println("json is valid")

		json.Unmarshal(JsonDataFromweb, &DeepCourse) //now decode tha json using Unmarshal
		//here json data will be decode and store in to &Deepcourse

		fmt.Println(DeepCourse, "course") //print the encoded json data

	} else {
		fmt.Println("json was not valid")
	}

	// some cases where i just want to add data to key and value in map(dictionary)

	var myOnlinedata map[string]interface{} //here value could be int,string or array that we are getting from web so we used interface here\

	json.Unmarshal(JsonDataFromweb, &myOnlinedata)
	fmt.Printf("%#v\n", myOnlinedata)

	for key, value := range myOnlinedata {
		fmt.Printf("key is %v and value is %v and type is: %T\n", key, value, value)
	}

}
