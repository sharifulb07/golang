package main

import "fmt"


func main(){

	arr1:=[]int{5,8,9,7,6,3}
	arr2:=[] int{1,3,2,4,10,11,15,20}

	fmt.Println(mergedArrays(arr1, arr2))

}


// merge two arrays

func mergedArrays(nums1[]int, nums2[] int)[]int {

	m, n:=len(nums1),len(nums2)
	merged:=make([]int,0, m+n)
	i,j:=0,0

	for i<m && j<n{
		if nums1[i]<nums2[j]{
			merged=append(merged, nums1[i])
			i++
		}else{
			merged=append(merged, nums2[j])
			j++
		}
	}

	merged=append(merged, nums1[i:]...)
	merged=append(merged, nums2[j:]...)

	fmt.Println(len(merged))
	return  merged
}



// func mergedArrays(nums1[]int, nums2[] int)[]int{
// 	m, n:=len(nums1), len(nums2)
// 	merged:=make([]int,0)

// 	i, j:=0,0


// 	for i<m && j<n{
// 		if nums1[i]<=nums2[j]{
// 			merged=append(merged, nums1[i])
// 			i++
// 		}else{
// 			merged=append(merged, nums2[j])
// 			j++
// 		}


// 	}


// 	for i<m{
// 		merged=append(merged, nums1[i])
// 		i++
// 	}

// 	for j<n{
// 		merged=append(merged, nums2[j])
// 		j++
// 	}

// 	return  merged
// }