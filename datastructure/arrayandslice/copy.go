package main 
import "fmt"


func main(){

	source:=[]int{1,2,3,4,5,6,7}
	dest:=[]int{8,9,10,11}


	result:=copy(dest, source)

	fmt.Println(result)

	fmt.Println(dest)
	fmt.Println(source)
}