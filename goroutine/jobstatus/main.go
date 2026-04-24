package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type JobStatus string

const (
	Pending JobStatus = "pending"
	Running JobStatus = "running"
	Done    JobStatus = "done"
	Failed  JobStatus = "failed"
)

type Job struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Data   string `json:"data"`
	Status string `json:"status"`
	Result string `json:"result"`
}

var (
	jobQue     = make(chan *Job, 1000)
	jobStore   = make(map[int]*Job)
	storeMutex sync.RWMutex
)

// ----worker-----
func worker(id int){

	for job:=range jobQue{
		updateStatus(job.ID, Running)

		fmt.Printf("Worker %d processing job %d \n", id, job.ID)

		time.Sleep(time.Duration(rand.Intn(3)+1)*time.Second)


		// simulate success/failure

		if rand.Intn(10)<8{
			job.Result="Process successfully "+job.Data
			updateStatus(job.ID, Done)
		}else{
			job.Result="Processing Failed"
			updateStatus(job.ID, Failed)
		}

		fmt.Printf("Worker %d finished job %d \n", id, job.ID)
	}

}


// ---update status----


func updateStatus(id int, status JobStatus){
	storeMutex.Lock()
	defer storeMutex.Unlock()

	if job, ok:=jobStore[id]; ok{
		job.Status=string(status)
	}
}



// get job

func getJob(id int)(*Job, bool){
	storeMutex.RLock()
	defer storeMutex.RUnlock()

	job, ok:=jobStore[id]
	return job, ok 
}


// create job 

func createJob(c *gin.Context){

	var input struct{
		Type string `json:"type"`
		Data string `json:"data"`
	}

	if err:= c.ShouldBindJSON(&input); err != nil{

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	job:=&Job{
		ID: rand.Intn(1000000),
		Type: input.Type,
		Data: input.Data,
		Status: string(Pending),
	}



	storeMutex.Lock()

	jobStore[job.ID]=job

	storeMutex.Unlock()

	jobQue <- job

	c.JSON(http.StatusAccepted, gin.H{
		"message":"Job submitted",
		"job_id":job.ID, 
	})
}



// get job status

func getJobStatus(c *gin.Context){
	var input struct{
		ID int `json:"id"`
	}

	id:=c.Param("id")

	fmt.Sscanf(id, "%d", &input.ID)

	job, exits:=getJob(input.ID)

	if !exits{
		c.JSON(http.StatusNotFound, gin.H{"error":"job not found"})
		return 
	}

	c.JSON(http.StatusOK, job)
}


func main(){
	rand.Seed(time.Now().UnixNano())

	// start workers

	for i:=1;i<=5;i++{
		go worker(i)
	}

	r:=gin.Default()

	r.POST("/jobs", createJob)
	r.GET("/jobs/:id", getJobStatus)

	fmt.Println("Server is running on 8080")

	r.Run(":8080")
}
