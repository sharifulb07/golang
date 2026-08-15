package main

import "fmt"


func main(){

	arr:=[]int{1,2,3,4,5,6,8,9}
	fmt.Println(missingNum(arr))

}




// using formula

// func missingNum(nums[]int)int{
// 	n:=len(nums)+1
// 	actualSum:=(n*(n+1))/2

// 	expectedSum:=0

// for _, num:=range nums{
// 	expectedSum+=num
// }

// return actualSum-expectedSum

// }

// xor approach

// func missingNum(nums[]int)int {
// 	n:=len(nums)+1
// 	xorAll:=0
// 	xorArr:=0


// 	for i:=1;i<=n;i++{
// 		xorAll ^=i
// 	}

// 	for _, num:=range nums{
// 		xorArr ^=num
// 	}
// 	return xorAll ^xorArr

// }





// hashing approach


// func missingNum(nums[]int) []int{
// 	seen:=make(map[int]bool)
// 	missing:=[]int{}



// 	// all attempts mark true

// 	for _, num:=range nums{
// 		seen[num]=true
// 	}

// 	for i:=1;i<=len(nums)+1;i++{
// 		if !seen[i]{
// 			missing=append(missing, i)

// 		}
// 	}

// 	return missing

// }