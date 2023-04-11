package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

// Receiver function cuando se instancia que la funcion recibe un struct
// calc.operate
func (calc) operate(entrada string, operador string) (salida int) {
	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		salida = operador1 + operador2
	case "-":
		salida = operador1 - operador2
	case "*":
		salida = operador1 * operador2
	case "/":
		salida = operador1 / operador2
	default:
		salida = 0
		fmt.Println(operador, " No esta soportado")
	}
	fmt.Println("El resultado de tu operacion es: ", salida)
	return
}

func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

func leerEntrada() (texto string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texto = scanner.Text()
	return
}

func main() {
	entrada := leerEntrada()
	operador := leerEntrada()
	c := calc{}
	fmt.Println(c.operate(entrada, operador))
}
