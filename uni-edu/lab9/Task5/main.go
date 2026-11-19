// Task 5: select
//
//	go run .         — все пункты
//	go run -race .   — проверка гонок (в т.ч. merge)
package main

import (
	"fmt"
	"runtime"
	"time"
)

// ---------- 5.1 random choice / priority ----------

func filled(name string, n int) chan string {
	ch := make(chan string, n)
	for i := 0; i < n; i++ {
		ch <- name
	}
	return ch
}

func randomChoice() map[string]int {
	a, b := filled("a", 1000), filled("b", 1000)
	count := map[string]int{}
	for i := 0; i < 1000; i++ {
		select {
		case v := <-a:
			count[v]++
		case v := <-b:
			count[v]++
		}
	}
	return count
}

// preferA: первый select с default пробует только a; если a не готов —
// второй select ждёт любой из двух каналов.
func preferA() map[string]int {
	a, b := filled("a", 1000), filled("b", 1000)
	count := map[string]int{}
	for i := 0; i < 1000; i++ {
		select {
		case v := <-a:
			count[v]++
			continue
		default:
		}
		select {
		case v := <-a:
			count[v]++
		case v := <-b:
			count[v]++
		}
	}
	return count
}

// ---------- 5.2 fetchWithTimeout без утечки ----------

func slow() int {
	time.Sleep(200 * time.Millisecond)
	return 42
}

func fetchWithTimeout(timeout time.Duration) (int, error) {
	ch := make(chan int, 1) // буфер 1: горутина отправит и завершится, даже если её никто не ждёт
	go func() { ch <- slow() }()
	select {
	case v := <-ch:
		return v, nil
	case <-time.After(timeout):
		return 0, fmt.Errorf("timeout after %v", timeout)
	}
}

// ---------- 5.3 merge ----------

func source(vals []int, delay time.Duration) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range vals {
			time.Sleep(delay)
			out <- v
		}
	}()
	return out
}

// merge reads from a and b until BOTH are closed.
func merge(a, b <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for a != nil || b != nil {
			select {
			case v, ok := <-a:
				if !ok {
					a = nil // a nil channel blocks forever: this case is now off
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				out <- v
			}
		}
	}()
	return out
}

// mergeNoNil — тот же merge без строки a = nil (и b = nil).
// Закрытый канал всегда «готов» и мгновенно отдаёт (0, false), поэтому
// select крутится вхолостую. Считаем такие пустые приёмы.
func mergeNoNil(a, b <-chan int) (<-chan int, *int) {
	out := make(chan int)
	useless := new(int)
	go func() {
		defer close(out)
		aClosed, bClosed := false, false
		for !aClosed || !bClosed {
			select {
			case v, ok := <-a:
				if !ok {
					aClosed = true
					*useless++
					continue
				}
				out <- v
			case v, ok := <-b:
				if !ok {
					bClosed = true
					continue
				}
				out <- v
			}
		}
		*useless-- // первое (0,false) от a — это не лишний, а сигнал закрытия
	}()
	return out, useless
}

// ---------- Participation: ping-pong ----------

func player(table chan int, done <-chan struct{}) {
	for {
		var ball int
		select {
		case ball = <-table:
		case <-done:
			return
		}
		ball++
		select {
		case table <- ball:
		case <-done:
			return
		}
	}
}

func main() {
	fmt.Println("=== 5.1 random choice (3 запуска) ===")
	for i := 1; i <= 3; i++ {
		fmt.Printf("  run %d: %v\n", i, randomChoice())
	}
	fmt.Println("  preferA:", preferA())

	fmt.Println("\n=== 5.2 fetchWithTimeout ===")
	before := runtime.NumGoroutine()
	fmt.Println(fetchWithTimeout(500 * time.Millisecond))
	fmt.Println(fetchWithTimeout(100 * time.Millisecond))
	fmt.Println("  goroutines right after timeout:", runtime.NumGoroutine())
	time.Sleep(200 * time.Millisecond)
	fmt.Println("  goroutines 200ms later:        ", runtime.NumGoroutine(), "(before:", before, ")")

	fmt.Println("\n=== 5.3 merge ===")
	a := source([]int{1, 2, 3}, 10*time.Millisecond)
	b := source([]int{100, 101, 102, 103, 104}, 7*time.Millisecond)
	fmt.Print("  ")
	for v := range merge(a, b) {
		fmt.Print(v, " ")
	}
	fmt.Println()

	for i := 1; i <= 3; i++ {
		a = source([]int{1}, 0)                                 // a закрывается почти сразу
		b = source([]int{100, 101, 102}, 3*time.Millisecond)     // b ещё ~9ms шлёт значения
		out, useless := mergeNoNil(a, b)
		for range out {
		}
		fmt.Printf("  without a = nil, run %d: %d useless zero values from a\n", i, *useless)
	}

	fmt.Println("\n=== Participation: ping-pong ===")
	before = runtime.NumGoroutine()
	table := make(chan int)
	done := make(chan struct{})
	go player(table, done)
	go player(table, done)
	table <- 0 // подаём мяч
	time.Sleep(100 * time.Millisecond)
	ball := <-table // забираем мяч
	close(done)     // останавливаем обоих игроков
	time.Sleep(10 * time.Millisecond)
	fmt.Println("  hits in 100ms:", ball)
	fmt.Println("  goroutines before:", before, "after:", runtime.NumGoroutine())
}
