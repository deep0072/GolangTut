package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func getBooks(c *gin.Context) { //gincontext is a type of context which is used to handle the request and response
	c.IndentedJSON(http.StatusOK, books) //c.IndentedJSON is used to send the response in json format
}

// function createBook is used to create a new book in the books slice variable

// to run post method  in terminal use ```curl localhost:8080/create --include --header "Content-Type: application/json" -d @body.json --request "POST"`
func createBooks(c *gin.Context) {
	var newbook book
	if err := c.BindJSON(&newbook); err != nil {
		return
	} //bind the json data to the newbook variable, kind of conversion from json to struct type

	books = append(books, newbook)            //add the new book to the books slice variable
	c.IndentedJSON(http.StatusCreated, books) // return status of creation for books

}

// function to get book by its id
func getaBookByid(id string) (*book, error) {

	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("Book not found")

}

func bookByid(c *gin.Context) {
	fmt.Println(books)
	id := c.Param("id")           //get the id from the url
	book, err := getaBookByid(id) //get the book by its id
	if err != nil {
		return
	}

	c.IndentedJSON(http.StatusOK, book) //return the book in json format
}

// function checkoutBook is used to checkout a book from the books slice variable
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	book, err := getaBookByid(id)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Book is not available"})
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()                 // allows to handle different routes
	router.GET("/books", getBooks)          // triggget the getbooks function when the /books is called
	router.POST("/create", createBooks)     // trigger the createbooks function when the /create is called
	router.GET("/book/:id", bookByid)       // trigger the bookByid function when the /book/:id is called
	router.PATCH("/checkout", checkoutBook) // trigger the checkoutBook function when the /checkout is called this is kind  of updating
	router.Run("localhost:8080")            // it run the local host on port 8080

}
