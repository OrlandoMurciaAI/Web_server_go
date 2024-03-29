package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Name  string `json:"name"` //Definir cuando serializar o deserializar segun el json
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJson() ([]byte, error) { //*User nos permite evitar crear copias
	return json.Marshal(u) // Encodear el struct a un json
}
