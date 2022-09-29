package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", PlayerServer)
	http.ListenAndServe(":8080", nil)
	fmt.Println("starting the server on port 8080")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello", r.URL.Path)
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}

func GETPlayers() string {
	return "20"
}
