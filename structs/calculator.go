package structs

import "math"


type Shape interface {
	Area() float64
}


type Rectangle struct{
	width, height float64
}

type Circle struct {
	radius float64
}



func Perimeter(rectangle Rectangle) float64{
	return 2 * (rectangle.width + rectangle.height)
}

func (r Rectangle) Area() float64{
	return r.height * r.width
}

func (c Circle) Area() float64{
	return math.Pi * c.radius * c.radius
}