package main
import "fmt"

func makeAccumulator() func()int{
	sum := 0
	return func(n int){
		sum += n
		return sum
	}
}

func main(){
	acc := makeAccumulator()
	fmt.Println(acc(10))
	fmt.Println(acc(5))
	fmt.Println(acc(4))
}