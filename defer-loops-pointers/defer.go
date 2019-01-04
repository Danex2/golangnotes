package main

import "fmt"

/*
* Read more about defer
- defer stalls the function execution until the surrounding function
has returned
- deferred calls are pushed onto a stack and returned in a last in
first out order
*/
func main() {
	deferLoop()
}

func deferFunc() {
	fmt.Println("Defer function")
}

func normalFunc() {
	fmt.Println("Normal function call")
}

func deferLoop() {
	// The start text is printed
	fmt.Println("Start of loop")

	for i := 0; i < 15; i++ {
		/*
			we get to this line but every number is stalled and instead
			put on a stack waiting for the surrounding function to return
		*/
		defer fmt.Println(i)
	}
	// The end text is printed right after the start text
	fmt.Println("End of loop")
	/*
		we get to the end of the function and there is nothing else to do
		so the deferred print statements are printed in reverse order because
		defer works as last in first out
	*/

	// output: 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1
}
