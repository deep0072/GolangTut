package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
q
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

func getOnecourse(w http.ResponseWriter, r *http.Request) { // function to get course using id from user input id
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through the courses. find matching id and return the response

	for _, course := range courses {

		if course.CourseId == params["id"] { //check if user requested id is in db
			json.NewEncoder(w).Encode(course)
			//NewEncoder() takes reponse write and Encode() encode the message that we want to send as response

			return

		}

	}
	json.NewEncoder(w).Encode("no course is found")

	return

}

func CreateOneCOurse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create  One Course")
	w.Header().Set("Content-Type", "application/json")

	//what is user send empty request
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")

	}

	// what about data is sent by user  with cempty curly braces {}

	var course Course // creating struct type variable 
	_ = json.NewDecoder(r.Body).Decode(&course) // here take the request and decode the course

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return

	}

	// generate unique id and append course in courses

	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100)) //string is converted into int using Itoa ==> equivalent to integer format

	courses = append(courses, course) // add created course id into courses

	json.NewEncoder(w).Encode(course) //send response
	return

}
