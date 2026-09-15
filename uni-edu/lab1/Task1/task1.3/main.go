package main

import "fmt"

func main() {
	arr := [2][2]int{
		{2, 2},
		{2, 2},
	}

	for _,  row := range arr {
		for _, cell := range row{
			fmt.Print(cell, "")
		}
		fmt.Println()
	}

}
