func TestFetchNoLeak(t *testing.T) {
    before := runtime.NumGoroutine()
    for i := 0; i < 50; i++ {
        fetchLeaky(time.Millisecond) // после фикса заменить на fetchFixed
    }
    time.Sleep(200 * time.Millisecond) // дать горутинам завершиться
    after := runtime.NumGoroutine()
    if after > before {
        t.Fatalf("goroutine leak: before=%d after=%d", before, after)
    }
}