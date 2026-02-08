package main

import "fmt"

// func min(a int, b int) int{
// 	if a<b{
// 		return a
// 	}
// 	return b
// }


func minNum(num[]int) int{
	if len(num)==0{
		return  0
	}

	min:=num[0]
	
	for _, value:=range num{
		// 5,3
		// 3<5
		// 5->3

		if value<min{
			min=value
		}
	}
	return min
}

func main(){

	arr:=[]int{9,8,10,11,5,6,7}
	fmt.Println(minNum(arr))
}