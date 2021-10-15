package main

import "fmt"

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
