package main

import "fmt"

type Student struct{
	Name     string
	ID       string
	GPA      float64
	Cources  []string
}

//function addCourse that appends a course to a courses
func addCourse1(s Student, course string){
	s.Cources = append(s.Cources, course)
} 

func addCourse2(s *Student, course string){
	s.Cources = append(s.Cources, course)
}

func main() {
	//declaring in 4 ways
	var s1 Student
	s2 := Student{Name:"Ibragim", ID:"24B030254", GPA:3.00, Cources: []string{"GO", "Differential equations"}}
	s3 := Student{"Ibragim", "24B030254", 3.00, []string{"golang", "Diffur"}}

	q := new(Student)
	
	addCourse1(s2, "Physical education")
	addCourse2(&s2, "Linear algebra")

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(q)

}


