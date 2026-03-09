// selection sort

package main 
import (
	"fmt"
)

// ascending order by selection sort in a slice in go lang 

func selectionSort(nums [] int){

	n:=len(nums)

	for i:=0; i<n-1;i++{
		minIndex:=i

		for j:=i+1;j<n;j++{
			if nums[j]< nums[minIndex]{
				minIndex=j
			}
		}



		nums[i], nums[minIndex]=nums[minIndex], nums[i]
	}
}



func main(){

	nums:=[]int {1,2,8,6,12,11,40,10,3,20,15,}

	selectionSort(nums)
	fmt.Println(nums)

}