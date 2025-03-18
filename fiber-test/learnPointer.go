package main

import "fmt"

func changeValue(n *int) {
	*n = 3
}

func main() {
	x := 2
	y := &x
	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(x)
	changeValue(y)
	fmt.Println(x)
}
