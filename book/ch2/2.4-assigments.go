package main

import "fmt"

func main() {
	fmt.Println("")

	x := 1 // named variable
	fmt.Println(x)
	var p (bool)
	fmt.Print(p)
	//	*p = true // indirect variable

	//person.name = "bob" // struct field
	//	count[x] = count[x] * scale // array or slice or map element
}

type person struct {
	name string
}
