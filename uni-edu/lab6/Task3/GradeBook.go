package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("student not found")

type GradeError struct {
	Student string
	Grade   int
}

func (e *GradeError) Error() string {
	return fmt.Sprintf("invalid grade %d for %s", e.Grade, e.Student)
}

type GradeBook struct {
	grades map[string]int
}

func NewGradeBook() *GradeBook {
	return &GradeBook{grades: make(map[string]int)}
}

func (gb *GradeBook) Get(name string) (int, error) {
	g, ok := gb.grades[name]
	if !ok {
		return 0, ErrNotFound
	}
	return g, nil
}

func (gb *GradeBook) Set(name string, grade int) error {
	if grade < 0 || grade > 100 {
		return &GradeError{Student: name, Grade: grade}
	}
	gb.grades[name] = grade
	return nil
}

func AddBonus(gb *GradeBook, name string, bonus int) error {
	current, err := gb.Get(name)
	if err != nil {
		return fmt.Errorf("add bonus to %q: %w", name, err)
	}
	if err := gb.Set(name, current+bonus); err != nil {
		return fmt.Errorf("add bonus to %q: %w", name, err)
	}
	return nil
}

func report(err error) {
	if err == nil {
		fmt.Println("OK")
		return
	}
	if errors.Is(err, ErrNotFound) {
		fmt.Println("404: student not found")
		return
	}
	var ge *GradeError
	if errors.As(err, &ge) {
		fmt.Printf("400: invalid grade %d for student %s\n", ge.Grade, ge.Student)
		return
	}
	fmt.Println("500:", err)
}
