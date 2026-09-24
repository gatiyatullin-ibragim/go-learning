package main
import "fmt"



//methods to create maps
//1) m := make(map[string]int)
//2) m := make(map[string]int, 100)          optional size
//3) m := map[string]int{"a":1, "b":2}       map literal
//4) var m map[string]int                    zero value - nil, not empty



func main(){
	// first method
	var a = map[string]string{"Student1": "loh", "Student2": "loh2"}
	
	a["Student1"] = "ne loh" //change the value for Student1
	fmt.Println(a["Student1"])	
	delete(a,"Student1") //deletes the key and value

	fmt.Print(a)

}