package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

type book struct {
	ID       string `json:"id"`
	title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", title: "Book One", Author: "John Doe", Quantity: 1},
	{ID: "2", title: "Book Two", Author: "Steve Smith", Quantity: 2},
	{ID: "3", title: "Book Three", Author: "Jane Doe", Quantity: 3},
}


// func to get the all books from the books slice variable

func getBooks(c *gin.Context) {  //gincontext is a type of context which is used to handle the request and response
	c.IndentedJSON(http.StatusOK, books) //c.IndentedJSON is used to send the response in json format
}



func main() {
	router := gin.Default() // allows to handle different routes
	router.GET("/books", getBooks) // triggget the getbooks function when the /books is called
	router.Run("localhost:8080") // it run the local host on port 8080

}
