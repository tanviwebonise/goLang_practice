package main 

import ("fmt" ; "errors"; "math")

func getSquareRoot(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	return math.Sqrt(value), nil
}

type negativeValueError struct {
    value  float64
    error string
}

func getCalculate(value float64) (float64, error){
	if value <= 0 {
		return 0, &negativeValueError{value, "negative value"}
	}
	return 10/value, nil
}

func (negativeValueError *negativeValueError) Error() string {
    return fmt.Sprintf("%f : %s", negativeValueError.value, negativeValueError.error)
}

func main() {
	_, errorMessage := getSquareRoot(-1)
	fmt.Println("Error : ", errorMessage)

	_, negativeValueError := getCalculate(-3)
	fmt.Println("Error : ", negativeValueError)
}