
package main

import (
	"fmt"
)


func main(){
	 maps:=map[string]string{"one":"Sujit sir", "two":"Kamrul Hossain", "three":"Asad sir", "four":"monirul"}

fmt.Println(maps["one"])

	//  for index :=range maps{
	// 	fmt.Println(index)
	// 	fmt.Println(maps[index])
	//  }

	
}







// package main 
// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"sync"
// )


// func worker (id int, jobs <-chan string, results chan <- string, wg *sync.WaitGroup){
// 	defer wg.Done()

// 	// range 1....5->1,2,3,4,5





// 	for url :=range jobs{
// 		fmt.Printf("Worker %d fetching %s \n", id, url)

// 		// fetching url 

// 		resp, err:=http.Get(url)

// 		if err !=nil{
// 			fmt.Printf("Worker %d fetching errorUrl %s ", id, url)
// 			continue
// 		}

// 		body, _:=ioutil.ReadAll(resp.Body)

// 		resp.Body.Close()

// 		results <- fmt.Sprintf("Worker %d: %s -> %d bytes", id, url, len(body))

// 	}
// }

// func main (){

// 		// List of URLs to fetch
// 		urls := []string{
// 			"https://golang.org",
// 			"https://google.com",
// 			"https://github.com",
// 			"https://stackoverflow.com",
// 			"https://reddit.com",
// 		}

// 		jobs:=make(chan string, len(urls))
// 		results:=make(chan string, len(urls))

// 		var wg sync.WaitGroup


// 		// worker start 

// 		for i:=0;i<=3;i++{
// 			wg.Add(1)
// 			go worker(i, jobs, results, &wg)
// 		}

// 		for _, url:=range urls{

// 			jobs <- url
// 		}

// 		close(jobs)

// 		go func(){
// 			wg.Wait()
// 			close(results)
// 		}()


// 		for r:=range results{

// 			fmt.Println("Result", r)
// 		}


// }

























// package main 
// import (
// 	"fmt"
// 	"sync"
// )


// // worker function

// func worker(id int, jobs <-chan int, results chan<-int , wg *sync.WaitGroup){

// 	defer wg.Done()

// 	for job:=range jobs{
// 		fmt.Printf("Worker %v processing %v", id, job)
// 		results <- job*3
// 	}
// }

// func main(){

// 	jobs:=make(chan int,5)
// 	results:=make(chan int,5)

// 	var wg sync.WaitGroup

// 	// worker pool start

// 	for i:=1; i<=5;i++{
// 		wg.Add(1)
// 		go worker(i, jobs, results, &wg)
// 	}

// 	// jobs sends
// 	for i:=1;i<=5;i++{

// 	jobs <- i
// 	}
// 	close(jobs)


// 	// another go routine syntax 

// 	go func ()  {
// 		wg.Wait()
// 		close(results)
		
// 	}()


// 	// Results 

// 	for r:=range results{
// 		fmt.Println("Result", r)
// 	}


// }