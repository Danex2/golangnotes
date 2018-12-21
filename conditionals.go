package main

import "fmt"

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
