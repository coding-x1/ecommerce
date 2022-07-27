package main

import (
	"fmt"
	"net/http"
	"simpleapp/controller"
)

func main() {
	controller.Master()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	fmt.Println("Server serving...")
	http.ListenAndServe("localhost:3000", nil)

}
