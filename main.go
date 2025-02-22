package main

import (
	"fmt"
	"github.com/matoval/airtHive/backend/llama"
	"github.com/matoval/airtHive/backend/server"
)


func main() {
    fmt.Println("Calling gtk_gui application from Go...")
    err := server.StartServer("localhost:50051", &llama.LLM{})
    if err != nil {
        panic(err)
    }
}

