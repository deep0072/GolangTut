package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("hello welcome to post request")
	PerformPostUrl()
	PerformPostFormRequest()
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

//Post form data as json fpr example in use of image upload
func PerformPostFormRequest() {

	const myUrl = "http://localhost:8000/postform"

	data := url.Values{} //use url.Values{} to add data in variable

	data.Add("firstname", "Deepak")
	data.Add("lastname", "Kumar")
	data.Add("Age", "23")

	response, err := http.PostForm(myUrl, data) //post form data
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()                   //close connection
	content, err := ioutil.ReadAll(response.Body) //read response

	fmt.Println(string(content)) //get content

}
