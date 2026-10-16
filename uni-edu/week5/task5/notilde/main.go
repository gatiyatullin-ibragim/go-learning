// Step 5: Number WITHOUT ~ -> compile error. (This file intentionally fails.)
package main

import "fmt"

type Number interface {
	int | int64 | float64 // no tilde!
}

type Score int

func Sum[T Number](nums []T) T {
	var s T
	for _, n := range nums {
		s += n
	}
	return s
}

func main() {
	fmt.Println(Sum([]Score{1, 2, 3}))
}
