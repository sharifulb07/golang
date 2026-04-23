// concurrent sum calculator

package main

import (
	"fmt"
	"sync"
)


func partialSum(nums[]int, ch chan int, wg *sync.WaitGroup){

	defer wg.Done()
	sum:=0

	for _, num:=range nums{
		sum+=num 
	}
	ch <- sum
}


func main(){
	arr:=[]int {12,5,4,7,89,63,64,2,8,9}

	mid:=len(arr)/2

	ch:=make(chan int, 2)
	var wg sync.WaitGroup

	wg.Add(2)

go partialSum(arr[:mid], ch,&wg)
go partialSum(arr[mid:], ch, &wg)

wg.Wait()
close(ch)

total:=0
for n:=range ch{
	total+=n 
}
fmt.Println(total)


}