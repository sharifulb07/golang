package main 
import "fmt"


type base struct{
	num int 
}

func (b base) describe()string{
	return fmt.Sprintf("Value of Base is %v: ", b.num)
}

type container struct{
	base
	str string
}


func main(){
	co:=container{
		base:base{
			num:1,
		},
		str:"some text here now ",
	}
	fmt.Printf("num %v str %v", co.num, co.str)

	type describer interface{
		describe()string
	}

	var c describer=co 
	fmt.Println(c.describe())
}