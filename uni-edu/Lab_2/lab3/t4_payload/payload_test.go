package main

import (
	"strings"
	"testing"
)

var payload = []byte("  " + strings.Repeat("some line of text\r\n", 200) + "  ")

var (
	sinkString string
	sinkBytes  []byte
)

func TestSameResult(t *testing.T) {
	if processPayload(payload) != string(processPayloadOptimized(payload)) {
		t.Fatal("results differ")
	}
}

func BenchmarkProcessPayload(b *testing.B) {
	b.Run("original", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkString = processPayload(payload)
		}
	})
	b.Run("optimized", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			sinkBytes = processPayloadOptimized(payload)
		}
	})
}
