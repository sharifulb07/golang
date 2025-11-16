


package main

import (
    "fmt"
    
)

func worker(ch chan string){
    ch<-"Task Completed"
}






func main() {

    ch:=make(chan string)

    go worker(ch)

    msg:=<-ch

    fmt.Println(msg)

    
}



// go printMessage()

// time.Sleep(3*time.Second)
// fmt.Println("Main finished")



// Goroutines example 

// func printMessage(){
//     fmt.Println("Goroutine is started......")
// }


// process(ConsoleLogger{})



// type Logger interface{
//     Log(msg string)
// }


// type ConsoleLogger struct{}

// func (ConsoleLogger) Log(msg string){
//     fmt.Println("This is the message: ", msg)
// }

// type FileLogger struct{
//     file *os.File
// }


// func (f FileLogger) Log(msg string){
//     f.file.WriteString(msg+"\n")
// }

// func process(l Logger){
//     l.Log("Process is started.............")
// }




// var x interface{}
//     x=10
//     fmt.Println(x)
//     x="Shariful Islam"
//     fmt.Println(x)
//     x=[]int{12,14,15,17}
//     fmt.Println(x)
    


// type shape interface{
//     Area() float64
// }

// type Rectangle struct{
//     width, height float64
// }

// func (r Rectangle) Area()float64{
//     return r.width*r.height
// }


// type Circle struct{
//     radius float64
// }

// func (r Circle) Area()float64{
//     return 3.1416*r.radius*r.radius
// }

// func printArea(s shape){
//     fmt.Printf("This Area:%v \n", s.Area())
// }

// r:=Rectangle{width:10, height:30}
// c:=Circle{30}
// printArea(r)
// printArea(c)

// err:=process()
// if err!=nil{
//     fmt.Println("error is:", err)
//     return
// }

// fmt.Println("File Processed successfully")




// func process()error{
//     f, err:=os.Open("file.txt")
//     if err!=nil{
//         return err
//     }
//     defer f.Close()

//     data:=make([]byte, 100)
//     _, err =f.Read(data)

//     if err!=nil{
//         return err
//     }
//     return nil
// }
// test()


// func test(){
//     defer fmt.Println("Last")
//     fmt.Println("First")
// }


// my:=Person{name:"Shariful islam", age:36}
// my.greet()
   

// Methods (Functions with Receiver)
// type Person struct{
//     name string
//     age int
// }

// func (p Person) greet(){
//     fmt.Printf("I am %v and i am %v years old", p.name, p.age)
// }

// c:=counter()
// fmt.Println(c())
// fmt.Println(c())
// fmt.Println(c())
// fmt.Println(c())
// fmt.Println(c())



// a closure function 
// func counter() func() int{
//     count:=-1
//     return func()int{
//         count++
//         return count
//     }
// }

// describer:=multiplier(10)
// fmt.Println(describer(50))
// f:=func(a int, b int)int {
//     return a+b
// }

// fmt.Println(f(20,50))



//     var fullname, country string
// var age int
// var gpa float32

// fmt.Print("Enter Your Info: ")
// fmt.Scan(&fullname, &age, &country, &gpa)

// fmt.Printf("%v is a student \n", fullname)
// fmt.Printf("%v is %v years old. \n", fullname, age)
// fmt.Printf("%v has got GPA %v \n", fullname, gpa)
// fmt.Printf("%v is from %v\n", fullname, country)

  // Function Returning Another Function
//   func multiplier (x int) func(int) int{
//     return func(n int)int {
//         return x*n
//     }

// }


// package main

// import "fmt"

// import "rsc.io/quote"

// func main(){
// 	fmt.Println("Hello world")
// 	fmt.Println(quote.Go())
// }


// variadic func with unlimited arguments

// func sum(nums ...int)int {
//     total:=0

//     for _, value:= range nums{
//         total+=value
//     }
//     return total 
// }



// function with named return values

// func rectangle(w, h float64)(area, peremiter float64){
//     area=w*h
//     peremiter=2*(w+h)
//     return
// }


// func divide(a, b float64)( float64, error){
//     if b==0{
//         return 0, fmt.Errorf("divided by zero")
//     }
//     return a/b, nil
// }


// func add(a int, b int)int{
//     return a+b
// }


// func varSec(){


//     fmt.Println("Be Cool Man: ")

// name:="sharif"

// result:=fmt.Sprintf("This is %v", name)
// fmt.Println(result)


// loop in map



    // for i, c :=range "Golang"{
    //     fmt.Println(i, string(c))
    // }

    // mymap:=map[string]int{"a":1, "b":3,"c":5}


    // for i,v :=range mymap{
    //     fmt.Println(i, v)
    // }




    // using loop in array

    // arr:=[]int{1,11,111,1111,2,22,222,2222}

    // for index, value:=range arr{
    //     fmt.Println(index, value)
    // }


    // loop statement

    // for i:=0;i<=5;i++{
    //     fmt.Println(i)
    // }


    // while loop
    // j:=6

    // for j<=20{
    //     fmt.Println("Value",j)
    //     j++
    // }


    // switch statement

    // score:=65

    // switch{
    // case score>=80:
    //     fmt.Println("A Grade")
    // case score>=70:
    //     fmt.Println("B Grade")
    // case score>=60:
    //     fmt.Println("C Grade")
    // }


    // fruit:="apple"

    // switch fruit{
    // case "mango":
    //     fmt.Println("Mango")
    // case "apple":
    //     fmt.Println("Apple")
    // default:
    //     fmt.Println("Unknown Fruit")
    // }

    // switch day:=2; day{
    // case 1,2,3,4,5:
    //     fmt.Println("Weekday")
    // case 6,7:
    //     fmt.Println("Weekend")
    // default:
    //     fmt.Println("Invalid Number ")
    // }


    // if conditional statement

// age:=22

// if age>=21{
//     fmt.Println("Adult")
// }else{
//     fmt.Println("Minor")
// }


// if num:=11; num%2==0{
//     fmt.Println("Even Number")
// }else{
//     fmt.Println("Odd Number")
// }





// rune used for unicode in go lang 
    // var r rune = 'গ'
    // fmt.Println(r)




    // pointer in  go lang

    // var x=50
    // var y=&x
    // var z=&y
    // fmt.Println(y)
    // fmt.Println(x)
    // fmt.Println(z)







    // map in go lang

    // names:=map[string]string{
    //     "one":"sharif",
    //     "two":"adi",
    //     "three":"ashia",
    // }

    // for _, value:= range names{
    //     fmt.Println(value)
    // }
    // fmt.Println(names)





    // array in go 

// numbers :=[]int {40,50,60,70}
// fmt.Println(numbers)






    // var arr []int=[]int{1,2,3,4,5,6,7}
    // for index, value:=range arr{

    //     fmt.Println(index, value)
    // }





    
    // constant in go lang


    // const (
    //     A=iota
    //     B
    //     C
    //     D
    //     E
    //     F
    // )

    // fmt.Println(A,B,C, D, E, F)







    // const PI=3.14159
    // const name="shariful"

    // fmt.Println(PI, name)



    // variable in go lang
    
    // var (
    //     username string="cureone"
    //     isLoggin bool=false
    //     score int=32
    // )

    // fmt.Println(username, isLoggin, score)




    // var friend string
    // var price int
    // var confident bool

    // friend="Adiyat"
    // price=07
    // confident=true
    // fmt.Println(friend, price, confident)

// var x, y, z=10,20,30
// fmt.Println(x,y,z)

    // name :="Shariful Islam"
    // age :=37
    // city :="Khulna"

    // fmt.Println(name,"who is", age, "years old lives in", city, "City")
// }
