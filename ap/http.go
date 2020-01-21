package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("SERVER_NAME")
	fmt.Fprintf(w, "%s got request.", name)
	fmt.Printf("%s got request.\n", name)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
