package main
import "fmt"


// func increasingSubArr(nums[]int)int{

// 	currentLength:=0
// 	maxLength:=0

// 	for i:=1;i<len(nums);i++{
// 		if nums[i]>nums[i-1]{
// 			currentLength++
// 			if currentLength>maxLength{
// 				maxLength=currentLength
// 			}
// 		}else{
// 			currentLength=1
// 		}
// 	}

// 	return  maxLength

// }


func increasingSubArr1(nums[]int)[]int{

	maxLength:=0
	maxStart:=0
	maxEnd:=0

	currentLength:=1
	currentStart:=1

	for i:=1;i<len(nums);i++{
		if nums[i]>nums[i-1]{
			currentLength++
			if currentLength>maxLength{
				maxLength=currentLength
				maxStart=currentStart
				maxEnd=i
			}
		}else{
			currentLength=1
			currentStart=i

		}
	}

	return nums[maxStart:maxEnd+1]


	
}



func main(){
	arr:=[]int {1,5,4,8,7,4,8,2,14,1,4,7,5,7,9,10}
	fmt.Println(increasingSubArr1(arr))
}