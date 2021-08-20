package main

import "fmt"

func main() {

	//arrays have to be FIXED Length.
	var ages [4]int = [4]int{1, 2, 3, 4}

	//Infers for us from right hand side.
	var ages2 = [4]int{1, 2, 3, 4}

	fmt.Println(ages, len(ages))
	fmt.Println(ages2)

	//String array
	names := [3]string{"me", "you", "them"}
	fmt.Println(names)

	//SLICES can add/remove length.
	//with no value in brackets it creates a SLICE
	var scores = []int{50, 60, 70}
	fmt.Println(scores, len(scores))

	//APPEND KEY WORD FOR SLICES
	scores[2] = 99
	scores = append(scores, 999999)
	fmt.Println("NEW SCORES :", scores, len(scores))

	//Show rangers of  arrays/slices
	//inclusive of first number but not second.
	rangeOne := names[1:3]
	rangeTwo := names[:]
	rangeThree := names[:3]

	fmt.Println(" rangeOne :", rangeOne)
	fmt.Println(" rangeTwo :", rangeTwo)
	fmt.Println(" rangeThree :", rangeThree)

}
