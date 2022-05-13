package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 12
	fmt.Println(a, ",", b)
	a, b = b, a
	fmt.Println(a, ",", b)

	// OR

	x := 100
	y := 200
	fmt.Println(x, ",", y)
	fmt.Println(a, ",", b)
	x = x + y
	y = x - y
	x = x - y
	fmt.Println(x, ",", y)

}

//https://go.dev/play/p/UONOQw11Dvh
