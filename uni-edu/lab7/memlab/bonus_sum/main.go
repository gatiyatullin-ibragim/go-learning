package main

import "fmt"

func sum() int {
	nums := make([]int, 1000) // попробуйте 8192, 8193, 10000
	for i := range nums {
		nums[i] = i
	}
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() { fmt.Println(sum()) }
