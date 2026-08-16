package main

import(

	"fmt"
)





func main() {

	
var p *int
change(&p)

fmt.Println(p)


	
}

func change(p *int){
	*p=5001
}


// type User struct {
// 	ID int  `json="id"`
// 	Name string `json="name"`
// 	Email string `json="email"`
// }


// u1:=User{
// 		ID: 12,
// 		Name:"sharif",
// 		Email:"sharif",
// 	}

// 	data, _:=json.Marshal(u1)

// 	fmt.Println(string(data))






// struct tag
// real banking example


// user := NewAccount("Sharif")
// 	user.Deposit(500)
// 	user.Deposit(700)
// 	user.Deposit(200)

// 	user.withdraw(600)
// 	bl := user.GetBalance()

// 	fmt.Println(bl)
// type Account struct {
// 	Name    string
// 	Balance float64
// }

// func NewAccount(name string) *Account {
// 	return &Account{
// 		Name:    name,
// 		Balance: 0,
// 	}
// }

// // deposit
// func (a *Account) Deposit(amount float64) {
// 	a.Balance += amount
// }

// // withdraw

// func (a *Account) withdraw(amount float64) bool {

// 	if amount > a.Balance {
// 		return false
// 	}
// 	a.Balance -= amount
// 	return false
// }

// // getbalance

// func (a Account) GetBalance() float64 {
// 	return a.Balance
// }

// pointer

// type Person struct{
// 	Name string
// 	Email string
// }

// func newUser(name, email string) *Person{
// 	return &Person{
// 		Name:name,
// 		Email:email,
// 	}
// }

// p1:=Person{"hasan", "hasan@gmail.com"}
// fmt.Println(p1)

// result:=newUser("hamim", "moon@gmail.com")
// fmt.Println(result.Email)
// fmt.Println(result.Name)
// fmt.Println(newUser("Mehrab", "mehrab@gmail.com"))

// type User struct {
//     Name string
//     Age  int
// }

//  user1 := User{
//         Name: "Sharif",
//         Age:  35,
//     }

//     user2 := &user1

//     user2.Name = "Rahim"

//     fmt.Println(user1.Name)
//     fmt.Println(user2.Name)
//     fmt.Println(user1.Name)

// u:=User{"sharif"}
// fmt.Println(u.Name)
// u.changeName("adi")
// fmt.Println(u.Name)

// type User struct{
// 	Name string
// }

// func (u *User) changeName(name string){
// 	u.Name=name
// }

// change original value with pointer
// x := 10
// 	p := &x

// 	fmt.Println(x)
// 	fmt.Println(p)
// 	fmt.Println(*p)
// 	*p=500
// 	fmt.Println(*p)

// x := 10
// p := &x

// fmt.Println(x)
// fmt.Println(p)
// fmt.Println(*p)

// struct

// type Person struct {
// 	Name  string
// 	Age   int
// 	Email string
// }

// func (u User)changeName(name string){

// 	u.Name=name
// 	fmt.Println(u.Name)
// }

// type User struct {
// 	Name    string

// }

// type Address struct {
// 	City    string
// 	Country string
// }

// user:=User{Name: "Shariful Islam"}
// fmt.Println(user.Name)
// user.changeName("Hunaif Sharif")
// fmt.Println(user.Name)

// p1 := user{
// 		Name: "Sharif",
// 		Address: Address{
// 			City:    "Khulna",
// 			Country: "Bangladesh",
// 		},
// 	}

// 	fmt.Println(p1)

// modify value here
// user:=Person{"Adiyat", 7, "adi@gmail.com"}
// 	fmt.Println(user)
// 	user.Name="Ashia khatun"
// 	fmt.Println(user)

// var user Person

// fmt.Printf("%+v\n", user.Name)

// Positional initialization

// user:=Person{"Shariful islam", 28, "shariful@gmail.com"}
// fmt.Println(user)

// Named Fields

// user := Person{
// 	Name:  "Shariful Islam",
// 	Age:   36,
// 	Email: "shariful@gmail.com",
// }

