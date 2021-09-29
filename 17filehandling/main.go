package main

import (
	"fmt"
	"io"
	"io/ioutil"
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

	readFile("./myfile.txt")

	defer file.Close() // here file should need to be close after intialisation. and here i used defer to execute it in the end
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName) // ioutil will read the file

	if err != nil {
		panic(err)

	}

	fmt.Println("file", data)
	//output ===> file [32 104 105 32 99 104 101 99 107 32 116 104 101 32 116 101 120 116 32 102 105 108 101]

	fmt.Println("file", string(data)) //output ===>>> file  hi check the text file

	//if i convert data variable in string format then actual data will be shown otherwise it will shown in bytes

}
