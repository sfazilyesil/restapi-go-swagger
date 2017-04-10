package main

import (
	"github.com/sfazilyesil/restapi-go-swagger/server"
)

//go:generate swagger generate spec
func main() {
	s := server.NewServer()
	s.Run()
}
