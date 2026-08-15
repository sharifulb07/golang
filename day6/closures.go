package main

import "fmt"

func initSeq() func() int {

	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

nextInit:=initSeq()
	fmt.Println(nextInit())
	fmt.Println(nextInit())





}
