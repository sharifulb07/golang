package main

import "fmt"

func main() {

	arr := []int{4, 5, 8, 6, 7, 9, 3}

	// fmt.Println(rightRotation(arr, 4))

	fmt.Println(leftRotation(arr, 2))

}

func rightRotation(nums []int, k int) []int {

	n := len(nums)
	if n == 0 {
		return nums
	}
	k = k % n

	reverse(nums, 0, n-1)
	fmt.Println("Whole array roation :", nums)
	reverse(nums, 0, k-1)
	
	fmt.Println("Zero to k roation :", nums)
	reverse(nums, k, n-1)
	fmt.Println("k to n  roation :", nums)

	return nums

}

func leftRotation(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nums
	}
	k = k % n

	reverse(nums, 0, k-1)
	
	fmt.Println("Zero to k roation :", nums)
	reverse(nums, k, n-1)
	fmt.Println("k to n  roation :", nums)
	reverse(nums, 0, n-1)
	fmt.Println("Whole array roation :", nums)

	return nums
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
