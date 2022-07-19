package main

import (
	"github.com/albadauto/projeto/database"
	"github.com/albadauto/projeto/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()
	server.Run()
}