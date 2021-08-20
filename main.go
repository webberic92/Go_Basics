package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	//Searching a string, returns true.
	greeting := "hello errrybody"
	fmt.Println(strings.Contains(greeting, "hello"))

	//returns new string with replacement.
	greeting2 := strings.ReplaceAll(greeting, "hello", "hi")
	fmt.Println(greeting2)

	// To uppercase
	fmt.Println(strings.ToUpper(greeting))

	//Shows position where string occurs.
	fmt.Println(strings.Index(greeting, "ll"))

	//Split example
	fmt.Println(strings.Split(greeting, " "))

	ages := []int{2, 6, 3, 5, 4}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 6)
	fmt.Println(index)

	// can also sort and search strings with sort.SearchStrings(nameOFVariable, "stringyourlookingfor")
}
