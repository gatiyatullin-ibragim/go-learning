// Task 1: Channel semantics
//
//	go run .            — все три части задания
//	go run . cell N     — запустить только одну ячейку таблицы (N = 1..12)
//
// Ячейки, которые паникуют или зависают навсегда, убили бы весь процесс,
// поэтому main запускает каждую ячейку в отдельном дочернем процессе
// (тот же бинарник с аргументом "cell N") и классифицирует результат.
package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// ---------- 1.1 ping ----------

func ping(sleepInSender bool) {
	start := time.Now()
	since := func() time.Duration { return time.Since(start).Round(100 * time.Millisecond) }

	ch := make(chan string) // unbuffered
	go func() {
		if sleepInSender {
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("  sender:   about to send", since())
		ch <- "ping"
		fmt.Println("  sender:   send finished", since())
	}()
	if !sleepInSender {
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("  receiver: waiting...  ", since())
	msg := <-ch
	fmt.Println("  receiver: got", msg, "     ", since())
	time.Sleep(10 * time.Millisecond) // дать sender'у напечатать "send finished"
}

// ---------- 1.2 таблица nil / open / closed ----------

type cell struct {
	name string
	fn   func()
}

var cells = []cell{
	// nil channel
	{"nil    | ch <- v", func() { var ch chan int; ch <- 1 }},
	{"nil    | <-ch", func() { var ch chan int; <-ch }},
	{"nil    | close(ch)", func() { var ch chan int; close(ch) }},
	{"nil    | len/cap", func() { var ch chan int; fmt.Print("len=", len(ch), " cap=", cap(ch)) }},
	// open channel
	{"open   | ch <- v", func() {
		ch := make(chan int) // никто не принимает → блок навсегда
		ch <- 1
	}},
	{"open   | <-ch", func() {
		ch := make(chan int) // никто не шлёт → блок навсегда
		<-ch
	}},
	{"open   | close(ch)", func() { ch := make(chan int); close(ch); fmt.Print("ok") }},
	{"open   | len/cap", func() {
		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		fmt.Print("len=", len(ch), " cap=", cap(ch))
	}},
	// closed channel
	{"closed | ch <- v", func() { ch := make(chan int, 1); close(ch); ch <- 1 }},
	{"closed | <-ch", func() {
		ch := make(chan int, 2)
		ch <- 7
		close(ch)
		v1, ok1 := <-ch
		v2, ok2 := <-ch
		fmt.Printf("(%d,%v) then (%d,%v)", v1, ok1, v2, ok2)
	}},
	{"closed | close(ch)", func() { ch := make(chan int); close(ch); close(ch) }},
	{"closed | len/cap", func() {
		ch := make(chan int, 3)
		ch <- 1
		close(ch)
		fmt.Print("len=", len(ch), " cap=", cap(ch))
	}},
}

func runCellsInSubprocesses() {
	self, err := os.Executable()
	if err != nil {
		fmt.Println("cannot find own binary:", err)
		return
	}
	for i, c := range cells {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		cmd := exec.CommandContext(ctx, self, "cell", strconv.Itoa(i+1))
		var out bytes.Buffer
		cmd.Stdout, cmd.Stderr = &out, &out
		_ = cmd.Run()
		cancel()

		text := out.String()
		var result string
		switch {
		case strings.Contains(text, "fatal error: all goroutines are asleep - deadlock!"):
			result = "FATAL deadlock  (" + firstLineWith(text, "goroutine 1 [") + ")"
		case strings.Contains(text, "panic:"):
			result = "PANIC  " + firstLineWith(text, "panic:")
		default:
			result = "OK     " + strings.TrimSpace(text)
		}
		fmt.Printf("  %2d. %-20s → %s\n", i+1, c.name, result)
	}
}

func firstLineWith(text, prefix string) string {
	for _, l := range strings.Split(text, "\n") {
		if strings.HasPrefix(l, prefix) {
			return strings.TrimSuffix(l, ":")
		}
	}
	return ""
}

// ---------- 1.3 say("goroutine") с каналом вместо WaitGroup ----------

func say(msg string, done chan<- struct{}) {
	for i := 0; i < 3; i++ {
		fmt.Println(" ", msg, i)
		time.Sleep(100 * time.Millisecond)
	}
	done <- struct{}{} // "я закончил" — main ждёт этого значения
}

func main() {
	if len(os.Args) == 3 && os.Args[1] == "cell" {
		n, _ := strconv.Atoi(os.Args[2])
		cells[n-1].fn()
		return
	}

	fmt.Println("=== 1.1 ping: Sleep в sender ===")
	ping(true)
	fmt.Println("=== 1.1 ping: Sleep перенесён в receiver ===")
	ping(false)

	fmt.Println("\n=== 1.2 таблица nil / open / closed (каждая ячейка — отдельный процесс) ===")
	runCellsInSubprocesses()

	fmt.Println("\n=== 1.3 say с каналом вместо WaitGroup ===")
	done := make(chan struct{})
	go say("goroutine", done)
	say("main", make(chan struct{}, 1)) // буфер 1: свой сигнал main никто не читает
	<-done
	fmt.Println("  main: goroutine finished")
}
