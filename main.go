package main
import "fmt"

<<<<<<< HEAD
func main(){

	// age:=40
	// fmt.Println(age)
	// var price float64=63.56

	// fmt.Println(price)
	// fmt.Println("Hello Go lang ")

	// var active bool=true

	// fmt.Println(active)
	// const MYAPPNAME string="My API"
	// fmt.Println(MYAPPNAME)

	// fmt.Printf("Age: %d, Price: %f, Active: %t, app name: %v", age, price, active, MYAPPNAME)


// 	var name string

// 	fmt.Print("Enter Your Name ")
// 	fmt.Scanln(&name)
// 	fmt.Println("Hello", name)


// value:=getDay(4)
// fmt.Println(value)



fmt.Println(calculate(40,50))


}



func calculate(a , b int) (int, int){
	return a+b, a-b
}


// func getDay(day int) string{
// switch day{
// case 1: return "Saturday"
// case 2: return "Sunday"
// case 3: return "Monday"
// case 4: return "Tuesday"
// case 5: return "Wednesday"
// case 6: return "Thursday"
// case 7: return "Friday"
// default: return "Unknown"
// }
// }
=======

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

>>>>>>> 79f8fca1b5150c19e2f523128bfed05fe1229326
