package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("TrimRight(\"123oxo\", \"xo\")   = %q\n", strings.TrimRight("123oxo", "xo"))
	fmt.Printf("TrimSuffix(\"123oxo\", \"xo\")  = %q\n", strings.TrimSuffix("123oxo", "xo"))
	fmt.Printf("TrimLeft(\"xoxo123\", \"xo\")   = %q\n", strings.TrimLeft("xoxo123", "xo"))
	fmt.Printf("TrimPrefix(\"xoxo123\", \"xo\") = %q\n", strings.TrimPrefix("xoxo123", "xo"))
}
