package main

import "fmt"

func concat(s1, s2 string) string {
	return s1+s2
}



func main(){

	test("Lane", "happy birthday")
	test("Elon", "hope that Test things work out")
	test("Go ", "is fantastic")



}



func test(s1, s2 string){
	
	fmt.Println(concat(s1,s2))
}