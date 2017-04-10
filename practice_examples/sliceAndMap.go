package main

import "fmt"

func main() {
	books := make([]string, 3)
	books[0] = "Programming Pearls"

	books := append(books, "Design Patterns", "C Primer Plus", "Java Puzzlers")
	fmt.println("books:", books)

	sliceThree := books[1:len(books)-1]
	sliceOne := books[len(books)-1:]
	fmt.println("sliceThree:", sliceThree)
	fmt.println("sliceOne:", sliceOne)

	limitedBooks := copy(limitedBooks, sliceThree)
	fmt.println("Copy slice:", limitedBooks)
}