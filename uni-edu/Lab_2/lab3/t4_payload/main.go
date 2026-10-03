package main

import "fmt"
import "bytes"
import "strings"

func main() {
	in := []byte("  hello\r\nworld\r\n  ")
	fmt.Printf("original:  %q\n", processPayload(in))
	fmt.Printf("optimized: %q\n", processPayloadOptimized(in))
}

func processPayload(data []byte) string {
	trimmed := string(bytes.TrimSpace(data))
	clean := strings.ReplaceAll(trimmed, "\r", "")
	return clean
}

// Оптимизированный: только bytes
func processPayloadOptimized(data []byte) []byte {
	trimmed := bytes.TrimSpace(data)
	return bytes.ReplaceAll(trimmed, []byte("\r"), nil)
}