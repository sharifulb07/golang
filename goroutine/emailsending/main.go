package main

import (
	"fmt"
	"sync"
	"time"
)

type EmailJob struct {
	UserId int
	Email  string
}

func emailWorker(id int, jobs <-chan EmailJob, wg *sync.WaitGroup){
	defer wg.Done()
	for job :=range jobs{
		fmt.Printf("Worker %d started sending email to user %d email: %s \n", id, job.UserId, job.Email )
		time.Sleep(1*time.Second)

		fmt.Printf("Worker %d finished sending email to user %d email: %s \n", id, job.UserId, job.Email )
	}

}

func main(){
	const totalJobs=30
	const totalWorker=5

	var wg sync.WaitGroup
	jobs:=make(chan EmailJob, totalJobs)
	for i:=0; i<totalWorker;i++{
		wg.Add(1)
		go emailWorker(i, jobs, &wg)
	}


	for j:=1;j<totalJobs;j++{
		jobs <-EmailJob{
			UserId: j,
			Email: fmt.Sprintf("user %d@gmail.com", j),
		}
	}

	close(jobs)

	wg.Wait()


	fmt.Println("All worker sending email done ")
}