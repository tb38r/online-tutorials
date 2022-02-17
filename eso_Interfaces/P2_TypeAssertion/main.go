package main

import (
	"fmt"
	"math"
)

//Type Assertion
//Empty Interfaces
//Type Switch

type shape interface {
	area() float64
	perimeter() float64
}

type polygon interface {
	getSides() int
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
	sides         int
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}
func (r rectangle) getSides() int {
	return r.sides
}

func Calculate(s interface{}) {

	// rect := s.(rectangle) //this is the type assertion (makimg assertion of the underlying
	//type of the passed in Interface)
	//we can also assert an interface (that shares a method signature
	// with the type)
	//we can now call all the methods of the rectangle
	// fmt.Printf("I am of Type %T\n", rect2) //rectangle type is asserted as it uses the
	//getSides() which is polygon's Interface signature

	//error handling, in case wrong type is asserted

	if rect, ok := s.(rectangle); ok {
		fmt.Println(ok)
		fmt.Println(rect.perimeter())
		fmt.Printf("My type is %T\n", rect)
		fmt.Println()
	}

	//Type Switch statements

	switch x := s.(type) {
	case circle:
		fmt.Printf("My type is %T\n", x)
	case rectangle:
		fmt.Printf("My type is %T\n", x)
	case int:
		fmt.Printf("My type is %T\n", x)
	default:
		fmt.Printf("My type is %T\n", x)
	}

}

//ordinarily couldn't call getSides() within the shape interface as it doesn't satisfy the method
//signatures. This is where type assertions come in...

//Type Assertion lets us extract the underlying concrete type of the value that gets passed in
//(in this case a rectangle) then calls the methods

//Type : rectangle

/* Value :
rectangle {
	width : 5
	height : 8
	sides : 4
} */

func main() {
	r := rectangle{5, 8, 4}
	c := circle{radius: 3}
	n := 2

	Calculate(r)
	Calculate(c)
	Calculate(n)

}
