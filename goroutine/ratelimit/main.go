package main

import (
	"fmt"
	"time"
)

func main() {

	bucket := make(chan struct{}, 5)

	for i := 0; i < 5; i++ {
		bucket <- struct{}{}
	}

	ticker := time.NewTicker(1*time.Second)
	defer ticker.Stop()


	go func(){

		for range ticker.C{
			for len(bucket)<5{
				bucket <- struct{}{}
			}
			fmt.Println("Bucket refilled here ")
		}
	}()

	// simulate 30 request incoming requests here 

	for i:=0;i<30;i++{
		select{
		case <- bucket:
			fmt.Printf("Request %d allowed \n", i)
		default:
			fmt.Printf("Request %d rejected here \n", i)
		}

		time.Sleep(100*time.Millisecond)
	}
}