package main

import ("fmt" ; "math")

type Shape interface {
	getPerimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	height, width float64
}

func (rectangle *Rectangle) getPerimeter() float64 {
	return 2*rectangle.width + 2*rectangle.height
}

func (circle *Circle) getPerimeter() float64 {
    return 2 * math.Pi * circle.radius
}

func getShape(shape Shape) {
	fmt.Println("Perimeter:", shape.getPerimeter())
}

func main() {
	circle := Circle{5}
	rectangle := Rectangle{10, 20}
	fmt.Println("Perimeter of Circle : ", circle.getPerimeter())
	fmt.Println("Perimeter of Rectangle : ", rectangle.getPerimeter())
	getShape(&circle)
	getShape(&rectangle)
}