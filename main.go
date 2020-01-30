package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// NewInMemoryPlayerStore is InMemoryPlayerStore instance
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

// InMemoryPlayerStore is store the player in memory
type InMemoryPlayerStore struct {
	store map[string]int
}

// GetPlayerScore from memory
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

// RecordWin is record the user then store in memory
func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

// GetLeague is show the list of player with wins
func (i *InMemoryPlayerStore) GetLeague() []Player {
	var league []Player

	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	return league
}

func main() {
	server := NewPlayerServer(NewInMemoryPlayerStore())
	if err := http.ListenAndServe(GetPort(), server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

// GetPort Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "3000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
