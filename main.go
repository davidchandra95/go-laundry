package main

import (
	"flag"
	"fmt"
	"github.com/davidchandra95/go-laundry/config"
	server "github.com/davidchandra95/go-laundry/routes"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.InitConfig(*environment)
	config.InitDB()
	var db = config.GetDB()
	server.InitServer()

	fmt.Println("server is running", db)
}
