package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	"werewolf/protos"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
)

const (
	port        = "localhost:8080"
	projectID   = "werewolf-51259"
	credentials = "werewolf-51259-firebase-adminsdk-f27q8-5ec9287167.json"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx := context.Background() // create context
	// create firestore client
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credentials))
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	grpcServer := grpc.NewServer()
	gameServer := &Server{
		db: client, // Firestore client
	}

	protos.RegisterGameServer(grpcServer, gameServer)

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// A single game session
type Game struct {
	ID      string
	Host    *protos.PlayerInfo
	Players []*protos.PlayerInfo
	State   protos.State // Game state: WAITING, IN_PROGRESS, etc.
}

// the gRPC server
type Server struct {
	protos.UnimplementedGameServer
	db *firestore.Client // Firestore client
}

func validatePlayerInfo(player *protos.PlayerInfo) error {
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

func (s *Server) generateGameId(ctx context.Context) (string, error) {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for {
		code := make([]byte, 6)
		rand.Seed(uint64(time.Now().UnixNano()))
		for i := range code {
			code[i] = charset[rand.Intn(len(charset))]
		}
		gameId := string(code)

		// Check Firestore if this gameId already exists
		gameRef := s.db.Collection("games").Doc(gameId)
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

func convertGameToGameInfo(game *Game) (*protos.GameInfo, error) {
	if game == nil {
		return nil, fmt.Errorf("game does not exist")
	}

	// Ensure GameId is valid
	if game.ID == "" {
		return nil, fmt.Errorf("missing game ID")
	}

	// Ensure Players is never nil
	if game.Players == nil {
		game.Players = []*protos.PlayerInfo{}
		return nil, fmt.Errorf("game has no host")
	}

	protoResponse := &protos.GameInfo{
		GameId:  game.ID,
		Players: game.Players,
		State:   game.State,
	}

	return protoResponse, nil
}

func (s *Server) HostGame(ctx context.Context, player *protos.PlayerInfo) (*protos.GameInfo, error) {
	if err := validatePlayerInfo(player); err != nil {
		return nil, err
	}

	gameId, err := s.generateGameId(ctx)
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
	_, err = s.db.Collection("games").Doc(gameId).Set(ctx, newGame)
	if err != nil {
		log.Printf("Failed to save game to Firestore: %v", err)
		return nil, fmt.Errorf("failed to save game")
	}

	fmt.Printf("Game created with ID: %v and saved to Firestore\n", gameId)

	// Create the game into a response object and send to client
	response, err := convertGameToGameInfo(newGame)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *Server) JoinGame(ctx context.Context, req *protos.JoinRequest) (*protos.GameInfo, error) {
	// Validate player info
	if err := validatePlayerInfo(req.Player); err != nil {
		return nil, err
	}

	// create reference to game firestore document for games
	// reference created outside the transaction
	gameRef := s.db.Collection("games").Doc(req.GameId)

	err := s.db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		// get game's document snapshot from firestore
		gameDoc, err := tx.Get(gameRef)
		// Check if the game exists
		if err != nil {
			log.Printf("Error retrieving game: %v", err)
			return fmt.Errorf("error getting game with ID: %s", req.GameId)
		}

		// decode game document into game struct
		var game Game
		err = gameDoc.DataTo(&game) // function decodes firestore json into game struct
		if err != nil {
			log.Printf("Error decoding Firestore game data: %v", err)
			return fmt.Errorf("error decoding game with ID: %s", req.GameId)
		}

		log.Printf("Game successfully retrieved from Firestore and decoded: %+v", game)

		// Check if game is in waiting state
		if game.State != protos.State_WAITING {
			return fmt.Errorf("game with ID %s not currently joinable", req.GameId)
		}
		// Check if player is already in the game
		for _, p := range game.Players {
			if p.PlayerId == req.Player.PlayerId {
				return fmt.Errorf("player %s is already in the game", req.Player.PlayerId)
			}
		}

		// Add player to the game
		game.Players = append(game.Players, req.Player)

		// set new players list in firestore
		err = tx.Set(gameRef, game)
		if err != nil {
			log.Printf("Error updating Firestore with new player: %v", err)
			return fmt.Errorf("failed to update game")
		}

		return nil

	})

	// verify transaction succeeded
	if err != nil {
		return nil, err
	}

	// fetch latest game doc to return to client
	// get game doc outside transaction
	latestGameDoc, err := gameRef.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting latest game state after transaction: %s", err)
	}

	var updatedGame Game
	if err := latestGameDoc.DataTo(&updatedGame); err != nil {
		return nil, fmt.Errorf("error decoding updated game data: %s", err)
	}
	updatedGameResponse, err := convertGameToGameInfo(&updatedGame)
	if err != nil {
		return nil, err
	}
	return updatedGameResponse, nil

}
