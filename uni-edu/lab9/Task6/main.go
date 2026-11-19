// Task 6: A worker pool with backpressure and cancellation
package main

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Result struct {
	Word  string
	Upper string
}

var words = []string{"alpha", "bravo", "charlie", "delta", "echo", "foxtrot",
	"golf", "hotel", "india", "juliet", "kilo", "lima", "mike", "november"}

// ---------- версия с отменой через done ----------

func producer(words []string, done <-chan struct{}) <-chan string {
	jobs := make(chan string)
	go func() {
		defer close(jobs)
		for _, w := range words {
			select {
			case jobs <- w:
			case <-done: // the consumer gave up: stop producing
				return
			}
		}
	}()
	return jobs
}

func startWorkers(n int, jobs <-chan string, done <-chan struct{}) <-chan Result {
	results := make(chan Result)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for w := range jobs {
				time.Sleep(5 * time.Millisecond) // "работа"
				select {
				case results <- Result{w, strings.ToUpper(w)}:
				case <-done:
					return
				}
			}
		}()
	}
	go func() { // closer: единственный, кто закрывает results
		wg.Wait()
		close(results)
	}()
	return results
}

// ---------- та же версия, но без case <-done ----------

func producerNoDone(words []string) <-chan string {
	jobs := make(chan string)
	go func() {
		defer close(jobs)
		for _, w := range words {
			jobs <- w
		}
	}()
	return jobs
}

func startWorkersNoDone(n int, jobs <-chan string) <-chan Result {
	results := make(chan Result)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for w := range jobs {
				time.Sleep(5 * time.Millisecond)
				results <- Result{w, strings.ToUpper(w)}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	return results
}

// ---------- Submit ----------

var ErrQueueFull = errors.New("queue is full")

func Submit(queue chan<- string, job string, timeout time.Duration) error {
	t := time.NewTimer(timeout)
	defer t.Stop()
	select {
	case queue <- job:
		return nil
	case <-t.C:
		return ErrQueueFull // an HTTP server would answer 503 "try again later"
	}
}

func settle() int {
	time.Sleep(50 * time.Millisecond)
	return runtime.NumGoroutine()
}

func main() {
	const workers = 4

	fmt.Println("=== run to the end with range ===")
	done := make(chan struct{})
	n := 0
	for r := range startWorkers(workers, producer(words, done), done) {
		fmt.Printf("  %-9s → %s\n", r.Word, r.Upper)
		n++
	}
	close(done)
	fmt.Println("  results:", n, " goroutines:", settle())

	fmt.Println("\n=== take only 3, then close(done) ===")
	done = make(chan struct{})
	results := startWorkers(workers, producer(words, done), done)
	for i := 0; i < 3; i++ {
		r := <-results
		fmt.Printf("  %-9s → %s\n", r.Word, r.Upper)
	}
	close(done)
	fmt.Println("  goroutines after close(done):", settle(), "(must be 1)")

	fmt.Println("\n=== without <-done cases: take only 3 ===")
	results = startWorkersNoDone(workers, producerNoDone(words))
	for i := 0; i < 3; i++ {
		<-results
	}
	g := settle()
	fmt.Println("  goroutines:", g, "→ leaked:", g-1,
		fmt.Sprintf("(%d workers + producer + closer)", workers))

	fmt.Println("\n=== Submit ===")
	queue := make(chan string, 1)
	fmt.Println("  Submit #1:", Submit(queue, "job-1", 50*time.Millisecond)) // в буфере есть место
	fmt.Println("  Submit #2:", Submit(queue, "job-2", 50*time.Millisecond)) // буфер полон, читателя нет
}
