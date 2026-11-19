package main

import (
	"fmt"
	"testing"
)

var sizes = []int{0, 1, 100, 10000}

func TestPassN(t *testing.T) {
	want := N * (N - 1) / 2
	for _, s := range sizes {
		if got := passN(s); got != want {
			t.Errorf("passN(%d) = %d, want %d", s, got, want)
		}
	}
	if got := passBatches(); got != want {
		t.Errorf("passBatches() = %d, want %d", got, want)
	}
}

var sink int

func BenchmarkPassN(b *testing.B) {
	for _, s := range sizes {
		b.Run(fmt.Sprintf("buf=%d", s), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = passN(s)
			}
		})
	}
}

func BenchmarkPassBatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sink = passBatches()
	}
}
