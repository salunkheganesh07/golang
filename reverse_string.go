package main

import "fmt"

func main() {
	s := "ganesh"
	r := []rune(s)
	var reverse []rune
	for ii := len(r) - 1; ii >= 0; ii-- {
		reverse = append(reverse, r[ii])
	}
	fmt.Println(string(reverse))
	//https://go.dev/play/p/glL70nMTGtl
}
