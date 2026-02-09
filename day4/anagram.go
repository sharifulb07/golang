package main

import "fmt"

func isAnagram(s string, t string)bool{

	if len(s) !=len(t){
		return false
	}

	characterMap:=make(map[rune]int)

	for _, char:=range s{
		characterMap[char]++
	}

	for _, char:=range t{
		characterMap[char]--
		if characterMap[char]<0{
			return  false
		}
	}

	return  true
}

func main(){
	fmt.Println(isAnagram("Hello", "olleH"))
}