package main

import "fmt"


func main(){


	arr:=[]int {2,5,4,7,8,2,1,5,1,2,4,2,2,2,2,2,2,2,2}
	fmt.Println(findMajority(arr))

}


// Boyer Moore Voting Algorithm to find majority elements in a slice

func findMajority(nums[]int) int{
	count:=0
	candidate:=0

	for _, num:=range nums{
		if count==0{
			candidate=num
		}

		if num==candidate {
			count++
		}else{
			count--
		}
	}
	count=0
	for _, num:=range nums{
		if num==candidate{
			count++
		}
	}
	if count>len(nums)/2{
		return  candidate
	}

	return  -1
}