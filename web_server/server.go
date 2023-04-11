// package main indica que el main va leer todos los archivos
package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

// Permitira instanciar nuestro servidor y a escuchar las demas funciones
// el servidor es capaz de ser modificado, por eso el return es un ampersan
// Se va directamente al valor que queremos y no la copia
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]

	if !exist { //SI no existe vamos a crearlo
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	s.router.rules[path][method] = handler
}

func (s *Server) AddMidleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)              // Nuestro punto de entrada
	err := http.ListenAndServe(s.port, nil) // El servidor se pone a escuchar
	if err != nil {
		return err
	} else {
		return nil
	}
}
