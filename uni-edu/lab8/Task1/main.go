package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func isPrime(n int) bool {
    if n < 2 { return false }
    for d := 2; d*d <= n; d++ {
        if n%d == 0 { return false }
    }
    return true
}

func countPrimes(from, to int) int {
    c := 0
    for n := from; n < to; n++ {
        if isPrime(n) { c++ }
    }
    return c
}

func cpuWork(workers int) (int, time.Duration) {
    const N = 6_000_000
    start := time.Now()
    results := make([]int, workers)
    var wg sync.WaitGroup
    chunk := N / workers
    for w := 0; w < workers; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            to := (w + 1) * chunk
            if w == workers-1 { to = N }
            results[w] = countPrimes(w*chunk, to)
        }()
    }
    wg.Wait()
    total := 0
    for _, r := range results { total += r }
    return total, time.Since(start)
}

func ioWork(workers int) time.Duration {
    start := time.Now()
    var wg sync.WaitGroup
    for w := 0; w < workers; w++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            time.Sleep(100 * time.Millisecond)
        }()
    }
    wg.Wait()
    return time.Since(start)
}

func main() {
    fmt.Println("NumCPU:", runtime.NumCPU(), "GOMAXPROCS:", runtime.GOMAXPROCS(0))
    for _, w := range []int{1, 2, 4, runtime.NumCPU(), 2 * runtime.NumCPU()} {
        p, d := cpuWork(w)
        fmt.Printf("CPU work, %d goroutine(s): %d primes in %v\n", w, p, d)
    }
    for _, w := range []int{1, 10, 1000} {
        fmt.Printf("I/O work, %d goroutine(s): %v\n", w, ioWork(w))
    }
}