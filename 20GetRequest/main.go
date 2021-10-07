// http.Get(url) ==> ioutil.ReadAll(response.Body) ==> response.StatusCode ==> response.Body.Close()

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
	defer response.Body.Close() //close the response as defer i used so it will be execute in the end

	fmt.Println("status code of response:", response.StatusCode) //get status of response
	fmt.Println("length of content is:", response.ContentLength) // get content length

	content, err := ioutil.ReadAll(response.Body) //read actual content that is body of response

	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
	
}
