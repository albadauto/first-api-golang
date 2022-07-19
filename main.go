package main

import "github.com/albadauto/projeto/server"

func main() {
	server := server.NewServer()

	server.Run()
}