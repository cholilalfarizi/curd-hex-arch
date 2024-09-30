package main

import (
	"crud-hex/internals/server"
	"log"
)

func main(){
	app := server.Setup()
	log.Fatal(app.Listen(":3000"))
}