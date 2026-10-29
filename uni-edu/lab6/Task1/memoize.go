package main

import "fmt"

func memoize(f func (int) int) func(int) int{
	cache := make(map[int]int)
	return func(n int) int{
		if val, exists := cache[n]; exists{
			return val
		}
		result := f(n)
		cache[n] = result
		return result
	}
}

func main(){
	calls := 0
	slowSquare := func(n int) int{
		calls++
		return n*n
	}
	fastSquare := memoize(slowSquare)
	
	fmt.Println(fastSquare(4), "(calls:", calls, ")")
	fmt.Println(fastSquare(4), "(calls:", calls, ")") 
	fmt.Println(fastSquare(5), "(calls:", calls, ")") 
}