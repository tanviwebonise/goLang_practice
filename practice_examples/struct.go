package main

import ("fmt" ;"math")

type Circle struct {
	radius, x_cordinate, y_cordinate float64
}
//function
func getCircleArea(circle *Circle) float64 {
	return math.Pi * circle.radius * circle.radius
}

//method
func (circle *Circle) getArea() float64 {
	return math.Pi * circle.radius * circle.radius
}

func main() {
	circle := Circle{10, 0, 0}
	circle.x_cordinate = 20
	fmt.Println("Circle radius : ", circle.radius)
	fmt.Println("Circle x-coordinate : ", circle.x_cordinate)
	fmt.Println("Circle area : ", getCircleArea(&circle))
	fmt.Println("Circle area : ", circle.getArea())
}