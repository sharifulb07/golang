package main

func main() {

	// signed integers

	// var a int=10
	// var b int8=-123
	// var c int16=23146
	// var d int32=884622156
	// var e int64=65451232224522111

	// fmt.Println(a,b,c,d, e)

	// unsigned integers

	// var x uint=20
	// var p uint8=68
	// var y uint16=25454
	// var z uint32=965655644
	// var q uint64=6898431312522103132

	// fmt.Println(x,p,y,z,q)

	// floating-points Types

	// var f1 float32=3.25584597496
	// var f2 float64=64693246.054487941564

	// fmt.Println(f1, f2)

	//  bool type variable

	// var isOnline bool = true
	// var isAdmin bool = false

	// fmt.Println(isAdmin, isOnline)

	// string data type

	// var name string ="Shariful Islam"
	// greetings:="Hello!-> "

	// fmt.Println(greetings,name)

	// string length
	// fmt.Println(len(name))
	// fmt.Println(len(greetings))

	// Derived / Composite Data Types

	// var numbers [3]int = [3]int{1, 2, 3}

	// fmt.Println(numbers)
	// fmt.Printf("Type of varible: %T ", numbers)

	// Array Inference

	// ProgramingLanguages:=[...]string{"Go", "Python", "C-Sharp"}

	// fmt.Println(ProgramingLanguages)
	// fmt.Printf("Type of Variable is : %T ", ProgramingLanguages)

	// slice -> Dynamic array
	// nums := []int{10, 20, 30, 40}
	// nums = append(nums, 50)

	// fmt.Println(nums)
	// fmt.Printf("%T", nums)

	// slice from an array

	// arr:=[5]int{1,2,3,4,5}

	// slice:=arr[3:4]
	// fmt.Println(slice)
	// fmt.Println(len(slice))
	// fmt.Println(cap(slice))

	// Map data type

	// student:=map[string]int{"Bangla":90, "Math":40, }
	// fmt.Println(student)
	// student["Science"]=100
	// fmt.Println(student)

	// delete(student,"Bangla")
	// fmt.Println(student)

	// struct data type

	// type User struct{
	// 	Name string
	// 	Age int
	// }

	// user:=User{
	// 	Name: "Shariful islam",
	// 	Age: 10,
	// }

	// fmt.Println(user.Name, user.Age)

	// pointer data type

	// x := 50
	// ptr := &x
	// fmt.Println(ptr)
	// fmt.Println(*ptr)

	// modifying value using pointer

	// *ptr=100
	// fmt.Println(x)

	// special data type

	// var ch rune='A'
	// fmt.Println(ch)
	// fmt.Printf("%c\n", ch)

	// byte data types

	// var b byte='G'
	// fmt.Println(b)
	// fmt.Printf("%c \n", b)

	// interface ->used as dynamic variable

	// var value interface{}

	// value = "shariful islam"

	// fmt.Println(value)

	// value=30
	// fmt.Println(value)

	// Type Assertion

	// var data interface{}="Go Lang"

	// str:=data.(string)

	// fmt.Println(str)

	// Safe type assertion

	// data=20

	// num,ok:=data.(int)
	// fmt.Println(num, ok)

	// Function as a Type

	// var add func(int, int) int

	// add = func(a int, b int) int {
	// 	return a + b
	// }

	// fmt.Println(add(10,80))

	// Channel Concurrency

	// ch := make(chan int)

	// go func() {
	// 	ch <- 500
	// }()

	// fmt.Println(<-ch)




	// type conversion


// var a int=500
// b:=float64(a)
// fmt.Printf("%T",b)



}