package main

import "fmt"

//what is output?

type Team struct {
	Name    string
	Scores  [3]int
	Members []string
}

func main() {
	t1 := Team{Name: "gooners", Scores: [3]int{1,2,3}, Members:[]string{"ali", "lox"}}
	t2 := t1
	
	t2.Name	= "Debily"
	t2.Scores[0] = 100
	t2.Members[0] = "Nurlan"

	fmt.Println(t1.Name, t1.Scores, t1.Members)
	fmt.Println(t2.Name, t2.Scores, t2.Members)
}

