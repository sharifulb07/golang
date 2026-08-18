package main

import "fmt"

func main(){

	defer fmt.Println("last")
	fmt.Println("First")
}

