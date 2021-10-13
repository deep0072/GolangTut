package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go list -m all ==> list all package that perticular related to program file
func main() {
	fmt.Println("welcome to the modules")

	r := mux.NewRouter()                        //create route to handle request and response using gorilla mux
	r.HandleFunc("/", ServeHome).Methods("GET") // it will trigger the response writer Method which is GEt method

	log.Fatal(http.ListenAndServe(":4000", r)) //use port 4000 to show the response
}

func Greeter() {
	fmt.Println("welcome to modules in go lang")
}

func ServeHome(w http.ResponseWriter, r *http.Request) { //function is used to take request and write response
	w.Write([]byte("<h1> welcome to Deepak world</h1?"))
}
