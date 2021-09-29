package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("ho welcome to files handling")

	content := " hi check the text file" //this is content that will be in file

	file, err := os.Create("./myfile.txt") // create text file

	if err != nil {
		panic(err) //it will check error during creation of file
		// panic keyword ==>  just shut down the execution of program and show the error

	}

	length, err := io.WriteString(file, content) //here "io" package will write the content into file

	if err != nil {
		panic(err)
	}

	fmt.Println("length is :", length)

	defer file.Close() // here file should need to be close after intialisation. and here i used defer to execute it in the end
}