// fmt.Println(user)
// fmt.Println(user.Name)
// fmt.Println(user.Email)
// fmt.Println(user.Age)

// type Person struct{
// 	Age int
// 	Active bool
// }

// var p Person

// fmt.Println(unsafe.Sizeof(p))

// map examples

// // name is nil
// var names map[string]string
// 	fmt.Println(len(names))

// type Point struct{
// 	X int
// 	Y int
// }

// m:=map[Point]string{
// 	Point{X:0, Y:0}:"Zero Axis",
// 	Point{X:3, Y:3}:"Multiply Axis",
// }

// m[Point{X: 4, Y:8}]="Point A"
// m[Point{X: 12, Y:32}]="Point B"

// key, value:=m[Point{X:4, Y:8}]

// fmt.Println(key)
// fmt.Println(value)

// fmt.Println(m)

// users := make(map[string]string)

// 	users["sharif"] = "good"
// 	users["ashia"] = "bad"
// 	users["adiyat"] = "very bad"

// 	delete(users, "sharif")

// 	fmt.Println(users)

// users := make(map[string]string)

// 	users["sharif"] = "good"
// 	users["ashia"] = "bad"
// 	users["adiyat"] = "very bad"

// 	for key, value := range users {
// 		fmt.Print(key, "-----")
// 		fmt.Println(value)
// 	}

// info := map[string]int{
// 		"Rahim": 20,
// 		"Karim": 39,
// 		"Kum":   50,
// 	}

// fmt.Println(info["Rahim"])

// for index, value:=range info{
// 	fmt.Println(index)
// 	fmt.Println(value)
// }
// 	fmt.Println(info)

// num:=[4]int{}
// fmt.Println(num)

// a := make([]int, 3, 5)
// fmt.Println(a[2])
// a=append(a, 50)
// a=append(a, 20)
// a=append(a, 30)
// a=append(a, 60)
// a=append(a, 60)
// a=append(a, 60)
// a=append(a, 60)
// fmt.Println(a)

// taking input number or string or slice or array

// var num1 int
// var num2 int

// fmt.Println("Enter second number ")
// fmt.Scanln(&num2)

// fmt.Println("Enter first number ")
// fmt.Scanln(&num1)

// fmt.Println(num1)
// fmt.Println(num2)

// a:=[]int{2,5,9}
// b:=make([]int, len(a))

// copy(b, a)

// fmt.Println(b)
// fmt.Println(a)

// array := [5]int{10, 20, 30, 40, 50}

// slice := array[1:4]

// fmt.Println(slice)

// numbers:=[3]int{129, 30,4Println(p)
// 	fmt.Println(numbers)
// 	change(&numbers)
// 	fmt.Println(numbers)

// func  change(num *[3]int)  {

// 	num[0]=500

// }

// a:=[]int{10,20,30,50,60}
// b:=a
// fmt.Println(a[0])
// a[0]=200
// fmt.Println(b[0])
// fmt.Println(a[0])

// numbers:=[...]int{5,6,1,8,9}

// 	fmt.Println(numbers[3])
// 	numbers[0]=10
// 	fmt.Println(numbers[0])
// numbers := [5]int{1, 2, 3, 4, 5}
// 	fmt.Println(numbers[0])
// 	fmt.Println(numbers[1])
// 	fmt.Println(numbers[2])
// 	fmt.Println(numbers[3])
// 	fmt.Println(numbers[4])
// var numbers [5] int
// fmt.Println(numbers[0])
// fmt.Println(numbers[1])
// fmt.Println(numbers[2])
// fmt.Println(numbers[3])
// fmt.Println(numbers[4])

// text := "বাংলা"

// for i, ch:=range text{
// 	fmt.Printf("Index%d \n, unicode%c \n", i, ch)
// }
// fmt.Println(text)
// fmt.Println(len(text))

// fmt.Println("Hello Wold")

// slc:=[]rune{'♛', '♠', '♧', '♡', '♬'}

// for i, value:=range slc{

// 	fmt.Printf("Character%c, Unicode%U, Position %i", value, value, i)
// }