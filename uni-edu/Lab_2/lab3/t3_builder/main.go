package main

import (
	"fmt"
	"strings"
	"testing"
)

func concatPlus(values []string) string {
	s := ""
	for _, v := range values {
		s += v
	}
	return s
}

func concatBuilder(values []string) string {
	var sb strings.Builder
	for _, v := range values {
		sb.WriteString(v)
	}
	return sb.String()
}

func concatGrow(values []string) string {
	total := 0
	for _, v := range values {
		total += len(v)
	}
	var sb strings.Builder
	sb.Grow(total)
	for _, v := range values {
		sb.WriteString(v)
	}
	return sb.String()
}

func main() {
	values := make([]string, 1000)
	for i := range values {
		values[i] = "item-" + fmt.Sprint(i) + ";"
	}
	fmt.Println("результаты совпадают:",
		concatPlus(values) == concatBuilder(values) && concatBuilder(values) == concatGrow(values))

	for _, c := range []struct {
		name string
		f    func([]string) string
	}{{"+=", concatPlus}, {"Builder", concatBuilder}, {"Builder+Grow", concatGrow}} {
		r := testing.Benchmark(func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				c.f(values)
			}
		})
		fmt.Printf("%-13s %10d ns/op %8d B/op %5d allocs/op\n",
			c.name, r.NsPerOp(), r.AllocedBytesPerOp(), r.AllocsPerOp())
	}
}
