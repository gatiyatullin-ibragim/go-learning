// DEMO ONLY: do not run! Error() calls fmt.Sprint(e) -> infinite recursion.
// Run `go vet ./task4/recursive` to see how vet detects it.
package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint(e)
}

func main() {
	fmt.Println(ErrNegativeSqrt(-2))
}
