package main

import "fmt"

func main() {
	elements := map[string]string {
		"H" : "Hydrogen",
		"B" : "Boron",
		"C" : "Carbon",
		"O" : "Oxygen"}
	fmt.Println("Initial map elements :", elements)

	delete(elements, "H")
	fmt.Println("After deleting first element :", elements)

	_, isElementExist := elements["H"]
	fmt.Println("Is hydrogen exists? ", isElementExist)

	value, status := elements["B"]
	fmt.Println("Element :", value, ", status:", status)
}