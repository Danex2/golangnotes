package main

import "fmt"

/*
- A pointer is a memory address of a value
- & is used to point a variable to the memory address
- * is used to read through the pointer back to the original value
- you can also reassign the value of a variable by reading through
its pointer
* Need more on pointers
*/

func main() {
	var d, a = 33, 89
	k := &d
	i := *k
	o := &a
	fmt.Println(k, *o, i)
}
