package main
import "fmt"



func rotateByKth(arr[]int, k int){
	n:=len(arr)-1

	if n==0{
		return
	}
	k=k%n
	// full reverse 
	reverse(arr, 0, n)
	// first part rotation before k value
	reverse(arr, 0, k-1)

	// end part reverse 
	reverse(arr,k, n)


}

func reverse(arr[]int,start int, end int){
	for start<end{
		arr[start], arr[end]=arr[end], arr[start]
		start++
		end--
	}
}

func main(){
	arr:=[]int{1,2,3,4,5,6,7,8,9}

	fmt.Println("Main Array: ", arr)
	rotateByKth(arr, 4)
	fmt.Println(arr)
}