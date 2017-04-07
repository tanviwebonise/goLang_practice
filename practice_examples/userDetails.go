package main

import "fmt"

const userName string = "Tanvi"

func main() {
	const userAddress = "Mumbai"
	var userPin, userId int = 10, 1 
	fmt.Println(userName, "stays at", userAddress)
	fmt.Println(userName, "with id", userId, "and pin", userPin)

	userProgrammingSkills := [3]string{"Java", "Groovy", "Go"}
	if userId == 1 {
		for iterator := 0 ; iterator <= len(userProgrammingSkills)-1 ; iterator++ {
			userSkill := userProgrammingSkills[iterator]
			fmt.Println("userSkill", userSkill)
		}
	}

	switch (userId) {
		case 1 : fmt.Println("Address is :", userAddress)
		case 2 : fmt.Println("second users Address")
		default: fmt.Println("No user found")
	}

}
