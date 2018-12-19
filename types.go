package main

import "fmt"

/*


// These are all the basic types for Go


bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128




*/

func main() {
	// Type inference happens when no type is added to the definition, Go will add a type based on the value assigned
	i := 42                // Will be treated as an int
	f := 3.142             // Will be treated as a float64
	t := "My name is Dane" // Will be treated as a string
	fmt.Println("Types")
}

// A type can be converted by wrapping the value with the type

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
