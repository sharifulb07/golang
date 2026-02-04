package main
import "fmt"






func main(){

	
	

}

// container with most water

// 
func maxWater(height[]int) int{
	left, right:=0, len(height)
	maxArea:=0


	for left<right{
		h:=min(height[left], height[right])
		area:=h*right-left

		if area>maxArea{
			maxArea=area
		}

		if height[left]<height[right]{
			left++
		}else{
			right--
		}


	}
	return maxArea
} 

func




// driver code 




// officers:=[] string {"Sharif", "Hunaif", "Ashia", "Safi"}


// 	for i:=0;i<len(officers);i++{
// 		r:=[]rune(officers[i])
// 		stringReverse(r)
// 		officers[i]=string(r)
// 	}

	

// 	fmt.Println(officers)


// 	myQoute:=[]rune("I am a developer")

// 	stringReverse(myQoute)

// 	fmt.Println(string(myQoute))



// string reverse using two two pointer approach

// func stringReverse(st[]rune){

// 	left, right:=0, len(st)-1

// 	for left<right{

// 		st[left], st[right]=st[right], st[left]
// 		left++
// 		right--
// 	}


// }



	
	// arr:=[]int{10,20,30,40,50,60,70,80,90,100}
	
	// fmt.Println(maxSumArr(arr, 3))


	// arr:=[]int{10,20,30,40,50,60,70,80,90,100}
	// fmt.Println(twopointer(arr, 90))
	// fmt.Println(twopointer(arr, 900))



	// arr:=[]int {1,2,2,2,3,5,10,6,5,6,6,8,8,9}


	// ascending sorting of slice in go

	// sort.Ints(arr)
	// fmt.Println(arr)
	
	// sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	// sort.Reverse(sort.IntSlice(arr))

	// sort.Sort-> function
	// sort.Reverse-> function logic
	// sort.intSlice-> sortable type
	
	// fmt.Println(arr)



// 	result:=removeDuplicate(arr )

// fmt.Println(result)

	
	




	// 


	// remove duplicates from an sorted array and return length of array 
	// use two pointer approach 



	// func removeDuplicate(nums[]int) int{

	// 	sort.Ints(nums)
	// 	fmt.Println(nums)
	// 	if len(nums)==0{
	// 		return 0
	// 	}

	// 	i:=0

	// 	for j:=1;j<len(nums);j++{
	// 		if nums[j]!=nums[i]{
	// 			i++
	// 			nums[i]=nums[j]
	// 		}
	// 		// fmt.Println(nums[i])
	// 	}

	// 	// fmt.Println(nums[:i+1])
	// 	fmt.Println(nums)
	// 	return i+1


	// }




// sliding window technique in go lang 



// func maxSumArr(arr[]int, k int) int{
// 	windowSum:=0
// 	maxSum:=0

// 	for i:=0;i<k;i++{
// 		windowSum+=arr[i]
// 	}

// 	maxSum=windowSum
// 	for i:=k;i<len(arr);i++{
// 		windowSum+=arr[i]-arr[i-k]
// 		if windowSum>maxSum{
// 			maxSum=windowSum
// 		}
// 	}

// 	return maxSum
// }




	// Two pointers technique now
	// in sorted array
	
func twopointer(arr[]int, target int)bool{
	i, j:=0, len(arr)-1;

	for i<j{
		sum:=arr[i]+arr[j]
		if(sum==target){
			return  true
		}else if(sum<target){
			i++
		}else{
			j--
		}

		
	}
	return  false
}