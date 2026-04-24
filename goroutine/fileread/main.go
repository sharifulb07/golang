package main

import (
	"fmt"
	"os"
	"sync"
)


type finalResult struct{
	FileName string
	Content string
	Err error
}




func readFile(file string, wg *sync.WaitGroup, ch chan finalResult){

	defer wg.Done()
	data, err:=os.ReadFile(file)



	result:=finalResult{
		FileName: file,
		Content: string(data),
		Err: err,
	}

	ch <-result

}



func main(){
	files:=[]string{
	"C:\\Users\\HP\\Desktop\\go\\goroutine\\fileread\\file1.txt",
	"C:\\Users\\HP\\Desktop\\go\\goroutine\\fileread\\file2.txt",
	"C:\\Users\\HP\\Desktop\\go\\goroutine\\fileread\\file3.txt",
	}
	var wg sync.WaitGroup
	ch:=make(chan finalResult)

	for _, file:=range files{
		wg.Add(1)
		go readFile(file, &wg, ch)
	}


	go func(){
		wg.Wait()
		close(ch)
	}()


	// result receiver

	for result:=range ch{

		if result.Err!=nil{
			fmt.Printf("fetching data err %s:%v \n", result.FileName, result.Err)
			continue
		}

		fmt.Println("Content: ")
		fmt.Println(result.FileName)
		fmt.Println(result.Content)
		fmt.Println("....................")



	}
}

