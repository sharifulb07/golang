package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)



type Job struct{
	ID int 
	URL string 
}

type Result struct{
	ID int 
	URL string 
	StatusCode int 
	Success bool 
	Err string
}




const (
	totalWorkers = 5
	requestDelay = 500 * time.Millisecond
	timeout=10*time.Second
)

func worker(
	id int,
	jobs <- chan Job ,
	result chan <- Result,
	wg *sync.WaitGroup,
	client *http.Client,
	throttle <-chan time.Time, 
) {

	for job:=range jobs{

		<-throttle

		fmt.Printf("Worker %d scraping %s \n", id, job.URL)

		resp, err:=client.Get(job.URL)

		if err!=nil{
			result <- Result{
				ID: job.ID,
				URL: job.URL,
				Success: false,
				Err: err.Error(),
			}
			continue 
		}

		resp.Body.Close()

		result <- Result{
			ID: job.ID,
			URL: job.URL,
			StatusCode: resp.StatusCode,
			Success: true,
		}
	}

}

func main() {
	urls:=[]string{
		"https://example.com",
		"https://golang.org",
		"https://github.com",
		"https://httpbin.org/get",
		"https://jsonplaceholder.typicode.com/posts",
		"https://stackoverflow.com",
		"https://news.ycombinator.com",
		"https://go.dev",
		"https://www.cloudflare.com",
		"https://www.wikipedia.org",

	}

	jobs:=make(chan Job, len(urls))
	results:=make( chan Result, len(urls))


	var wg sync.WaitGroup

	client:=&http.Client{
		Timeout: timeout,
	}

	throttle:=time.Tick(requestDelay)


	for i:=0;i<totalWorkers;i++{
		wg.Add(1)

		go worker(i, jobs, results, &wg, client, throttle)
	}


	// sending jobs

	for i, url:=range urls{
		jobs <- Job{
			ID: i+1,
			URL: url,
		}
	}

	close(jobs)

	wg.Wait()

	close(results)


	for result:=range results{

		if result.Success{
			fmt.Printf("[Success] Job %d | %s | status %d \n", result.ID, result.URL, result.StatusCode )
			}else{
				
				fmt.Printf("[Failed] Job %d | %s | status %d \n", result.ID, result.URL, result.StatusCode )
		}
	}
}