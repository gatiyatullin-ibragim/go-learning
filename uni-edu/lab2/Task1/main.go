package main

import (
	"fmt"
	"strings"
)


func main(){
	var b strings.Builder
	b.Grow(90)
	b.WriteString("user list\n")
	for i:=0; i<10; i++{
		fmt.Fprintf(&b, "user #%d\n", i)
	}
	fmt.Println(b.String())
}
