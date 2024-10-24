package main

import "fmt"

func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

func main() {
	a := 1
	b := 2
	fmt.Printf("a до применения функции = %d\n", a)
	fmt.Printf("b до применения функции = %d\n", b)
	Swap(&a, &b)
	fmt.Printf("a после применения функции = %d\n", a)
	fmt.Printf("b после применения функции = %d\n", b)
}
