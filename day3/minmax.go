package main

import "fmt"



func reverseStr(s string) bool {


	runes:=[]rune(s)
	left, right:=0, len(runes)-1
	count:=0

	for left<right{
		if(runes[left]==runes[right]){
			runes[left], runes[right]=runes[right], runes[left]
		left++
		right--
		count++
		}
		
	}
	if count==len(runes)/2{
		return true
	}

	return  false


	// return  string(runes)
}

func main(){

	str:="wowsome"
	fmt.Println(reverseStr(str))

	

}