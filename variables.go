package main

import "fmt"

/*
- the keyword var defines a list of variables
- they can be defined in the global scope or inside of a function
- the type always goes at the end of variable declarations
- Any numeric variable declared without a value is given the number 0
- Any boolean declared without true/false is given false
- Any string declared without a value is given an empty string
- Variables can also be declared with the "const" keyword. Like in javascript, anything declared as const can not be redeclared
*/
var name, school, grade string

// if the variable is assigned a value then the type can be omitted
var num1, num2, c, b = 1, 2, true, false

func main() {
	fmt.Println(name, school, grade)
	fmt.Println(num1, num2, c, b)
	shortVar()
}

// inside of a function the short variable declaration can be used ":=", does not work in the global scope. Variables must be assigned with var, func, etc.

func shortVar() {
	j := "Does this work?"
	k := true
	fmt.Println(j, k)
}
