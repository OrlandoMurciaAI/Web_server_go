package main

import "net/http"

func main() {
	server := NewServer(":3000")
	server.Handle(http.MethodGet, "/", HandleRoot)
	server.Handle(http.MethodPost, "/create", PostRequestMetadata)
	server.Handle(http.MethodPost, "/user", PostRequestUser)
	server.Handle(http.MethodPost, "/api", server.AddMidleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}
