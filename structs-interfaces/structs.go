package main

import "fmt"

/*
- A struct is a collection of fields
- The values inside of a struct can be accessed by using dot notation (similar to a javascript object)
*/

type Thing struct {
	N string
	U int
}

type Person struct {
	Name      string
	shoe_size int
}

func main() {
	o := Thing{"Dane", 24}
	fmt.Println(o)
	fmt.Println(o.N)

	// Assign struct to variable
	j := Person{"Dane", 10}

	// Set pointer to struct
	l := &j

	// Pointer value can be accessed with just dot notation and no asterisk
	fmt.Println(l.shoe_size)

	// Reassignment is also allowed
	l.shoe_size = 15
	fmt.Println(l.shoe_size)
}
