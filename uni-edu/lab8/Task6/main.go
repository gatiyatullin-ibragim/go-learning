type Report struct{ data [64 << 10]byte }

func buildReport() *Report {
    time.Sleep(20 * time.Millisecond)
    return &Report{}
}

func handle(timeout time.Duration) error {
    ch := make(chan *Report) // утечка: небуферизованный
    go func() { ch <- buildReport() }()
    select {
    case r := <-ch:
        _ = r
        return nil
    case <-time.After(timeout):
        return errors.New("timeout")
    }
}

func stats(label string) {
    runtime.GC()
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("%s goroutines=%d heap=%d MB\n", label, runtime.NumGoroutine(), m.HeapAlloc>>20)
}

func main() {
    stats("start")
    for round := 1; round <= 3; round++ {
        var wg sync.WaitGroup
        for i := 0; i < 1000; i++ {
            wg.Add(1)
            go func() { defer wg.Done(); handle(5 * time.Millisecond) }()
        }
        wg.Wait()
        stats(fmt.Sprintf("after round %d", round))
    }
}