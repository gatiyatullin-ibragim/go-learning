package main
import "fmt"

//initialization of struct values
type Student struct {
	Name     string
	age      int
	gpa      float64
	courses  []string

}


func main() {
	var s1 Student
	s2 := Student{Name: "Bakytzhan", age: 23, gpa: 4.00, courses: []string{"golang", "pp1"}}
	s3 := Student{"Bakytzhan", 23, 4.00, []string{"golang", "pp1"}}

	p := &Student{Name: "Bakytzhan"}
	p.gpa = 3.5
	q := new(Student)


	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s1.courses==nil)
	fmt.Println(q)
}

