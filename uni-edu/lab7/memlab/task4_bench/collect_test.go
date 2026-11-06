package bench

import "testing"

var (
	sinkP  Point
	sinkPP *Point
	sinkSl []int
)

func BenchmarkNewPointValue(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkP = NewPointValue(i, i)
	}
}

func BenchmarkNewPointPointer(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkPP = NewPointPointer(i, i)
	}
}

func BenchmarkCollectNoCap1000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkSl = CollectNoCap(1000)
	}
}

func BenchmarkCollectWithCap1000(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		sinkSl = CollectWithCap(1000)
	}
}

func TestAllocCounts(t *testing.T) {
	measure := func(name string, f func()) float64 {
		got := testing.AllocsPerRun(1000, f)
		t.Logf("%-22s %v allocs/op", name, got)
		return got
	}
	v := measure("NewPointValue", func() { sinkP = NewPointValue(1, 2) })
	p := measure("NewPointPointer", func() { sinkPP = NewPointPointer(1, 2) })
	no := measure("CollectNoCap(1000)", func() { sinkSl = CollectNoCap(1000) })
	with := measure("CollectWithCap(1000)", func() { sinkSl = CollectWithCap(1000) })

	if v != 0 {
		t.Errorf("NewPointValue: got %v, want 0", v)
	}
	if p != 1 {
		t.Errorf("NewPointPointer: got %v, want 1", p)
	}
	if with != 1 {
		t.Errorf("CollectWithCap: got %v, want 1", with)
	}
	if no <= with {
		t.Errorf("CollectNoCap (%v) must allocate more than CollectWithCap (%v)", no, with)
	}
}
