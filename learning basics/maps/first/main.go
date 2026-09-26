package main

import "fmt"

func main(){
	mp := map[string]int{
		"alice": 60,
		"Bakytzhan": 100,
	}

	fmt.Println(mp)

	if value, ok := mp["alice"]; ok{
		fmt.Println(value)
	}else{
		fmt.Println("NOnono")
	}


}
