package main

import "fmt"

func GetTotalMarksAndAverage(marks []float64) (float64, float64) {
	total := 0.0
	for _, iterator := range marks {
		total += iterator
	}
	return total, (total / float64(len(marks)))
}

func main() {
	marks := []float64 {66, 20, 78, 96, 45}
	total, average := GetTotalMarksAndAverage(marks)
	fmt.Println("Total marks:", total, "average:", average)
}