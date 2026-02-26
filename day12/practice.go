package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w, "Server is runing very well ")
}



func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)


}