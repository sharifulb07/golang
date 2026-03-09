// searching a value in rotated array

package main 
import ( 
	"fmt"
)


func search(nums[]int, target int) int{

	left, right:=0, len(nums)-1

	for left<=right{

		mid:=left+(right-left)/2

		if nums[mid]==target{
			return mid 
		}

		if nums[left]<=nums[mid]{
	if nums[left]<=target && target<nums[mid]{
		right=mid-1


		}else{
			left=mid+1
		}
		}else{

			if nums[mid]< target && target <=nums[right]{
				left=mid+1
			}else{
				right=mid-1
			}
		}
	
	}
	return -1
}


func main(){


	nums:=[]int{10,9,8,7,6,0,1,2,3,4,5}

	target:=5

	fmt.Println(search(nums,target))

}