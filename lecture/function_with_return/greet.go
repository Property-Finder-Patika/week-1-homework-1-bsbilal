package main

import (
	"fmt"
)

func main() {
	var name string = "Lalibb"
	var greeting = createGreet(name)
	fmt.Printf("%s", greeting)
}
func createGreet(name string) string {
	greeting := "Selam " + name + " :)"
	return greeting
}
