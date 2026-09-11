package main

import "fmt"

func main() {
	arr := [5]int{5, 4, 3, 2, 1}

	slice := arr[2:4]

	fmt.Println("arrray: ", arr)
	fmt.Println("slice: ", slice)

	slice[0] = 1000

	fmt.Println("\nafter change: ")
	fmt.Println("array: ", arr)
	fmt.Println("slice: ", slice)

	arr[2] = 2000

	fmt.Println("\nafter second change: ")
	fmt.Println("array: ", arr)
	fmt.Println("slice: ", slice)
}
