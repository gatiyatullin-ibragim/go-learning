// Bonus: comparing interface values
package main

import "fmt"

func main() {
	var a, b any = 1, 1
	fmt.Println(a == b) // true: same dynamic type and value

	var c, d any = 1, int64(1)
	fmt.Println(c == d) // false: dynamic types differ (int vs int64)

	defer func() {
		fmt.Println("recovered:", recover())
	}()
	var e, f any = []int{1}, []int{1}
	fmt.Println(e == f) // PANIC: comparing uncomparable type []int
}
