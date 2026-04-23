package main

import "fmt"

func main() {

	fmt.Println(add(50,80))

}

func add(a, b int) (result int) {
	result = a + b
	return
}