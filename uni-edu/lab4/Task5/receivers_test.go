package receivers

import "testing"

var sink float64

func BenchmarkFirstValue(b *testing.B) {
	var r Report
	for i := 0; i < b.N; i++ {
		sink = r.FirstValue() 
	}
}

func BenchmarkFirstPointer(b *testing.B) {
	var r Report
	for i := 0; i < b.N; i++ {
		sink = r.FirstPointer() 
	}
}

func BenchmarkPointValue(b *testing.B) {
	p := Point{1, 2}
	for i := 0; i < b.N; i++ {
		sink = p.SumValue()
	}
}

func BenchmarkPointPointer(b *testing.B) {
	p := Point{1, 2}
	for i := 0; i < b.N; i++ {
		sink = p.SumPointer()
	}
}