package main 

import (
	"fmt"

)


func main(){

	// What is the difference between array and slice?

// 	type slice struct {
//     arr:=[5]string{"i", "you", "they", "she", "it"}
//     len int  5// number of elements currently used
//     cap int  5// total available space from ptr onward
// }

arr:=[3]int{4,5,8}
fmt.Println(cap(arr))



}