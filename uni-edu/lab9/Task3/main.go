// Task 3: Ownership
//
//	go run .          — collect (правильная версия) + пайплайны
//	go run . wrong    — WRONG-версия: каждый producer закрывает общий канал
//	go run -race .    — проверка гонок
package main

import (
	"fmt"
	"os"
	"sync"
)

// ---------- 3.1 WRONG ----------

func wrong() {
	results := make(chan int)
	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				results <- w*10 + i
			}
			close(results) // "I'm done" — but the others are not!
		}()
	}
	sum := 0
	for v := range results {
		sum += v
	}
	fmt.Println("sum:", sum)
	wg.Wait()
}

// ---------- 3.1 collect с closer-горутиной ----------

func collect(workers int) <-chan int {
	results := make(chan int)
	var wg sync.WaitGroup
	for w := 1; w <= workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				results <- w*10 + i
			}
		}()
	}
	go func() {
		wg.Wait()      // all writers are finished...
		close(results) // ...so now it's safe to close
	}()
	return results
}

// ---------- 3.2 / 3.3 пайплайн ----------

func generator(from, to int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := from; n <= to; n++ {
			out <- n
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v
		}
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if v%2 == 0 {
				out <- v
			}
		}
	}()
	return out
}

// sum — последняя стадия: она ничего не отдаёт дальше, поэтому своего
// выходного канала у неё нет и закрывать нечего.
func sum(in <-chan int) int {
	s := 0
	for v := range in {
		s += v
	}
	return s
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "wrong" {
		wrong()
		return
	}

	s := 0
	for v := range collect(3) {
		s += v
	}
	fmt.Println("collect(3) sum:", s, "(want 189)")

	fmt.Println("generator→square→sum:           ", sum(square(generator(1, 10))), "(want 385)")
	fmt.Println("generator→square→filterEven→sum:", sum(filterEven(square(generator(1, 10)))), "(want 220)")
}
