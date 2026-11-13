func stackSys() uint64 {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    return m.StackSys
}

func main() {
    const N = 100_000
    before := stackSys()
    stop := make(chan struct{})
    var wg sync.WaitGroup
    start := time.Now()
    for i := 0; i < N; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            <-stop
        }()
    }
    fmt.Println("started in", time.Since(start))
    time.Sleep(100 * time.Millisecond)
    after := stackSys()
    fmt.Println("goroutines alive:", runtime.NumGoroutine())
    fmt.Printf("stack memory: %d MB (%.1f KB per goroutine)\n",
        (after-before)>>20, float64(after-before)/N/1024)
    close(stop)
    wg.Wait()
    time.Sleep(100 * time.Millisecond)
    fmt.Println("after close:", runtime.NumGoroutine(), "goroutine(s)")
}