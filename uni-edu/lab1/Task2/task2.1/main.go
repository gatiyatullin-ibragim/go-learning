package main

import (
	"fmt"
	"unsafe"
)

type Slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	cars := []string{"Ferrari", "Honda", "Ford", "BYD"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
}

//cars: [Ferrari Honda Ford BYD] has old length 4 and capacity 4
//cars: [Ferrari Honda Ford BYD Toyota] has new length 5 and capacity 8

