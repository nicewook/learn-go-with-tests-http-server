package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	log.Printf("r.URL.Path: %v, player name: %v", r.URL.Path, player)

	fmt.Fprint(w, GetPlayerScore(player))
}

func GetPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}
	if player == "Floyd" {
		return "10"
	}
	return ""
}

func main() {
	handler := http.HandlerFunc(PlayerServer) // type http.HandleFunc has the method of ServeHTTP()
	if err := http.ListenAndServe(":5050", handler); err != nil {
		log.Fatalf("could not listen on port 5050 %v", err)
	}
}
