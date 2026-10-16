// Task 5: Generic toolbox
package main

import (
	"cmp"
	"fmt"
)

type Number interface {
	~int | ~int64 | ~float64
}

type Score int

func Sum[T Number](nums []T) T {
	var s T
	for _, n := range nums {
		s += n
	}
	return s
}

func Average[T Number](nums []T) (float64, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	return float64(Sum(nums)) / float64(len(nums)), true
}

func MaxOf[T cmp.Ordered](items []T) (T, bool) {
	var zero T
	if len(items) == 0 {
		return zero, false
	}
	m := items[0]
	for _, v := range items[1:] {
		if v > m {
			m = v
		}
	}
	return m, true
}

func IndexOf[T comparable](items []T, target T) int {
	for i, v := range items {
		if v == target {
			return i
		}
	}
	return -1
}

func Filter[T any](items []T, keep func(T) bool) []T {
	var out []T
	for _, v := range items {
		if keep(v) {
			out = append(out, v)
		}
	}
	return out
}

func main() {
	scores := []Score{90, 60, 75, 82, 49}

	fmt.Println("Sum:", Sum(scores))
	fmt.Println("Sum float:", Sum([]float64{1.5, 2.5}))
	fmt.Println(Average(scores))
	fmt.Println(Average([]Score{})) // 0 false

	fmt.Println(MaxOf(scores))
	fmt.Println(MaxOf([]int{3, 9, 4}))
	fmt.Println(MaxOf([]string{"apple", "banana", "cherry"}))
	fmt.Println(MaxOf([]int{})) // 0 false

	fmt.Println(IndexOf(scores, 75))                     // 2
	fmt.Println(IndexOf(scores, 100))                    // -1
	fmt.Println(IndexOf([]string{"go", "rust"}, "rust")) // 1

	fmt.Println(Filter(scores, func(s Score) bool { return s >= 75 })) // [90 75 82]
}
