package main

import "fmt"


func main(){
	nums:=[]int{4,5,7,8,9,6,3,50}

for _, num:=range nums{
	fmt.Println("Value of Slice:", num)
}

for i, num:=range nums{

	fmt.Printf("Key: %T and Value: %v",i, num)
}


	values:=map[string]string{"a":"apple", "b":"banana", "c":"Cowliflower"}

for key, v:=range values{

	fmt.Printf(key, v)

}


}
