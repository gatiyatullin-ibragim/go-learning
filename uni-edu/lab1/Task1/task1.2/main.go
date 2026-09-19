package main

import "fmt"

func main(){
	var arr = [5]int{1,2,3}
	fmt.Println(arr)
}

//output: [1 2 3 0 0]
//free spaces will be filled up with the zeroes 