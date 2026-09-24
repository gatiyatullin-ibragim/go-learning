package main
import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int{
	mp := make(map[string]int)
	for _, values:= range(strings.Fields(s)){
		mp[values]++
	}
	return mp
}

func main(){
	str := "i want five out of five points"

	fmt.Print(wordCount(str))
}
