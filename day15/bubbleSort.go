// bubble sort
package main

import "fmt"

// ascending order
func bubbleSort(nums[]int){
	n:=len(nums)

	for i:=0;i<n-1;i++{
		swap:=false 

		for j:=0;j<n-i-1;j++{ 

			if nums[j]>nums[j+1]{

				nums[j], nums[j+1]=nums[j+1], nums[j]
				swap=true 
			}

		}

		if(!swap){
			break
		}
		
	}

}

// descending order

func bubbleSortDescending(nums[]int){
	n:=len(nums)

	for i:=0;i<n-1;i++{
		swap:=false 

		for j:=0;j<n-i-1;j++{ 

			if nums[j]<nums[j+1]{

				nums[j], nums[j+1]=nums[j+1], nums[j]
				swap=true 
			}

		}

		if(!swap){
			break
		}
		
	}

}
// generic type bubble sort

func bubbleSortGeneric [T int| string | float64] (nums[]T){
	n:=len(nums)

	for i:=0;i<n-1;i++{

		swap:=false 
		for j:=0 ;j<n-i-1;j++{
			if nums[j]>nums[j+1]{
				nums[j], nums[j+1]=nums[j+1], nums[j]
				swap=true
			}
		}

		if !swap {
			break
		}
	}


}
// bubble sort by using struct 
type Student struct{
	Name string
	Grade float64


}


func bubbleSortByStruct(students [] Student){
	n:=len(students)

	for i:=0;i<n-1;i++{
		swap:=false 

		for j:=0;j<n-i-1;j++{
			if students[j].Grade>students[j+1].Grade{
				students[j], students[j+1]=students[j+1], students[j]
				swap=true
			}
		}

		if(!swap){
			break 
		}
	}
}



func main(){
	nums:=[]int{1,2,5,4,8,9,7,32,4,5,8,4,8,4,60,54,4,6,8,79,44,55,6,0}
	names:=[]string{"Mango", "Banana", "Lichi", "Owds", "Strawberry"}

	students:=[]Student{{"sharif",50},{"adiyat",40},{"mohammad",30},{"humayon",20},{"sakib",10}}

	fmt.Println(nums)
	
	bubbleSortDescending(nums)
	fmt.Println(nums)
	
	bubbleSortGeneric(nums)
	
	fmt.Println(nums)
	
	bubbleSortGeneric(names)
	fmt.Println(names)

	bubbleSortByStruct(students)

	fmt.Println(students)



	

	
}