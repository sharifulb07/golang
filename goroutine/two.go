package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Err error
}

func fetchUrl(url string, ch chan Result, wg *sync.WaitGroup){


	defer wg.Done()

	start:=time.Now()

	resp, err:=http.Get(url)

	if err !=nil{
		ch <- Result{
			URL: url,
			Err: err,
		}
		return 
	}

	defer resp.Body.Close()

	duration:=time.Since(start)

	ch <- Result{
		URL: url,
		StatusCode: resp.StatusCode,
		Duration: duration,
		Err: nil,

	}
}


func main(){
	urls:=[]string{
		"https://example.com",
		"https://github.com",
		"https://golang.org",
	}

	ch:=make(chan Result, len(urls))

	var wg sync.WaitGroup


	for _, url:=range urls{
		wg.Add(1)
		go fetchUrl(url, ch, &wg)

	}

	wg.Wait()
	close(ch)

	for result:=range ch{

		if result.Err!=nil{
			fmt.Printf("Fetching error %s in %d \n", result.URL, result.StatusCode)
		}

		fmt.Printf("Url: %s Status: %d Time: %v \n", result.URL, result.StatusCode, result.Duration)
	}

}