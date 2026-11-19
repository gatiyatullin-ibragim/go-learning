// Task 7: Benchmark — how expensive is a channel?
//
//	go run .
//	go test -v .
//	go test -bench . -benchmem -cpu 1,2 .
package main

import "fmt"

const N = 10_000

const batchSize = 100

func passN(size int) int {
	ch := make(chan int, size)
	go func() {
		for i := 0; i < N; i++ {
			ch <- i
		}
		close(ch)
	}()
	sum := 0
	for v := range ch {
		sum += v
	}
	return sum
}

// passBatches отправляет те же N чисел пачками []int по 100 штук:
// 100 операций с каналом вместо 10 000.
func passBatches() int {
	ch := make(chan []int, 1)
	go func() {
		for start := 0; start < N; start += batchSize {
			batch := make([]int, 0, batchSize)
			for i := start; i < start+batchSize && i < N; i++ {
				batch = append(batch, i)
			}
			ch <- batch
		}
		close(ch)
	}()
	sum := 0
	for batch := range ch {
		for _, v := range batch {
			sum += v
		}
	}
	return sum
}

func main() {
	fmt.Println("want:", N*(N-1)/2)
	for _, s := range []int{0, 1, 100, 10000} {
		fmt.Printf("passN(%d) = %d\n", s, passN(s))
	}
	fmt.Println("passBatches() =", passBatches())
}
