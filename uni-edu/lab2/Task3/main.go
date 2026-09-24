package main
import "fmt"


func main(){
	inventory := map[string]int{"apples": 10, "bananas":5}
	fmt.Println(inventory)
	fmt.Println(inventory["apples"])
	inventory["bananas"] = 12
	inventory["oranges"] = 8
	delete(inventory, "apples")
	fmt.Println(inventory)
}