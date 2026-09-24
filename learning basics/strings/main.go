package main
import "unicode/utf8"

func main(){
	
	var str string
	var stre string
	str = "abc"
	stre = "абц"
	
	println(str) //prints the string itself
	println(str[2]) //prints the second letter's value, because GO supports unicode
	println(len(str)) //prints lenght of string, but it will be correct only if all letters are from ASCII and their size is under 1 Bite
	println(utf8.RuneCountInString(str)) // Rune it's a special type that can contain more than 1 Byte. So we can get the exact number of symbols 
	println(utf8.RuneCountInString(stre)) 
}	