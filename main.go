package main

import "fmt"

func main() {

	//Strings

	//Explicitly typing
	var nameOne string = "mario"
	//Infering typing
	var nameTwo = "luigi"
	//initialize emtpy string
	var nameThree string
	//Short hand for initializing variable.Infers type.
	//Cant do this OUTSIDE function.
	nameFour := "yoshi"

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	//Can specify integer size with int8,16,32,64 etc..
	var numOne int8 = 25
	var numTwo int8 = -128

	//UINT Can NOT be negative, allows higher then counterpart since no longer accepting minuses.
	var numThree uint8 = 128
	fmt.Println(numOne, numTwo, numThree)

	//Floats have decimal points.
	var scoreOne float32 = 99.99
	var scoreTwo float64 = 99.999999999
	//Comes out as 100
	scoreThree := 99.9999999999999999999999999999999999999999

	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
