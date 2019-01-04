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
	//forLoop(1)
	//forWhile(1)
	loop2()
}

// regular for loop
func forLoop(x int) {
	for x := x; x < 10; x++ {
		fmt.Println(x)
	}
}

// while loop but still using the keyword `for`
func forWhile(x int) {
	for x < 10 {
		x++
		fmt.Println(x)
	}
}

// the initializer in the for loop can be omitted as long as it is defined somewhere
// you can also drop the increment part of the loop, linting won't let me do it >:(
func loop2() {
	x := 0
	for x < 10000 {
		fmt.Println(x)
	}
}

// and with nothing it is just an infinite loop
func inf() {
	for {

	}
}
