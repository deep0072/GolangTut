package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursenam=reactjs&paymentid=jsdnkjn"

func main() {
	fmt.Println("welcome to handling url in golang")

	result, err := url.Parse(myUrl) //to parse the url information  such as host,port etc we net package

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	//if i want to extraxt raw query seperatly .Query () method

	params := result.Query()

	fmt.Println(params) // this eill gnerate dictionary or map

	//output ===>> map[coursenam:[reactjs] paymentid:[jsdnkjn]]

	for _, val := range params { // in for loop only values will be printed
		fmt.Println(val)
	}

	// construct proper url using url.URL with scheme, Host, Path, RawPath

	partsOfUrl := &url.URL{
		Scheme:  "https",    //Scheme is like https connection
		Host:    "deep.dev", //Host is like web domain
		Path:    "/bio",
		RawPath: "/info=Deep",
	}

	fmt.Println(partsOfUrl.String())

}
