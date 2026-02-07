package main


func main(){



	



}


// moves zeroes to end 




// arr:=[]int{12,5,8,1,3,3,6,4,8,6,8,9,5,7,5}
// 	fmt.Println(removeDuplicates(arr))

// remove duplicate values from slice 


// func removeDuplicates(nums[] int) [] int {

// 	if(len(nums)==0){
// 		return nums
// 	}

// 	sort.Slice(nums, func(i, j int) bool {return nums[i]<nums[j]})

// 	k:=1

// 	for i:=1;i<len(nums);i++{
// 		if nums[i]!=nums[i-1]{
// 			nums[k]=nums[i]
// 			k++
// 		}

// 	}
// 	return nums[:k]

// }




// func removeDuplicates(nums[]int )[]int{

// 	sort.Slice(nums, func(i, j int) bool {return nums[i]<nums[j]})

// 	tempMap:=make(map[int]int, len(nums))
// 	tempArr:=make([]int,0)

// 	for _, v:=range nums{

// 		if id, found:=tempMap[v]; !found{
// 			tempMap[v]=id
// 			tempArr=append(tempArr, v)
// 		}
// 	}
// 	return tempArr
// }



	// arr:=[]int {10,20,30,40,50,60}
	
	// fmt.Println(twoSum(arr, 110))

// // using struct and two pointer approaches 

// type Pair struct{
// 	value int
// 	index int
// }

// func twoSum(nums[]int, target int)[]int{

// pairs:=make([]Pair, len(nums))

// for i, v:=range nums{
// 	pairs[i]=Pair{v,i}
// }

// // reverse pairs slice 
// sort.Slice(pairs,func(i, j int) bool {return pairs[i].value<pairs[i].value})

// left, right:=0, len(pairs)-1

// for left<right{
// 	sum:=pairs[left].value+pairs[right].value
// 	if sum==target{
// 		return []int{pairs[left].index, pairs[right].index}
// 	}else if sum<target{
// 		left++
// 	}else{
// 		right--
// 	}
// }
// return []int {}

// }



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