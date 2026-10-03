package main

import "fmt"

func main() {
	// 1. Init
	inventory := map[string]int{
		"apples":  10,
		"bananas": 5,
	}
	// 2. Read
	fmt.Println("Current stock of apples:", inventory["apples"])
	// 3. Update
	inventory["bananas"] = 12
	// 4. Insert
	inventory["oranges"] = 8
	// 5. Delete
	delete(inventory, "apples")
	// 6. Print
	fmt.Println("Updated inventory:", inventory)
}
