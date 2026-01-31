package main
import "fmt"




func square( num int, ch chan int ){

	// keep integer value in channel 
	ch <-num*num

}


func main(){


	ch:= make(chan int)

	for i:=0; i<=5;i++{
		go square(i, ch)
		fmt.Println(ch)
	}
	


	for i:=0;i<=5;i++{
		result:= <-ch
		fmt.Println("Result: ", result)
	}


}

















// package main
// import (
// 	"fmt"
// 	"sync"
// )


// func worker(id int, wg *sync.WaitGroup){
// 	defer wg.Done()

// 	fmt.Println("Worker", id, "started")
// }


// func main(){

// 	var wg sync.WaitGroup

// 	for i:=0;i<=5;i++{
		 
// 		wg.Add(1)
// 		go worker(i, &wg)







// 	}

// 	wg.Wait()
// 	fmt.Println("Allo goroutines tasks are finished")



// }










// package main

// import (
// 	"fmt"
// 	"time"
// )


// func PrintNum(n int){
// 	fmt.Println(n)
// }


// func main(){

// 	// why we use loop 
// 	// repetition  of same task

// 	// concurrency -> at the same time -> multiple task perform 

// 	// go PrintNum(0)->0
// 	// go PrintNum(1)->1
// 	// go PrintNum(2)->2
// 	// go PrintNum(3)->3
// 	// go PrintNum(4)->4
// 	// go PrintNum(5)->5

// 	// 2 5 3 4 0 1

// 	for i:=0;i<=5;i++{
// 		go  PrintNum(i)
// 	}

// 	time.Sleep(time.Second)

// }