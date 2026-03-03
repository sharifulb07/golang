package main 
import (
	"fmt"
)


func binarySearch(arr []int, target int) int{
	left:=0
	right:=len(arr)-1

	for left<right{
		mid:=left+(right-left)/2

		if arr[mid]==target{
			return mid

		}else if(arr[mid]<target){
			left=mid+1
		}else{
			right=mid-1

		}
	}

	return -1


}


func main(){

	arr:=[]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	result:=binarySearch(arr,13)

	fmt.Println(arr[result])
}