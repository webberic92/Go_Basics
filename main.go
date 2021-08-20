package main

import "fmt"

func main() {

	//Loop
	// x := 0
	// for x < 5 {
	// 	fmt.Println("Value is :", x)
	// 	x++
	// }

	//Loop through array.
	names := []string{"me", "them", "you", "us"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	for index, value := range names {
		fmt.Printf("the values at index %v is %v \n", index, value)

	}

}
