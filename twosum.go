package main

import (
	"fmt"
	"sort"
)


func main(){


	arr:=[]int {10,20,30,40,50,60}
	
	fmt.Println(twoSum(arr, 110))



}



// using struct and two pointer approaches 

type Pair struct{
	value int
	index int
}

func twoSum(nums[]int, target int)[]int{

pairs:=make([]Pair, len(nums))

for i, v:=range nums{
	pairs[i]=Pair{v,i}
}

// reverse pairs slice 
sort.Slice(pairs,func(i, j int) bool {return pairs[i].value<pairs[i].value})

left, right:=0, len(pairs)-1

for left<right{
	sum:=pairs[left].value+pairs[right].value
	if sum==target{
		return []int{pairs[left].index, pairs[right].index}
	}else if sum<target{
		left++
	}else{
		right--
	}
}
return []int {}

}



// two pointer approach

// func twoSumWithTwoPointer(nums[]int, target int) []int{

// 	sort.Slice(nums, func(i, j int) bool {return nums[i]<nums[j]})

// 	left, right:=0, len(nums)-1

// 	for left<right{
// 		sum:=nums[left]+nums[right]

// 		if sum==target{
// 			return  []int{left, right}
// 		}else if sum<target{
// 			left++
// 		}else{
// 			right--
// 		}
// 	}
// 	return []int {}
// }

// using hashmap 
// func twoSum(nums[]int, target int )[]int{

// 	hashMap:=make(map[int]int)


// 	for i, num :=range nums{
// 		complement:=target-num

// 		if idx, found:=hashMap[complement];
// 		found{
// 			return []int {idx, i}
// 		}
// 		hashMap[num]=i
// 	}

// 	return []int{}
// }