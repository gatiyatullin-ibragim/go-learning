package pool

import "testing"

var sinkN int

func makeIDs() []int {
	ids := make([]int, 200)
	for i := range ids {
		ids[i] = i
	}
	return ids
}

func BenchmarkRenderNew(b *testing.B) {
	ids := makeIDs()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkN = RenderNew(ids)
	}
}

func BenchmarkRenderPool(b *testing.B) {
	ids := makeIDs()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkN = RenderPool(ids)
	}
}

func TestSameLen(t *testing.T) {
	ids := makeIDs()
	if a, b := RenderNew(ids), RenderPool(ids); a != b {
		t.Fatalf("new=%d pool=%d", a, b)
	}
}

func TestNoResetBreaks(t *testing.T) {
	ids := makeIDs()
	first := RenderPoolNoReset(ids)
	second := RenderPoolNoReset(ids)
	t.Logf("without Reset: first=%d second=%d", first, second)
	if second <= first {
		t.Log("(пул мог быть очищен GC между вызовами; повтори)")
	}
}

func TestPoolAllocs(t *testing.T) {
	ids := makeIDs()
	RenderPool(ids)
	allocs := testing.AllocsPerRun(1000, func() { sinkN = RenderPool(ids) })
	t.Logf("RenderPool: %v allocs/op", allocs)
	if allocs > 0 {
		t.Errorf("RenderPool: %v allocs, want 0", allocs)
	}
}
