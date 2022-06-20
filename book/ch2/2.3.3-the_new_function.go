package main

import "fmt"

func main() {
	p := new(int)   // p, of type *int, points to an unnamed int variable
	fmt.Println(*p) // "0"
	fmt.Println(p)  // address of p
	*p = 2          // sets the unnamed int to 2
	fmt.Println(*p)
	fmt.Println(p)

	// distinct values
	a := new(int)
	b := new(int)
	fmt.Println(a == b)

}

func newInt() *int {

	return new(int)
}

func newInteger() *int {
	var dummy int
	return &dummy
}

func delta(old, new int) int {
	return new - old
}
