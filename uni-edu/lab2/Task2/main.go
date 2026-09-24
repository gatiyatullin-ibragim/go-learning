package main

import "fmt"

func main(){
	scores := map[string]int{
		"Alice" : 90,
		"Bob" : 0,
	} 
	

	//1)
	if value, ok := scores["Charlie"]; ok {
		fmt.Println("value: ", value, "Charlie is in the map")
	}else {
		fmt.Println("Charlie is missing")
	}

	_, ok := scores["Bob"]
	if ok {
		fmt.Println( "Bob is in the map")
	}else {
		fmt.Println("Bob is missing")
	}


	//2)

}
