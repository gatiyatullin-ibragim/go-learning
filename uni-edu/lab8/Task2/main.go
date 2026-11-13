func say(msg string, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        fmt.Println(msg, i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go say("goroutine", &wg)
    go say("main", &wg)
    wg.Wait()
}