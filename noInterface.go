package main

import "fmt"

type pajaro struct{}
type pez struct{}
type perro struct{}

func (p *perro) caminar() string {
	return "soy un perro y camino"
}

func (p *pez) nada() string {
	return "soy un pez y nado"
}

func (p *pajaro) vuela() string {
	return "soy un pajaro y vuelo"
}

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func moverPez(p pez) {
	fmt.Println(p.nada())
}

func moverPajaro(p pajaro) {
	fmt.Println(p.vuela())
}

func main() {
	p := perro{}

	pe := pez{}

	pa := pajaro{}

	moverPerro(p)
	moverPez(pe)
	moverPajaro(pa)
}
