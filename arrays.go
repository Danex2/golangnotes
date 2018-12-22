package main

import "fmt"

/*
- Arrays have a fixed length
- You can add values to an array by setting the index to a value
*/

func main() {
	// Empty array with a length of 2
	var even [2]int

	// Add elements
	even[0] = 2
	even[1] = 4

	e := even
	fmt.Println(e)

	// Or just define an array with all of your elements
	odd := [5]int{1, 3, 5, 7, 9}
	o := odd
	fmt.Println(o)

}
