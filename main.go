package main
import "fmt"


func sum(nums ...int )int {
total:=0

for _, num:=range nums{
total+=num
}

return total

}


func main(){

fmt.Println(sum(1,2))
fmt.Println(sum(2,3,5,4,6))

nums:=[]int{5,8,6,9,7,10}

fmt.Println(sum(nums...))

}

