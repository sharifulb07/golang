package main
import "fmt"






func main(){

	
	
	// arr:=[]int{10,20,30,40,50,60,70,80,90,100}
	
	// fmt.Println(maxSumArr(arr, 3))


	// arr:=[]int{10,20,30,40,50,60,70,80,90,100}
	// fmt.Println(twopointer(arr, 90))
	// fmt.Println(twopointer(arr, 900))



	
	}



// sliding window technique in go lang 



func maxSumArr(arr[]int, k int) int{
	windowSum:=0
	maxSum:=0

	for i:=0;i<k;i++{
		windowSum+=arr[i]
	}

	maxSum=windowSum
	for i:=k;i<len(arr);i++{
		windowSum+=arr[i]-arr[i-k]
		if windowSum>maxSum{
			maxSum=windowSum
		}
	}

	return maxSum
}




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