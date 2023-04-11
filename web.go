package main

import (
	"fmt"
	"io"
	"net/http"
)

// Se tienen que crear las interfaces para que los otros metodos funcionen
type escritorWeb struct{}

// Handlers y middlewares

//Middlewares sirve para autenticacioens

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	respuesta, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	e := escritorWeb{}
	io.Copy(e, respuesta.Body)

}
