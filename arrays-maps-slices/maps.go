package main

import "fmt"

/*
- Maps are a collection of key value pairs
*/

// Map literals are like struct literals but the keys need to be defined

var k = map[string]int{
	"Stacy": 30,
	"Adam":  22,
}

func main() {
	// Make an empty map
	m := make(map[string]int)
	// Add a key value pair to the map
	m["Dane"] = 24
	fmt.Println(m)
	fmt.Println(k)
}
