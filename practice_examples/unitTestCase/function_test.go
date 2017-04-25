package main 

import "testing"

func TestGetTotalMarksAndAverageReturnsValue(t *testing.T) {
	marks := []float64 {66, 20, 45}
	total, average := GetTotalMarksAndAverage(marks)

	if total == 0 {
		t.Errorf("Failed to calculate total of marks")
	}
	if average == 0 {
		t.Errorf("Failed to calculate average of marks")
	}
}