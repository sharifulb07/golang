package main

func main() {
	// basic variable
	// var age int
	// age=35
	// fmt.Println(age)

	// declare+ assign together

	// var name string="Shariful Islam"
	// fmt.Println(name)

	// Type inference (Go Detect Types)

	// var city = "Khulna"
	// var score = 100

	// fmt.Println(city,score)

	// short variable declaration or for best practice

	// language:="Go Lang"
	// version:="1.25.6"

	// fmt.Println(language, version)

	// [Not allowed outside functions]

	// Multiple variable detection

	// same type
	// var x,y,z int=10,20,30

	// fmt.Println("X: ",x, "Y: ",y, "Z: ",z)

	// multiple type variable

	// var (
	// 	name string="Shariful"
	// 	age int=36
	// 	active bool=true
	// )

	// fmt.Println("My name is ", name, ". I am ", age, "years old. I am also active go lang learner that's", active)

	// short form variable declaratoin in go lang

	// a,b,c:=10,20, "GO"

	// fmt.Printf("Variable Type: %T and Value: %v ",a,a, )
	// fmt.Printf("Variable Type: %T and Value: %v ",b,b, )
	// fmt.Printf("Variable Type: %T and Value: %v",c,c, )

	// Zero Values -> Default Values

	// var name string
	// var score float64
	// var speed float32
	// var weight int

	// var isAble bool

	// fmt.Println(name, score, speed, weight, isAble)

	// constants in go
	// [Constants cannot be changed]

	// const pi float64=3.1416

	// fmt.Println(pi)

	// get error for change the value of pi
	// pi=3.14

	// fmt.Println(pi)

	// multiple variable constants

	// const (
	// 	lang     = "Go"
	// 	verision = 1.25
	// 	author   = "Shariful Islam"
	// )

	// fmt.Println(lang, verision, author)

	// untyped constants

	// const x = 200
	// fmt.Println(x)

	// typed constants
	// const y int=500
	// fmt.Println(y)

	// constants expression

	// const (
	// 	A=10
	// 	B=20
	// 	C=A+B
	// )

	// fmt.Println(A,B,C)

	// iota with values

	// const (
	// 	KB=1<<(10*iota)
	// 	MB
	// 	GB
	// )

	// fmt.Println(KB, MB, GB)

	// swap variables mini project

	// var x, y int=10,20

	// fmt.Println("X: ",x)
	// fmt.Println("Y: ", y)

	// x=y
	// y=x
	// fmt.Println("After Swapping ")
	// fmt.Println("X: ",x)
	// fmt.Println("Y: ", y)

	// convert Celsius to Farenheit

	// fmt.Println(Converter(100.00))

	// calculate Area of a Circle

	// const PI=3.1416
	// radius:=30
	// Area:=PI*float64(radius)*float64(radius)

	// fmt.Println(Area)

	// user Profile Variables

	// var (
	// 	name="Shariful Islam"
	// 	age=35

	// 	isActiveLearner=true
	// )

	// fmt.Println(name, age, isActiveLearner)

}

// func Converter(Celsius float64)float64{
// 	return (Celsius+9/5)+32
// }