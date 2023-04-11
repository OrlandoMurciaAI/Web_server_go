package main

import "fmt"

type animal interface {
	moverse() string
}

type pajaro struct{}
type pez struct{}
type perro struct{}

func (perro) moverse() string {
	return "soy un perro y camino"
}

func (pez) moverse() string {
	return "soy un pez y nado"
}

func (pajaro) moverse() string {
	return "soy un pajaro y vuelo"
}

func moverAnimal(a animal) {
	fmt.Println(a.moverse())
}

func main() {
	p := perro{}
	moverAnimal(p)
	pe := pez{}
	moverAnimal(pe)
	pa := pajaro{}
	moverAnimal(pa)
}
