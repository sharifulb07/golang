package main

import (
	"fmt"
	"net/http"
)


type Person struct{
	Name string 
	Age int64
}


type speed interface{}


type handler func( w http.ResponseWriter, r *http.Request)


type Router struct{
	routes map[string] handler
}




func main() {

	p1:=Person{Name: "Sujit", Age: 45}
	var s1 speed=60

	fmt.Println(p1)
	fmt.Println(s1)
	fmt.Println("Hello world")

}