package main

import (
	"fmt"
	"runtime"
)

func main() {

	var m runtime.MemStats

	var holder [][] byte 


	for i:=0; i<1_00_0000; i++{
		b:=make([]byte, 1024) // allocate memeory
		holder=append(holder, b)
	}


	runtime.GC()
	runtime.ReadMemStats(&m)

	fmt.Println("Allocate: ", m.Alloc)
	fmt.Println("Total Allocate: ", m.TotalAlloc)

	fmt.Println("NumGC: ", m.NumGC)

}