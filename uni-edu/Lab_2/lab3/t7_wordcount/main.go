package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}

// Локальная копия wc.Test из golang.org/x/tour/wc (пакет недоступен без сети)
func Test(f func(string) map[string]int) {
	ok := true
	for _, c := range []struct {
		in  string
		out map[string]int
	}{
		{"I am learning Go!", map[string]int{"I": 1, "am": 1, "learning": 1, "Go!": 1}},
		{"The quick brown fox jumped over the lazy dog.", map[string]int{"The": 1, "quick": 1, "brown": 1, "fox": 1, "jumped": 1, "over": 1, "the": 1, "lazy": 1, "dog.": 1}},
		{"I ate a donut. Then I ate another donut.", map[string]int{"I": 2, "ate": 2, "a": 1, "donut.": 2, "Then": 1, "another": 1}},
		{"A man a plan a canal panama.", map[string]int{"A": 1, "man": 1, "a": 2, "plan": 1, "canal": 1, "panama.": 1}},
	} {
		got := f(c.in)
		if fmt.Sprint(got) != fmt.Sprint(c.out) {
			ok = false
			fmt.Printf("FAIL\n input: %q\n want: %v\n got:  %v\n", c.in, c.out, got)
		} else {
			fmt.Printf("PASS\n input: %q\n output: %v\n", c.in, got)
		}
	}
	if ok {
		fmt.Println("All tests passed")
	}
}

func main() { Test(WordCount) }
