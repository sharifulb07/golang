package main
import "fmt"

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