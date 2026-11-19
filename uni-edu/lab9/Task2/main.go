// Task 2: Buffered vs unbuffered
package main

import (
	"fmt"
	"time"
)

func run(name string, ch chan int) {
	start := time.Now()
	since := func() time.Duration { return time.Since(start).Round(100 * time.Millisecond) }
	done := make(chan struct{})
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Printf("%-12s sent %d at %v\n", name+":", i, since())
		}
		close(done)
	}()
	time.Sleep(300 * time.Millisecond) // the receiver is busy at first
	for i := 0; i < 3; i++ {
		<-ch
		time.Sleep(100 * time.Millisecond) // and slow
	}
	<-done
}

func main() {
	fmt.Println("=== 2.1 run() с разной ёмкостью ===")
	for _, c := range []int{0, 1, 2, 3, 10} {
		run(fmt.Sprintf("cap=%d", c), make(chan int, c))
		fmt.Println()
	}

	fmt.Println("=== 2.2 буфер 2 ===")
	ch := make(chan string, 2)
	ch <- "a"
	ch <- "b"
	fmt.Println(<-ch, <-ch)

	fmt.Println("=== 2.2 буфер 1: см. buf1() — вызов закомментирован, т.к. это deadlock ===")
	// buf1()
}

// buf1 — тот же код с make(chan string, 1).
// Второй send ch <- "b" блокируется: буфер полон, а принимать некому —
// приём стоит ниже в той же (единственной) горутине main.
// Результат: fatal error: all goroutines are asleep - deadlock!
func buf1() {
	ch := make(chan string, 1)
	ch <- "a"
	ch <- "b"
	fmt.Println(<-ch, <-ch)
}

var _ = buf1
