package main

import (
	"fmt"
	"runtime"
	"time"
)

type Record struct {
	ID   int
	Data [256]byte
}

var live []*Record 

var sink *Record

func consume(r *Record) int {
	sink = r
	return r.ID + int(r.Data[0])
}

func main() {
	start := time.Now()

	for i := 0; i < 80_000; i++ {
		live = append(live, &Record{ID: i})
	}

	sum := 0
	for i := 0; i < 10_000_000; i++ {
		r := &Record{ID: i} 
		sum += consume(r)
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("sum:               ", sum)
	fmt.Println("GC cycles:         ", m.NumGC)
	fmt.Println("heap from OS (MB): ", m.HeapSys>>20)
	fmt.Printf("total pauses (ms):  %.1f\n", float64(m.PauseTotalNs)/1e6)
	fmt.Println("time:              ", time.Since(start).Round(time.Millisecond))
}
