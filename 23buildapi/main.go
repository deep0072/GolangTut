package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"` // here we are using author that is struct type will be add into in Course
}

type Author struct {
	FUllname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB
var courses []Course

//middleware, helper - file
func (c *Course) IsEmpty() bool { // here function taking parameter with return type bool
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

	fmt.Println("welcome to building of api in golang")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome to Api by Deepcourse "))

}

func getAllcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses) //this line encode the response which will be passed in Encode() method
	//when we inject some fake data values in data base is known as seeding process

}

func getOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through the courses. find matching id and return the response

	for _, course := range courses {

		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			//NewEncoder() takes reponse write and Encode() encode the message that we want to send as response

			return

		}

	}
	json.NewEncoder(w).Encode("no course is found")
	return

}
