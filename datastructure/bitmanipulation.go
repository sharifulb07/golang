package main

import "fmt"

func main() {

	arr:=[]int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(Sum(arr))



	// fmt.Println(setkthbit(5,2))

	// fmt.Println(isEven(6))

	// isActive:=false
	// fmt.Println(isActive)

	// fmt.Println(isKthBitSet(5, 1))

	// a:=5
	// b:=3

	// fmt.Println(a & b)

}

// O(1) Time complexity constants

func Sum(temp []int) int {

	return (temp[len(temp)-1] * (temp[len(temp)-1] + temp[0])) / 2
}

// O(n) Time complexity
// func Sum(temp [] int) int{
// 	total:=0

// 	for i:=range temp{
// total+=temp[i]
// 	}

// 	return total
// }

// func setkthbit(n int, k int) int{
// 	return  n| (1<<k)
// }

// func isKthBitSet(n int, k int) bool{
// 	return (n & (1<<k)) !=0
// }

//even->  0, 2,4,6,8
// odd -> 1,3,5,7,9
// 101
//   1
//

// func isEven(n int) bool{
// 	return  n&1==0

// 	// 5&1
// 	// 1
// }