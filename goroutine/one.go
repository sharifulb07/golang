package main

import (
	"fmt"
	"sync"
)

func main() {
	counter:=0

	var wg sync.WaitGroup
	// var mu sync.Mutex

	for i:=0;i<10;i++{
		wg.Add(1)

		go func ()  {
			defer wg.Done()

			// mu.Lock()
			counter++
			fmt.Println(counter)
			// mu.Unlock()

		}()
	}

	wg.Wait()

	fmt.Println("all go routine are finished ")


}