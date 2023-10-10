package main

import "fmt"

func vars() {
	fmt.Println("Hello Ninjas!")

	// Strings
	var nameOne string = "Madi"
	nameTwo := "Winnie"
	var nameThree string
	namefour := "youshi"

	nameOne = "Mando"
	nameThree = "Tets"

	fmt.Println(nameOne, nameTwo, nameThree, namefour)

	// Integers
	ageOne := 30
	var ageTwo int = 40
	var ageThree = 20

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var ageFour int8 = 12
	var numTwo int8 = -128
	var numThree uint8 = 25
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 8888.67478
	fmt.Println(ageFour, numTwo, numThree, scoreOne, scoreTwo)

}
