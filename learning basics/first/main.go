//Обработка ошибок в Go 
package main
import "fmt"
//Первый способ

func div(a,b int)(int, error){
	if b == 0{
		return 0, fmt.Errorf("divisor is equal to 0")
	}
	return a/b, nil
}

func main(){
	d, err := div(10,0)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Printf("d = %d", d)
	}
}