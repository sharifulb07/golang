package main 
import (
	"fmt"
	"math"
)

// interface
type geomatry interface{
	area()float64
	perim()float64
}

// rect struct
type rect struct{
	width, height float64
}

// circle struct

type circle struct{
	radius float64
}

func (r rect) area()float64{
	return r.width*r.height
}

func (r rect) perim() float64{
	return 2*r.width+2*r.height
}

func (c circle)area()float64{
	return math.Pi*c.radius*c.radius
}

func (c circle) perim()float64{
	return 2*math.Pi*c.radius
}

func measure (g geomatry){
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectCircle(g geomatry){

	if c, ok:=g.(circle); ok{
		fmt.Println("Radius of a circle is : ", c.radius)
	}
}

func main(){
	r:=rect{width: 20, height: 30}
	c:=circle{radius: 40}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}