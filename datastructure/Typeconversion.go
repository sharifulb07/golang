package main

import (
"fmt"
"log"
// "reflect"
)


func main(){

err:="User is not found Here"
	fmt.Errorf("LoadUser failed: %w", err)
	log.Println(err)
// var v interface{}="Go Lang"

// t:=reflect.TypeOf(v)
// value:=reflect.ValueOf(v)

// fmt.Println(t)
// fmt.Println(value)
// fmt.Println(value.Kind())








// 	var a int =10

// 	var b interface{}=float64(a)
// fmt.Printf("Type: %T", b)

// 	v, ok:=b.(float64)

// 	if ok {
// 		fmt.Printf("Type of value %T", v)
// 	}else{
// 		fmt.Printf("Unknown type")
// 	}

}