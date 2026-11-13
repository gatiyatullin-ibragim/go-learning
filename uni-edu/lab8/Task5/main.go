func slowQuery() int { time.Sleep(50 * time.Millisecond); return 42 }

// с утечкой
func fetchLeaky(timeout time.Duration) (int, error) {
    ch := make(chan int)
    go func() { ch <- slowQuery() }()
    select {
    case v := <-ch:
        return v, nil
    case <-time.After(timeout):
        return 0, errors.New("timeout")
    }
}

// исправлено: буфер 1
func fetchFixed(timeout time.Duration) (int, error) {
    ch := make(chan int, 1)
    go func() { ch <- slowQuery() }()
    select {
    case v := <-ch:
        return v, nil
    case <-time.After(timeout):
        return 0, errors.New("timeout")
    }
}

func startTickerLeaky() {
    go func() {
        for { time.Sleep(10 * time.Millisecond) }
    }()
}

func startTicker(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            case <-time.After(10 * time.Millisecond):
            }
        }
    }()
}

func main() {
    fmt.Println("start:", runtime.NumGoroutine())
    for i := 0; i < 100; i++ { fetchLeaky(time.Millisecond) }
    fmt.Println("after 100 fetchLeaky:", runtime.NumGoroutine())
    for i := 0; i < 100; i++ { fetchFixed(time.Millisecond) }
    time.Sleep(100 * time.Millisecond)
    fmt.Println("after 100 fetchFixed:", runtime.NumGoroutine()) // не растёт
    for i := 0; i < 10; i++ { startTickerLeaky() }
    fmt.Println("after 10 leaky tickers:", runtime.NumGoroutine())
    ctx, cancel := context.WithCancel(context.Background())
    for i := 0; i < 10; i++ { startTicker(ctx) }
    fmt.Println("after 10 stoppable tickers:", runtime.NumGoroutine())
    cancel()
    time.Sleep(50 * time.Millisecond)
    fmt.Println("after cancel():", runtime.NumGoroutine()) // -10
    pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
    pprof.Lookup("goroutine").WriteTo(os.Stdout, 2)
}