const N = 200_000
var sink int

func isPrime(n int) bool {
    if n < 2 { return false }
    for d := 2; d*d <= n; d++ {
        if n%d == 0 { return false }
    }
    return true
}

func countRange(from, to int) int {
    c := 0
    for n := from; n < to; n++ {
        if isPrime(n) { c++ }
    }
    return c
}

func CountSeq(n int) int { return countRange(0, n) }

func CountPar(n, parts int) int {
    if parts < 1 { parts = 1 }
    results := make([]int, parts)
    var wg sync.WaitGroup
    chunk := n / parts
    for p := 0; p < parts; p++ {
        from := p * chunk
        to := from + chunk
        if p == parts-1 { to = n } // последний забирает остаток
        wg.Add(1)
        go func() {
            defer wg.Done()
            results[p] = countRange(from, to)
        }()
    }
    wg.Wait()
    total := 0
    for _, r := range results { total += r }
    return total
}