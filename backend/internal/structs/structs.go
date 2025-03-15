package structs

import "werewolf/protos"

// A single game session
type Game struct {
	ID      string
	Host    *protos.PlayerInfo
	Players []*protos.PlayerInfo
	State   protos.State // Game state: WAITING, INPROGRESS, etc.
}
