package main

import "fmt"


type Person struct{
	Name   string
	Email  string
}

func (p Person) Contact() string {
	return p.Name + " <" + p.Email + ">"
}

type Student struct {
	Person
	GPA float64
}

type Teacher struct {
	Person
	Department string
}

func (t Teacher) Contact() string {
	return t.Person.Contact() + " (" + t.Department + ")"
}

func main(){
	student := Student{
		Person: Person{Name: "Ibragim", Email: "student@example.com"},
		GPA:    3.0,
	}
	teacher := Teacher{
		Person:     Person{Name: "Alex Smith", Email: "teacher@example.com"},
		Department: "Computer Science",
	}

	fmt.Println(student.Contact())
	fmt.Println(teacher.Contact())
	fmt.Println(teacher.Person.Contact())
}
