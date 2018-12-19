package main

import (
	"fmt"
)

/*
- There is only one loop in Go, the for loop
- The for loop isn't only used as a for loop, can be used as a while loop also
- For loops also do not contain brackets
*/
func main() {
	forLoop(1)
}

func forLoop(x int) {
	for x := x; x < 10; x++ {
		fmt.Println(x)
	}
}
