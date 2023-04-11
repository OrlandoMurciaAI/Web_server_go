package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		//fmt.Println(servidor, "no esta disponible")
		canal <- servidor + " No se encuentra disponible "
	} else {
		//fmt.Println(servidor, "Esta funcionando normalmente")
		canal <- servidor + " Esta funcionando normalmente"
	}
}

func main() {
	inicio := time.Now()
	canal := make(chan string)
	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	i := 0

	for {
		if i > 10 {
			break
		} else {
			for _, servidor := range servidores {
				go revisarServidor(servidor, canal)
			}
			time.Sleep(1 * time.Second)
			fmt.Println(<-canal)
			i++
		}
	}

	//for i := 0; i < len(servidores); i++ {
	//	fmt.Println(<-canal)
	//}

	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecución: ", tiempoPaso)
}
