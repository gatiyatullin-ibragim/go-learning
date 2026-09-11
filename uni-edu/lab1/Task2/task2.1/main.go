package main

import "fmt"

func main() {
	var s1 []int

	s2 := []int{1, 2, 3, 4, 5}

	s3 := make([]int, 5, 10)

	arr := [5]int{10, 20, 30, 40, 50}
	s4 := arr[1:4]

	s5 := new([]int)

	fmt.Println(s1, s2, s3, s4, s5)
}