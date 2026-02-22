package main 

import (
	"fmt"
)


func main(){

	str:="Hello World"

	var ch byte

	for i := 0; i < len(str); i++ {
		ch = str[i]
		fmt.Println(ch)
		fmt.Printf("ASCII value of character '%c': %d\n", ch, ch)
	}

	// fmt.Println("ASCII value of first character:", str[0])

	// fmt.Println("Substring of string", str[1:4])
	// fmt.Println("converted string of first character:", string(str[0]))
}
