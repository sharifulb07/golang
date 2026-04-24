package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup){
	defer wg.Done()

	for job :=range jobs{
		fmt.Printf("Worker %d started jobs %d \n",id, job)
		time.Sleep(500*time.Millisecond)
		fmt.Printf("Worker %d finished jobs %d \n",id, job)

	}
}

func main(){
	const(
		totolJobs=100
		totalWorkers=5
	)

	var wg sync.WaitGroup
	jobs:=make(chan int, totolJobs)

	for i:=1; i<totalWorkers; i++{
		wg.Add(1)
		go worker(i,jobs, &wg)
	}


	for j:=1;j<totolJobs;j++{
		jobs <-j
	}

	close(jobs)

	wg.Wait()

	fmt.Println("All worker tasks are finished here ")
}