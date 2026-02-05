package main
import "fmt"


func main (){


	arr:=[]int{10,20,30,40,50,60,70,80,90}

	fmt.Println(twosum(arr, 80))
}


// optimal solution 
func twosum(nums[]int, target int ) []int{
	hashMap:=make(map[int]int)

	for i, num:=range nums{
		complement:=target-num

		if idx, found:=hashMap[complement]; found{
			return []int {idx, i}
		}
		hashMap[num]=i
	}

	return []int {}
}

