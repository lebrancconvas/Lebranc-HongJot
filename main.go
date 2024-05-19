package main

import (
	"github.com/lebrancconvas/Lebranc-HongJot/db"
	"github.com/lebrancconvas/Lebranc-HongJot/server"
)

func main() {
	// Init Database
	db.Init()

	// Start Server
	server.Start()
}
