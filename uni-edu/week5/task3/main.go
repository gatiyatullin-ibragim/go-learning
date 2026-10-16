// Task 3: Find the bug (nil error that is not nil)
package main

import (
	"errors"
	"fmt"
)

type ValidationError struct{ Field string }

func (e *ValidationError) Error() string { return "invalid field: " + e.Field }

// BUGGY version: returns a typed nil pointer wrapped in an error interface
func validateBuggy(age int) error {
	var err *ValidationError
	if age < 0 {
		err = &ValidationError{Field: "age"}
	}
	return err
}

// FIXED version: return the interface, and a literal nil when no error
func validate(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age"}
	}
	return nil
}

func main() {
	fmt.Println("=== buggy ===")
	err := validateBuggy(20)
	if err != nil {
		fmt.Println("error:", err) // prints: error: <nil>
	}
	fmt.Println("err == nil:", err == nil) // false
	fmt.Printf("%%T of err: %T\n", err)    // *main.ValidationError

	fmt.Println("\n=== fixed ===")
	err = validate(20)
	fmt.Println("validate(20) == nil:", err == nil) // true

	err = validate(-1)
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("validate(-1): field =", ve.Field) // age
	}
}
