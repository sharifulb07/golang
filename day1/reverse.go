// reverse an string

package main 
import "fmt"

func reverse(nums[]int){
	left, right:=0, len(nums)-1

	for left<right{
		nums[left], nums[right]=nums[right], nums[left]
		left++
		right--
	}
}

func main (){
	arr:=[]int{1,2,3,4,5,6,7}
	reverse(arr)
	fmt.Println(arr)
}