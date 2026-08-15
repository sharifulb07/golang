package main

import (
	"fmt"
	"strings"
)


func main(){

	original:="hello World. Journey to new World"
	reversed:=reverse(original)

	fmt.Println("Original", original)
	fmt.Println("Reversed : ", reversed)


	first:="I am Go developer"
	reversedSecondo:=reverseString(first)

	fmt.Println(first)
	fmt.Println(reversedSecondo)

}

func reverse(s string) string{
	str:=[]rune(s)

	for i, j:=0, len(str)-1; i<j; i, j=i+1, j-1{

		str[i], str[j]=str[j], str[i]
	}

	return  string(str)
}



// string builder for more string 

func reverseString(s string) string{
	str:=[]rune(s)
	var builder strings.Builder
	builder.Grow(len(s))

	for i:=len(str)-1; i>=0; i--{
		builder.WriteRune(str[i])
	}
	return  builder.String()
}