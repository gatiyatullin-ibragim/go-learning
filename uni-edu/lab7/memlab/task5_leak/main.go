package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

func readFile(i int) []byte {
	data := make([]byte, 1<<20) 
	copy(data, fmt.Sprintf("HEADER-%04d", i))
	return data
}

func keepHeadersSlice(n int) [][]byte {
	var headers [][]byte
	for i := 0; i < n; i++ {
		data := readFile(i)
		headers = append(headers, data[:11])
	}
	return headers
}

func keepHeadersCopy(n int) [][]byte {
	var headers [][]byte
	for i := 0; i < n; i++ {
		data := readFile(i)
		headers = append(headers, bytes.Clone(data[:11]))
	}
	return headers
}

func makeBig(i int) string {
	return fmt.Sprintf("%05d", i) + strings.Repeat("x", 1<<20)
}

func keepPrefixSub(n int) []string {
	var out []string
	for i := 0; i < n; i++ {
		big := makeBig(i)
		out = append(out, big[:5])
	}
	return out
}

func keepPrefixClone(n int) []string {
	var out []string
	for i := 0; i < n; i++ {
		big := makeBig(i)
		out = append(out, strings.Clone(big[:5]))
	}
	return out
}

func liveHeapMB() int {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return int(m.HeapAlloc >> 20)
}

func main() {
	const n = 100

	base := liveHeapMB()
	h1 := keepHeadersSlice(n)
	fmt.Println("subslices: ", liveHeapMB()-base, "MB")
	runtime.KeepAlive(h1) 

	base = liveHeapMB()
	h2 := keepHeadersCopy(n)
	fmt.Println("copies:    ", liveHeapMB()-base, "MB")
	runtime.KeepAlive(h2)

	base = liveHeapMB()
	s1 := keepPrefixSub(n)
	fmt.Println("substrings:", liveHeapMB()-base, "MB")
	runtime.KeepAlive(s1)

	base = liveHeapMB()
	s2 := keepPrefixClone(n)
	fmt.Println("clones:    ", liveHeapMB()-base, "MB")
	runtime.KeepAlive(s2)
}
