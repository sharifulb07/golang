
package main
import "fmt"


func main(){


	arr:=[]int{1,0,2,5,6,8,0,0,10,20,8,0,9}
	fmt.Println(MoveZeroToEnd(arr))

}

// moves zeroes to end of a slice using temp veriable tracking 

func MoveZeroToEnd(nums[]int)[]int{
	insertPos:=0

	if len(nums)==0{
		return  nums
	}

	for i, v:=range nums{
		if nums[i]!=0{
			nums[insertPos]=v
			insertPos++
		}
	}

	for insertPos< len(nums){
		nums[insertPos]=0
		insertPos++
	}
	return  nums

}





// moves zeroes to end of a slice

// func MoveZeroToEnd(nums[]int)[]int{
// 	if len(nums)==0{
// 		return nums
// 	}

// 	tempArr:=make([]int, 0)

// 	for _, v:=range nums{
// 		if v!=0{
// 		tempArr=append(tempArr, v)
// 		}
// 	}
// 	for _, v:=range nums{
// 		if v==0{
// 			tempArr=append(tempArr, v)
// 		}
// 	}
// 	return  tempArr

// }