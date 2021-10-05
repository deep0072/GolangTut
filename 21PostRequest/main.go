package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("hello welcome to post request")
	PerformPostUrl()
}

func PerformPostUrl() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`  
	{
		"coursename": "Golang",
		"Price": 0,
		"Platform": "Youtube"
	}`) //creat content as json

	response, err := http.Post(myUrl, "application/json", requestBody) //use .Post method takes 3 param url, type of content, content
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body) //get the response
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))

}
