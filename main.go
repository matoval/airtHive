package main

/*
#cgo pkg-config: gtk4
#include "gtk_gui/app.c"
*/
import "C"
import "fmt"

func main() {
    fmt.Println("Calling gtk_gui application from Go...")
    C.start(0, nil)
}

