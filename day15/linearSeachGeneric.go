

package main 

import (
	"fmt"
)



func linearSearchGeneric[T comparable] (arr[]T, target T) int {

	for i, val:=range arr{
		if val==target{
			return i 
		}
	}
	return -1


}


func main(){

	fmt.Println(linearSearchGeneric([]string{"a","b","c","d", "e"}, "e"))
	fmt.Println(linearSearchGeneric([]int{1,2,3,2,45,5,47,8},8))
}


