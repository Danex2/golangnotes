package main

import "fmt"

/*
- Slices are dynamically sized arrays
- Slices are made by defining a high and low bound ([stat index:end index])
- the numbers before the start and after the end index are not incluced in the slice
- It does not store any data, it is just a reference to the original array
- Anything changed in the slice will also be changed in the array it references
- Slice literals are like arrays but without a fixed length
*/

func main() {
	myArr := [5]int{1, 2, 3, 4, 5}
	// Take the first two elements and put them in firstTwo
	firstTwo := myArr[0:2]
	fmt.Println(firstTwo)
	// Change the value of the first element in the firstTwo array to 2
	firstTwo[0] = 2
	// The change also occurs in the original array
	fmt.Println(myArr)

	// Slice literal
	lit := []bool{true, true, false}
	for i, val := range lit {
		fmt.Println(i, val)
	}
}
