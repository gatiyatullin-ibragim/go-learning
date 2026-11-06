package main

import (
	"fmt"
	"os"
)

// 1. pad [64]byte
func depth64(n int) int {
	var pad [64]byte // делаем каждый фрейм побольше
	pad[n%len(pad)] = byte(n)
	if n == 0 {
		return 0
	}
	// читает 0, но компилятор этого не знает
	return depth64(n-1) + 1 + int(pad[(n+1)%len(pad)])
}

// 1. pad [1024]byte: фрейм ~1 КБ, 1_000_000 вызовов ~ 1 ГБ > лимита 1_000_000_000 байт
func depth1024(n int) int {
	var pad [1024]byte
	pad[n%len(pad)] = byte(n)
	if n == 0 {
		return 0
	}
	return depth1024(n-1) + 1 + int(pad[(n+1)%len(pad)])
}

// 2. Убрали строку, которая ЧИТАЕТ pad. Массив только пишется,
// компилятор делает dead store elimination и выбрасывает его: фрейм маленький.
func depthNoRead(n int) int {
	var pad [1024]byte
	pad[n%len(pad)] = byte(n)
	if n == 0 {
		return 0
	}
	return depthNoRead(n-1) + 1
}

// 3. Бесконечная рекурсия
func forever(n int) int {
	return forever(n+1) + 1
}

func main() {
	mode := "depth64"
	if len(os.Args) > 1 {
		mode = os.Args[1]
	}
	switch mode {
	case "depth64":
		fmt.Println("depth:", depth64(1_000_000))
	case "depth1024":
		fmt.Println("depth:", depth1024(1_000_000))
	case "nopad":
		fmt.Println("depth:", depthNoRead(1_000_000))
	case "forever":
		// recover НЕ сработает: stack overflow это fatal error, а не panic
		defer func() {
			fmt.Println("recovered:", recover()) // эта строка не напечатается
		}()
		fmt.Println(forever(0))
	default:
		fmt.Println("unknown mode:", mode)
	}
}
