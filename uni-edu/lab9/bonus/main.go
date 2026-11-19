// Bonus: select на закрытом и на nil канале
package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. закрытый канал: приём из него всегда готов → всегда "closed: 0 false"
	ch := make(chan int)
	close(ch)
	for i := 0; i < 5; i++ {
		select {
		case v, ok := <-ch:
			fmt.Println("closed:", v, ok)
		default:
			fmt.Println("default")
		}
	}

	// 2. nil канал: case никогда не готов → всегда "default"
	var nilCh chan int
	select {
	case v, ok := <-nilCh:
		fmt.Println("nil:", v, ok)
	default:
		fmt.Println("nil version: default")
	}

	// 3. nil без default: select блокируется навсегда →
	//    fatal error: all goroutines are asleep - deadlock!
	if len(os.Args) > 1 && os.Args[1] == "deadlock" {
		select {
		case v, ok := <-nilCh:
			fmt.Println("nil:", v, ok)
		}
	}
}
