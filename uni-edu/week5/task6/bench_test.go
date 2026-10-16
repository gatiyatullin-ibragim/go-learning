// Task 6: Benchmark concrete vs generic vs any
// Run: go test -bench=. -benchmem ./task6
package task6

import "testing"

type Number interface {
	~int | ~int64 | ~float64
}

func SumInts(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func SumAny(nums []any) int {
	total := 0
	for _, n := range nums {
		total += n.(int) // type assertion on every element
	}
	return total
}

func ToAny(nums []int) []any {
	out := make([]any, len(nums))
	for i, n := range nums {
		out[i] = n // boxing: int -> any
	}
	return out
}

func makeNums() []int {
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i // first number is 0
	}
	return nums
}

var sink int

func BenchmarkSumInts(b *testing.B) {
	nums := makeNums()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = SumInts(nums)
	}
}

func BenchmarkSumGeneric(b *testing.B) {
	nums := makeNums()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = Sum(nums)
	}
}

func BenchmarkSumAny(b *testing.B) {
	anys := ToAny(makeNums()) // converted ONCE, before ResetTimer
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = SumAny(anys)
	}
}

func BenchmarkSumAnyWithBoxing(b *testing.B) {
	nums := makeNums()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sink = SumAny(ToAny(nums)) // boxing inside the loop
	}
}
