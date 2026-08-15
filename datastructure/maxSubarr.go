package main
import "fmt"



func main(){

	arr:=[]int{12,4,5,9,8,6,3,7,90}
	fmt.Println(maxSubarr(arr))

}
// 

func maxSubarr(nums[]int)int{
maxSum:=nums[0]
curSum:=0

for i:=0;i<len(nums);i++{
	curSum+=nums[i]
	if curSum>maxSum{
		maxSum=curSum
	}

	if curSum<0{
		curSum=0
	}
}
return  maxSum


}



// func maxSubarr(nums[]int)int{
// 	maxSum:=nums[0]
// 	currentSum:=nums[0]
	

// 	for i:=1;i<len(nums);i++{

// 		currentSum=max(currentSum+nums[i], nums[i])

// 		maxSum=max(currentSum,maxSum)
// 	}
// 	return  maxSum
// }

// func max(a int, b int)int{
// 	if a>b{
// 		return  a
// 	}
// 	return  b
// }