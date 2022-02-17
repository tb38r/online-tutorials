package main

import(
	"fmt"
	"math"
)

type shape interface{
	area() float64
	perimeter()float64 
}

type rectangle struct{
	width, height float64
}

type circle struct{
	radius float64
}

func (r rectangle)area()float64{
	return r.height* r.width
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (r rectangle) perimeter()float64{
	return 2 * r.width + 2 * r.height
}

func (c circle) perimeter() float64{
	return 2 * math.Pi * c.radius
}

//with each type having an area & perimeter method that return a type float64, they can be 
//said to satisfy the shapes interface 


func Calculate (s shape){
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}


func main(){

r := rectangle {5,8}
c := circle{6}

Calculate (r)
Calculate(c)





}