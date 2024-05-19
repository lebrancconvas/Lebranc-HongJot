package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/lebrancconvas/Lebranc-HongJot/db"
	"github.com/lebrancconvas/Lebranc-HongJot/server"
)

func Init() {

}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}

	flag.Parse()

	// Init Database
	db.Init()

	// Start Server
	server.Start()

}
