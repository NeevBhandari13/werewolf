package game

import (
	"context"
	"fmt"
	"log"
	"time"
	"werewolf/protos"

	"cloud.google.com/go/firestore"
	"golang.org/x/exp/rand"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// A single game session
type Game struct {
	ID      string
	Host    *protos.PlayerInfo
	Players []*protos.PlayerInfo
	State   protos.State // Game state: WAITING, IN_PROGRESS, etc.
}

func ValidatePlayerInfo(player *protos.PlayerInfo) error {
	if player == nil {
		return fmt.Errorf("player info cannot be nil")
	}
	if player.PlayerId == "" {
		return fmt.Errorf("player ID is required")
	}
	if player.PlayerName == "" {
		return fmt.Errorf("player name is required")
	}
	return nil
}

func generateGameId(ctx context.Context, db *firestore.Client) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for {
		code := make([]byte, 6)
		rand.Seed(uint64(time.Now().UnixNano()))
		for i := range code {
			code[i] = charset[rand.Intn(len(charset))]
		}
		gameId := string(code)

		// Check Firestore if this gameId already exists
		gameRef := db.Collection("games").Doc(gameId)
		_, err := gameRef.Get(ctx)

		// If the document does not exist, the ID is unique and can be used
		if err != nil && status.Code(err) == codes.NotFound {
			return gameId, nil
		}

		// If Firestore returns an error other than "NotFound", return the error
		if err != nil {
			return "", fmt.Errorf("failed to check game ID existence: %v", err)
		}
	}
}

func CreateGame(ctx context.Context, db *firestore.Client, player *protos.PlayerInfo) (*Game, error) {
	if err := ValidatePlayerInfo(player); err != nil {
		return nil, err
	}

	gameId, err := generateGameId(ctx, db)
	if err != nil {
		return nil, err
	}

	// game data for database
	newGame := &Game{
		ID:      gameId,
		Host:    player,
		Players: []*protos.PlayerInfo{player}, // Add the host as the first player
		State:   protos.State_WAITING,
	}

	// Save game data to Firestore
	// save to collection games, with the name gameId and the contents of newGame
	_, err = db.Collection("games").Doc(gameId).Set(ctx, newGame)
	if err != nil {
		log.Printf("Failed to save game to Firestore: %v", err)
		return nil, fmt.Errorf("failed to save game")
	}

	fmt.Printf("Game created with ID: %v and saved to Firestore\n", gameId)

	return newGame, nil
}
