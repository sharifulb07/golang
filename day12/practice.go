package main 
import "fmt"

func main() {

	fmt.Println("Hello, World!")
	fmt.Println("The result of addition is ", processOp(10, 20, add))
}


func processOp(a int, b int, op func (x int, y int)int) int {

	fmt.Println("Process Operation works very well now ")
	return op(a, b)
	
}



func add(a int, b int) int {

	return a + b

}