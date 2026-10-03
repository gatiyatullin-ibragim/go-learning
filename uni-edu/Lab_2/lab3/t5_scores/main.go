package main

import "fmt"

func main() {
	scores := map[string]int{
		"Alice": 90,
		"Bob":   0, // Bob took the exam but scored 0
	}

	fmt.Println(`scores["Bob"] =`, scores["Bob"], `| scores["Charlie"] =`, scores["Charlie"], "(оба 0!)")

	if _, ok := scores["Bob"]; ok {
		fmt.Println("Bob is in the map")
	} else {
		fmt.Println("Bob is missing")
	}

	if _, ok := scores["Charlie"]; ok {
		fmt.Println("Charlie is in the map")
	} else {
		fmt.Println("Charlie is missing")
	}
}
