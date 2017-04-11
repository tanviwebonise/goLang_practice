package main

import "fmt"

//pass by value
func zero(val int) int {
	val = 0
	return val
}

//pass by reference
func getValueWithPointer(val *int) {
	*val++
}

func main() {
	i := 1
	fmt.Println("initial value of i:", i)
	value := zero(i)
	fmt.Println("value without passing reference:", value)
	getValueWithPointer(&i)
	fmt.Println("value after passing reference:", i)
}