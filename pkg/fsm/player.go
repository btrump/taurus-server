package fsm

import (
	"log"
)

// Player implements a Player in the game
type Player struct {
	ID   string
	Name string
}

func NewPlayer(id string, n string) *Player {
	log.Printf("fsm::Player(): Creating new player")
	return &Player{
		ID:   id,
		Name: n,
	}
}