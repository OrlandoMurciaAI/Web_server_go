package main

import "fmt"

func cambiarValor(a int) {
	fmt.Println(&a)
	a = 36
}

func main() {
	x := 25
	fmt.Println(x)
	fmt.Println(x)
	cambiarValor(x)
	y := &x
	*y = 305
	fmt.Println(x)
}
