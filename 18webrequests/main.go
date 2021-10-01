package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("hi welcome to web requests")

	response, err := http.Get(url)
	// http.Get will be used to fetch the data from url which is kind of intialisation of connection

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response  is of type: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	// as i used here defer so it will be execute in the end

	data, err := ioutil.ReadAll(response.Body) //iutil.ReadAll will read the response in proper manner

	if err != nil {
		panic(err)
	}

	fmt.Println("data is:", string(data)) //it will give the response of actual data we want make sure convert response body into string

}
