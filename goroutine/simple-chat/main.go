package main

import (
	"fmt"
	"time"
)

func sender(ch chan string) {

	messages := []string{
		"Hello",
		"How are You",
		"Channels in go lang are awesome",
		"yeah I'm learning go routine",
		"okay main , Bye",
	}

	for _, msg := range messages {
		fmt.Printf("sender: %s \n", msg)
		ch <-msg
		time.Sleep(1*time.Second)


	}
	close(ch)
}

func receiver(ch chan string) {

	for msg:=range ch{
		fmt.Println("receiver: ", msg)
		time.Sleep(500*time.Microsecond)
	}

}


func main(){
	ch:=make(chan string)

	go sender(ch)
	receiver(ch)


}