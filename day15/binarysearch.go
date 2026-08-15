// find first and last occurrence of a value in a sorted array

package main

import (
	"fmt"
	"sort"
)


func findFirst(nums[]int, target int) int{
	left, right:=0, len(nums)-1
	result:=-1

	for left<=right{
		mid:=left+(right-left)/2

		if nums[mid]==target{
			result=mid 
			right=mid-1
		}else if(nums[mid]<target){
			left=mid+1
		}else{
			right=mid-1
		}
	}
	return result
}
func findLast(nums[]int, target int) int{

	left, right:=0, len(nums)-1
	result:=-1

	for left<=right{
		mid:=left+(right-left)/2

		if nums[mid]==target{
			result=mid 
			left=mid+1
		}else if(nums[mid]<target){
			left=mid+1
		}else{
			right=mid-1
		}
	}
	return result
}


func searchRange(nums[] int, target int) [2]int  {
	
	return [2] int {findFirst(nums,target), findLast(nums,target)}
}


func main(){

	nums:=[]int{2,2,5,6,3,2,6,8,2,9,2,7,2,10,2,4}
	sort.Ints(sort.IntSlice(nums))
	target:=2

	fmt.Println(searchRange(nums,target))

}