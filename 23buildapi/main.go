/* important info when you get request as a json then use decode method to
   do manipulation in data . then after it "encode" method is used to send the json response  */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"author"` /* here we are using author that is   struct type will be add into in Course */

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
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "1", CourseName: "python", CoursePrice: 500, Author: &Author{FUllname: "Deepak Kumar", Website: "deeprofile.herokuapp.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Golang", CoursePrice: 500, Author: &Author{FUllname: "Deepak Kumar", Website: "deeprofile.herokuapp.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET") //here we are calling all function we mentioned below using HandleFunc method

	r.HandleFunc("/courses", getAllcourses).Methods("GET")       //get all courses
	r.HandleFunc("/course/{id}", getOnecourse).Methods("GET")    // here we calling getonecourse so we are passing id in {} same as in function it self.
	r.HandleFunc("/course", CreateOneCourse).Methods("POST")     //caling post method
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT") //update method
	r.HandleFunc("/course/{id}", DeleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

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

func getOnecourse(w http.ResponseWriter, r *http.Request) { // function to get course using id from user input
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through the courses. find matching id and return the response

	for _, course := range courses {

		if course.CourseId == params["id"] { //check if user requested id is in db
			json.NewEncoder(w).Encode(course)
			//NewEncoder() takes reponse write and "Encode()" method  encode the message that we want to send as response

			return

		}

	}
	json.NewEncoder(w).Encode("no course is found")

	return

}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create  One Course")
	w.Header().Set("Content-Type", "application/json")

	//what if user send empty request
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")

	}

	// what about data is sent by user  with cempty curly braces {}

	var course Course                           // creating struct type variable
	_ = json.NewDecoder(r.Body).Decode(&course) // here take the request and decode the course

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return

	}

	// generate unique id and append course in courses
	rand.Seed(time.Now().UnixNano())

	course.CourseId = strconv.Itoa(rand.Intn(100)) // string is converted into int using Itoa ==> equivalent to integer format

	courses = append(courses, course) // add created course id into courses

	json.NewEncoder(w).Encode(course) // send response
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("hello update the course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop,id,remove,add my new  id

	for index, course := range courses {
		if course.CourseId == params["id"] { //check id if available
			courses = append(courses[:index], courses[index+1:]...) // now remove the id using append fucntion

			//here first parameter is the values before the index that we find, 2nd param is the those value which will come after the inex we already got

			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course) // here it will take message and decode the json

			course.CourseId = params["id"]

			courses = append(courses, course) // add the data we want to add in course

			json.NewEncoder(w).Encode(course) //here sending encoded message that added data in json form
			return

		}

	}

}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("deleted one course")
	params := mux.Vars(r)

	// loop ,id ,remove

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			break

		}

	}

	json.NewEncoder(w).Encode("course is deleted")
}
