package join

import "testing"

var sinkS string
var ids10 = []int{1001, 1002, 1003, 1004, 1005, 1006, 1007, 1008, 1009, 1010}

func makeIDs(n int) []int {
	ids := make([]int, n)
	for i := range ids {
		ids[i] = 1000 + i
	}
	return ids
}

func TestSameResult(t *testing.T) {
	if a, b := JoinIDsSlow(ids10), JoinIDsFast(ids10); a != b {
		t.Fatalf("slow=%q fast=%q", a, b)
	}
}

func TestJoinIDsSlowAllocs(t *testing.T) {
	allocs := testing.AllocsPerRun(1000, func() {
		sinkS = JoinIDsSlow(ids10)
	})
	t.Logf("JoinIDsSlow: %v allocs", allocs)
}

func TestJoinIDsAllocs(t *testing.T) {
	allocs := testing.AllocsPerRun(1000, func() {
		sinkS = JoinIDsFast(ids10)
	})
	if allocs > 1 {
		t.Errorf("JoinIDsFast: %v allocs, want at most 1", allocs)
	}
}

func benchSlow(b *testing.B, n int) {
	ids := makeIDs(n)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkS = JoinIDsSlow(ids)
	}
}

func benchFast(b *testing.B, n int) {
	ids := makeIDs(n)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sinkS = JoinIDsFast(ids)
	}
}

func BenchmarkJoinIDsSlow(b *testing.B)     { benchSlow(b, 10) }
func BenchmarkJoinIDsFast(b *testing.B)     { benchFast(b, 10) }
func BenchmarkJoinIDsSlow1000(b *testing.B) { benchSlow(b, 1000) }
func BenchmarkJoinIDsFast1000(b *testing.B) { benchFast(b, 1000) }
