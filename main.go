package main

import (
	"fmt"
)

func main() {

	userInput := ""
	fmt.Println("This application will test if user input is a palindrome...")
	fmt.Println("Please enter string to test palindromity...")

	// Taking input from user
	fmt.Scanln(&userInput)

	fmt.Println("You entered : ", userInput)
	fmt.Println("Lets check to see if {", userInput, "} is a palindrome....")
	fmt.Println("Processing")
	fmt.Println("Processing")
	fmt.Println("Processing")
	fmt.Println("Checking the mainframe")
	fmt.Println("Hacking the blockchain.....")
	fmt.Println("This is a palindrome is : ", isPalindrome(userInput))

}

func isPalindrome(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	for i := range str {
		if str[i] != reversedStr[i] {
			return false
		}
	}
	return true
}
