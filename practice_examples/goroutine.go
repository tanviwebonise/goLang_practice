package main 

import ("fmt")

func getInside(value string) {
	for i:= 0; i < 5 ; i++ {
		fmt.Println(value)
	}
}

func printName(name string) {
	fmt.Println(name)
}

func main() {
	go getInside("Riya")
	go printName("Galaxy")
	go getInside("Tanvi")
	var input string
	fmt.Scanln(&input)
}


