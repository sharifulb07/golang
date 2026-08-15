package main 
import "fmt"

func main(){
	arr:=[]int {1,2,3,4,45,6,7,89}

	fmt.Println(findMax(arr))
	fmt.Println(findMin(arr))

}





// func findMax(nums[]int)int{
// 	if len(nums)==0{
// 		panic("empty array")
// 	}

// 	max:=nums[0]
// 	for i:=1;i<len(nums); i++{
// 		if nums[i]>max{
// 			max=nums[i]
// 		}
// 	}
// 	return  max 
// }
// func findMin(nums[]int)int{
// 	if len(nums)==0{
// 		panic("empty array")
// 	}

// 	min:=nums[0]
// 	for i:=1;i<len(nums); i++{
// 		if nums[i]<min{
// 			min=nums[i]
// 		}
// 	}
// 	return  min 
// }