package main

import (
	"fmt"
	"github.com/matoval/airtHive/backend/server"
)


func main() {
    fmt.Println("Calling gtk_gui application from Go...")
    err := server.StartServer()
    if err != nil {
        panic(err)
    }
}

