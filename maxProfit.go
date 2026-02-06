package main

import "fmt"

func main() {

	arr := []int{1, 5, 8,  9, 60, 7, 6, 4, 8, 50, 4, 10}
	fmt.Println(maxProfit(arr))


}

func maxProfit(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	minPrice := nums[0]
	maxProfit := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] < minPrice {
			minPrice = nums[i]
		} else {
			profit := nums[i] - minPrice

			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}