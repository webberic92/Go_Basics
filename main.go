package main

import "fmt"

func main() {

	age := 35
	name := "ducky"

	//at end of print does not add on a new line.
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("newline \n")

	//Println adds new line, plus example of syntax for variables.
	fmt.Println("Hello again")
	fmt.Println("goodbye")
	fmt.Println("my name is ", name, "my age is ", age)

	//PrintF (formatted Strings) %_ is format specifier
	//%q is quoutes so wont work on integers
	//%v is variable
	//%T will tell you type for example it will say float or int.
	//%f is float and can limit by adding number example %0.1f
	fmt.Printf("Formated String : my age is %v and my name is %v \n", age, name)

	//SprintF saves formatted strings.
	var str = fmt.Sprintf("Formated String : my age is %v and my name is %v", age, name)
	fmt.Println("SAVED STRING :", str)
}
