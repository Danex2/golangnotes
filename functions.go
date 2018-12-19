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
