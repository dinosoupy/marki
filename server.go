package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Marki is attempting to go online...")
	http.Handle("/", http.FileServer(http.Dir("frontend/dist")))
	http.ListenAndServe(":8080", nil)
}