package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, " Hello World")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the API ENDPOINT")
}

func PostRequestUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user) // lo que espera esto es una interface generica

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func PostRequestMetadata(w http.ResponseWriter, r *http.Request) {
	// cuando se pasa una interfaz admite cualquier tipo de valor
	// Esto es util cuando no sabemos como será la respuesta del json
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata) // lo que espera esto es una interface generica

	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	//fmt.Println(metadata["name"])
	fmt.Fprintf(w, "Payload %v \n", metadata)
}
