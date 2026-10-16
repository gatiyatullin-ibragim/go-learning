// Task 7: Tour of Go -> Generic types (List[T])
package main

import (
	"fmt"
	"strings"
)

type List[T any] struct {
	next *List[T]
	val  T
}

// Push puts v at the front and returns the new head. Works on a nil list.
func (l *List[T]) Push(v T) *List[T] {
	return &List[T]{next: l, val: v}
}

// Len works on a nil list too (the loop just does not run).
func (l *List[T]) Len() int {
	n := 0
	for ; l != nil; l = l.next {
		n++
	}
	return n
}

func (l *List[T]) String() string {
	var parts []string
	for ; l != nil; l = l.next {
		parts = append(parts, fmt.Sprint(l.val))
	}
	return "[" + strings.Join(parts, " ") + "]"
}

func main() {
	var l *List[int] // nil list
	fmt.Println(l.Len(), l)
	l = l.Push(1)
	l = l.Push(2)
	l = l.Push(3)
	fmt.Println(l.Len(), l) // 3 [3 2 1]

	var words *List[string]
	words = words.Push("go").Push("rust")
	fmt.Println(words.Len(), words) // 2 [rust go]
}
