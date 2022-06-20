package main

import "fmt"

func main() {
	// pointers
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	// zero pointers
	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil)

	//function pointer addresses
	var px = fx()
	fmt.Println(px == px)     // returun true
	fmt.Println(fx() == fx()) // return false

	fmt.Println("pointer increment")
	// function pointer increment
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)

}
func fx() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	fmt.Println(p)
	return *p
}
