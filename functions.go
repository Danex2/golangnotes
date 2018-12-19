package main

import (
	"fmt"
)

/*
 - Functions can take zero or more arguments
 - Variables are still defined the same way with the type after it
 - The return value needs to have a type, comes right after declaring the function arguments
*/

func addNums(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(addNums(6, 12))
}

/*
- If there are function arguments that share a type then the type can be omitted
*/
func someStrings(d, z string) string {
	return d + z
}

// When returning multiple values they must be wrapped in brackets
func multiReturn(x, y, z int) (int, int, int) {
	return x, y, z
}

/*
	- Return values can be named, they should be given a descriptive name to explain what they're returning
	- Can also just type 'return' and it'll return all of the named values but should only be used in really
	short functions as they can harm readability in longer functions
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
