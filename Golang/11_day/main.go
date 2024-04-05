package main

// model for course-- usually this goes in seprate file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//Case ID: 171222257800802

//fake database

var courses []Course

//middleware or helper file -- this goes in seprate file
func (c *Course) IsEmpty() bool {

}
func main() {

}
