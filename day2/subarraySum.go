package main

import "fmt"

func main(){

arr:=[]int{1,2,4,5,6,11,8,10,7,8,}

fmt.Println(subarraySum(arr, 19))



}

func subarraySum(nums[]int, k int) int{
	count:=0
	sum:=0

	prefixSum:=make(map[int]int)
	prefixSum[0]=1

	for _, num:=range nums{
		sum+=num
		fmt.Println("Sum:", sum)



		if freq, exists:=prefixSum[sum-k]; exists{
			fmt.Println(count)
			count+=freq
			fmt.Println("after: ", count)
		}
		prefixSum[sum]++

		fmt.Println("PrefixSux :", prefixSum)

	}

	return count
}