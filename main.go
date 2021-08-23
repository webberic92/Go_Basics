package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Fizz = 3, Buzz = 5")

	for i := 0; i <= 100; i++ {
		finalString := ""

		if i%3 == 0 {
			finalString = "Fizz"

		}
		if i%5 == 0 {
			finalString += "Buzz"

		}

		if i%3 != 0 && i%5 != 0 {
			finalString = strconv.Itoa(i)

		}
		fmt.Println(finalString)

	}

}
