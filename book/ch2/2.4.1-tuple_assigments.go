package main

func main() {

	x := 1
	y := 1
	x, y = y, x
	var a [2]int
	a[0] = 1
	a[1] = 2
	j := 1
	i := 0
	a[i], a[j] = a[j], a[i]

}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
