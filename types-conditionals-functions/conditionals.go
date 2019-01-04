package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	fmt.Println(checkAge(31))
}

/*
- Not much to explain here, same as any other language. Still no brackets cause that's just how we rockin
*/

func checkAge(age int) string {
	if age < 30 {
		return "Sorry you are much too young to be here!"
	}
	return "Welcome old man!"
}

// if short statements
// don't rly understand these, come back to it

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// Switch statements, not much different here
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
