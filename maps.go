package main

import "fmt"

// un map es una estructura llave valor
func main() {
	m1 := make(map[string]int)
	m1["orlando"] = 8
	fmt.Printf("%v \n", m1["orlando"])
}
