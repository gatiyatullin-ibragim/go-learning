package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrBadFormat = errors.New("missing comma or bad format")

type Record struct {
	Name  string
	Grade int
}

func ParseAll(input string) ([]Record, error) {
	var records []Record
	var errs []error

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			errs = append(errs, fmt.Errorf("line %d: %w", i+1, ErrBadFormat))
			continue
		}

		name := strings.TrimSpace(parts[0])
		gradeStr := strings.TrimSpace(parts[1])

		grade, err := strconv.Atoi(gradeStr)
		if err != nil {
			errs = append(errs, fmt.Errorf("line %d: %w", i+1, err))
			continue
		}

		records = append(records, Record{Name: name, Grade: grade})
	}

	if len(errs) > 0 {
		return records, errors.Join(errs...)
	}
	return records, nil
}
