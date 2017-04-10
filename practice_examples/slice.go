package main

import "fmt"

func main() {
	books := make([]string, 3)
	books[0] = "Programming Pearls"

	books = append(books, "Design Patterns", "C Primer Plus", "Java Puzzlers")
	fmt.Println("books:", books)

	sliceThree := books[1:len(books)]
	sliceOne := books[len(books)-1:]
	fmt.Println("sliceThree:", sliceThree)
	fmt.Println("sliceOne:", sliceOne)

	limitedBooks := make([]string, len(books))
	copy(limitedBooks, books)
	fmt.Println("Copy:", limitedBooks)
}