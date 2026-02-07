package main 
import "fmt"


func main(){


	arr:=[]int {2,5,8,9,6,7}
	fmt.Println(multiply(arr))

}

func multiply(nums[]int)[]int{
	n:=len(nums)-1
	
	result:=make([]int, n)
	product:=1

	// left multiplicaiton 
	for i:=0;i<n;i++{
		result[i]=product
		fmt.Println(result[i])
		product*=nums[i]
		fmt.Println(product)
	}

	// right multiplication

	product=1
	for i:=n-1;i>=0;i--{
		result[i]*=product
		product*=nums[i]

	}

	return  result


	
}