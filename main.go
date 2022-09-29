package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/players/", PlayerServer)
	http.ListenAndServe(":8080", nil)
	fmt.Println("starting the server on port 8080")
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	name := strings.Split(r.URL.Path, "/")[2]
	score := GETPlayers(name)
	fmt.Fprint(w, score)
}

func GETPlayers(name string) int {
	if name == "Pepper" {
		return 20
	} else {
		return 10
	}
}
