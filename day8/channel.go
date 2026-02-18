package main 

import "fmt"

func main(){

	message:=make(chan string)

	go func (){message <- "Hi Is am shaprif who is a react native developer right now. I always write new apps to protect my own appps"}()

	msg:=<- message 

	fmt.Println(msg)
}

