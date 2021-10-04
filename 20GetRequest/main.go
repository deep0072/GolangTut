// http.Get(url) ==> ioutil.ReadAll(response.Body) ==> response.StatusCode

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("welcome to get request in  go lang")
	GetResponse()
}

func GetResponse() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl) //get url response

	if err != nil {
		panic(err)
	}
	defer response.Body.Close() //close the response

	fmt.Println("status code of response:", response.StatusCode) //get status of response
	fmt.Println("length of content is:", response.ContentLength) // get content length

	content, err := ioutil.ReadAll(response.Body) //read body of response

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}
