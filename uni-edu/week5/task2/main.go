// Task 2: Describe anything (+ Participation: version without type switch)
package main

import (
	"errors"
	"fmt"
)

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%.1f°C", float64(c)) }

func describe(v any) string {
	switch x := v.(type) {
	case nil:
		return "nil"
	case int, int64:
		return fmt.Sprintf("integer %v", x)
	case string:
		return fmt.Sprintf("string of %d bytes", len(x))
	case []int:
		return fmt.Sprintf("slice with %d ints", len(x))
	case bool: // step 3
		return fmt.Sprintf("bool %t", x)
	case map[string]int: // step 3
		return fmt.Sprintf("map with %d keys", len(x))
	case error:
		return "error: " + x.Error()
	case fmt.Stringer:
		return "Stringer: " + x.String()
	default:
		return fmt.Sprintf("unknown type %T", x)
	}
}

// Participation: the same logic with only v, ok := i.(T)
func describeNoSwitch(v any) string {
	if v == nil {
		return "nil"
	}
	if n, ok := v.(int); ok {
		return fmt.Sprintf("integer %v", n)
	}
	if n, ok := v.(int64); ok {
		return fmt.Sprintf("integer %v", n)
	}
	if s, ok := v.(string); ok {
		return fmt.Sprintf("string of %d bytes", len(s))
	}
	if s, ok := v.([]int); ok {
		return fmt.Sprintf("slice with %d ints", len(s))
	}
	if b, ok := v.(bool); ok {
		return fmt.Sprintf("bool %t", b)
	}
	if m, ok := v.(map[string]int); ok {
		return fmt.Sprintf("map with %d keys", len(m))
	}
	if e, ok := v.(error); ok {
		return "error: " + e.Error()
	}
	if s, ok := v.(fmt.Stringer); ok {
		return "Stringer: " + s.String()
	}
	return fmt.Sprintf("unknown type %T", v)
}

// Question from the lab: order of cases matters
type Timeout struct{ Sec int }

func (t Timeout) Error() string  { return fmt.Sprintf("timeout after %ds", t.Sec) }
func (t Timeout) String() string { return "Timeout{" + fmt.Sprint(t.Sec) + "}" }

func kind(v any) string {
	switch v.(type) {
	case fmt.Stringer:
		return "Stringer"
	case error:
		return "error"
	case Timeout:
		return "Timeout"
	}
	return "other"
}

func main() {
	values := []any{
		nil, 7, int64(7), "привет", []int{1, 2, 3},
		errors.New("disk full"), Celsius(21.5), 3.14, struct{}{},
		true, map[string]int{"a": 1},
	}
	for _, v := range values {
		a, b := describe(v), describeNoSwitch(v)
		mark := "ok"
		if a != b {
			mark = "MISMATCH"
		}
		fmt.Printf("%-22s %-18T -> %-24s [%s]\n", fmt.Sprint(v), v, a, mark)
	}

	fmt.Println("\n--- Timeout question ---")
	fmt.Println(kind(Timeout{5})) // Stringer (first matching case wins)
	var e error = Timeout{3}
	fmt.Println(e) // timeout after 3s: fmt prefers Error() over String()
}
