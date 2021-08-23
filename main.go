package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Fizz = 3, Buzz = 5")

	for i := 0; i <= 100; i++ {

		if i%3 == 0 {
			fmt.Printf("fizz")
		}
		if i%5 == 0 {
			fmt.Printf("buzz")
		}

		if i%3 != 0 && i%5 != 0 {
			fmt.Printf(strconv.Itoa(i))
		}
		fmt.Println("")

	}

	//Case switch example
	// for i := 1; i <= 100; i++ {
	// 	switch {
	// 	   case i%15==0:
	// 	   fmt.Println("FizzBuzz")
	// 	   case i%3==0:
	// 	   fmt.Println("Fizz")
	// 	   case i%5==0:
	// 	   fmt.Println("Buzz")
	// 	   default:
	// 	   fmt.Println(i)
	// 	}
	//  }

}
