package main

import "fmt"

func main() {
	var arr = [5][5]string{
		{"-", "*", "-", "*", "-"},
		{"*", "-", "*", "-", "*"},
		{"-", "*", "-", "*", "-"},
		{"*", "-", "*", "-", "*"},
		{"-", "*", "-", "*", "-"},
	}

	for _, row := range arr {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}

}
