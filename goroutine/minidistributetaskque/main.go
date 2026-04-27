package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)


type Task struct{
	ID int 
	Type string
	Payload string
	Retry int 
	MaxTry int 
	CreatedAt time.Time
}




var(
	mainQueue= make(chan *Task, 100)
	dlq=make(chan *Task, 100)
)


func worker( id int, wg *sync.WaitGroup){

	defer wg.Done()

	

	for task:=range mainQueue{
		fmt.Printf(
			"Worker %d processing Task %d (%s) \n",
			id,
			task.ID,
			task.Type, 
		)

		success:=processTask(task)

		if success{
			fmt.Printf(
				"Task %d completed successfully \n",
				task.ID, 
			)
			continue 
		}

		task.Retry++
if task.Retry <= task.MaxTry {
			fmt.Printf(
				"Task %d failed → retrying (%d/%d)\n",
				task.ID,
				task.Retry,
				task.MaxTry,
			)

			time.Sleep(2 * time.Second)

			mainQueue <- task
		} else {
			fmt.Printf(
				"Task %d moved to DLQ\n",
				task.ID,
			)

			dlq <- task
		}
	}
}


func processTask(task *Task)bool{
	time.Sleep(1*time.Second)

	return rand.Intn(10)<7
}



func monitorDLQ(wg *sync.WaitGroup)  {

	defer wg.Done()

	for task:=range dlq{
		fmt.Printf(
			"|DLQ| Task %d permanently failed (%s) \n",
			task.ID,
			task.Payload, 
		)
	}


	
}
func main() {


	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	for i:=0; i<5; i++{
		wg.Add(1)
		go worker(i, &wg)

	}
	wg.Add(1)
	go monitorDLQ(&wg)


	// simulator producer 

	for i:=1; i<=10; i++{
		task:=&Task{
			ID: i,
			Type: "email",
			Payload: fmt.Sprintf("Send Welcome email to user %d", i ),
			MaxTry: 3,
			CreatedAt: time.Now(),
		}

		mainQueue <- task
	}

	time.Sleep(20*time.Second)

	close(mainQueue)
	close(dlq)

	wg.Wait()

	fmt.Println("\n All Tasks Processed ")

}