package main

import "fmt"

func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)

}

func merge(left []int, right []int) []int {

	var result []int
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result

}

func main() {

	arr := []int{10, 8, 6, 21, 9, 7, 1, 2, 4}

	fmt.Println("Original Array: ", arr)

	sorted:=MergeSort(arr)

	fmt.Println("Merge Sorted Array: ", sorted)
}