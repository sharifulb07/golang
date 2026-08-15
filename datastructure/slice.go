package main

import (
	"fmt"
	"sort"
)

// func reverse(arr[] string){
// 	i, j:=0,len(arr)-1

// 	for i<j{

// 		// swap
// 		arr[i],arr[j]=arr[j],arr[i]

// 		i++
// 		j--
// 	}
// }

// func insertAt(slice [] int, index int, value int)[]int{

// 	slice=append(slice, 0)
// 	fmt.Println(slice)
// 	copy(slice[index+1:], slice[index:])
// 	slice[index]=value

// 	return slice
// }

// deletefast without ordering

// func deletefast(arr[]int , index int)[] int{

// 	arr[index]=arr[len(arr)-1]
// 	return  arr[:len(arr)-1]
// }

// rotate right by k position

// func rotateright(arr [] int, k int)[] int{
// 	n:=len(arr)
// 	k=k%n
// 	return  append(arr[n-k:], arr[:n-k]...)
// }
// func rotateleft(arr [] int, k int)[] int{
// 	n:=len(arr)
// 	k=k%n
// 	return  append(arr[k:], arr[:k]...)
// }

// func findMinMax(arr[]int)(int, int){

// 	if(len(arr)==0){
// 		panic("Slice is empty")
// 	}

// 	min:=arr[0]
// 	max:=arr[0]

// 	for _, v:=range arr{

// 		if v>max{
// 			max=v
// 		}
// 		if v<min{
// 			min=v
// 		}
// 	}

// 	return min, max

// }


func main() {





	arr:=[]int{1,2,3,4,5,6,7,8,9}

	// ascending sort
	sort.Ints(arr)

	fmt.Println(arr)

	// descending sort
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)

	names:=[]string{"Adiyat", "Sharif", "hunaif", "Ruhul Amin", "Go man"}

	
	sort.Strings(names)
	fmt.Println(names)

	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Println(names)


	speed:=[]float64{12.0,20.45, 32.90, 21.1, 23.3} 

	sort.Float64s(speed)

	fmt.Println(speed)

	sort.Sort(sort.Reverse(sort.Float64Slice(speed)))
	fmt.Println(speed)




	

	// fmt.Println(findMinMax(arr))


// result:=rotateleft(arr, 2)

// fmt.Println(result)
// deletefast(arr,4)

// fmt.Println(arr)

	// fmt.Println(arr[:4])


	// fmt.Println(arr[3:])

	// fmt.Println(append(arr[:3], arr[4:]...))



	// result:=insertAt(arr,5,500)

	// fmt.Println(result)




	// traversal slice

	// arr:=[]string{"Sharif","Adiyat","Mamun", "Hasan", "Ashia","Shafi"}

	
					
	// reverse(arr)

	// fmt.Println(arr)





// one:=arr[3:]
// two:=arr[4:]

// fmt.Println(one)
// fmt.Println(two)


// copy(one, two)

// fmt.Println(arr)


	// fmt.Println(arr)

	// one:=arr[2:]
	// fmt.Println(arr)
	// two:=arr[3:]
	// fmt.Println(one)
	// fmt.Println(two)

	// arr=append(arr, "")
	// three:=copy(two, one)
	// fmt.Println(arr)

	// fmt.Println(three)




// for index, value:=range arr{
// 	fmt.Printf("Index: %d Value: %s \n", index, value)
// }


	// for i:=0;i<arr.length;i++{
	// 	fmt.Println(arr[i])
	// }


}

// driver code

// arr:=[]int{1,2,3,4,5,6,7,8,9,10}

// 	fmt.Println(arr)
// 	reverse(arr)
// 	fmt.Println(arr)

// reverse slice

// func reverse (nums []int){

// 	i, j:=0,len(nums)-1

// 	for i<j{

// 		nums[i], nums[j]=nums[j], nums[i]
// 		i++
// 		j--
// 	}
// }
