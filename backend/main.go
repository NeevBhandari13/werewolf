package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	"werewolf/protos"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	"golang.org/x/exp/rand"
	"google.golang.org/grpc"
)

const (
	port        = "localhost:8080"
	projectID   = "werewolf-51259"
	credentials = "/werewolf-51259-firebase-adminsdk-f27q8-5ec9287167.json"
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
		games: make(map[string]*Game),
		mu:    sync.Mutex{},
		db:    client, // Firestore client
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
	games map[string]*Game  // In-memory store for games
	mu    sync.Mutex        // Mutex to protect concurrent access to the games map
	db    *firestore.Client // Firestore client
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

func (s *Server) generateGameId() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for {
		code := make([]byte, 6)
		rand.Seed(uint64(time.Now().UnixNano()))
		for i := range code {
			code[i] = charset[rand.Intn(len(charset))]
		}
		gameId := string(code)

		// Check if the game ID already exists
		s.mu.Lock()
		_, exists := s.games[gameId]
		s.mu.Unlock()

		if !exists {
			return gameId
		}
	}
}

func (s *Server) HostGame(ctx context.Context, player *protos.PlayerInfo) (*protos.GameInfo, error) {
	if err := validatePlayerInfo(player); err != nil {
		return nil, err
	}

	gameId := s.generateGameId()

	newGame := &Game{
		ID:      gameId,
		Host:    player,
		Players: []*protos.PlayerInfo{player}, // Add the host as the first player
		State:   protos.State_WAITING,
	}

	// Save game data to Firestore
	// save to collection games, with the name gameId and the contents of newGame
	_, err := s.db.Collection("games").Doc(gameId).Set(ctx, newGame)
	if err != nil {
		log.Printf("Failed to save game to Firestore: %v", err)
		return nil, fmt.Errorf("failed to save game")
	}

	fmt.Printf("Game created with ID: %v and saved to Firestore\n", gameId)

	// Create the response object and send to client
	response := &protos.GameInfo{
		GameId: gameId,
		Players: []*protos.PlayerInfo{
			player,
		},
		State: protos.State_WAITING,
	}
	return response, nil
}

func (s *Server) JoinGame(ctx context.Context, req *protos.JoinRequest) (*protos.GameInfo, error) {
	// Validate player info
	if err := validatePlayerInfo(req.Player); err != nil {
		return nil, err
	}

	// Lock for safe access to games map
	s.mu.Lock()
	defer s.mu.Unlock()

	// Check if the game exists
	game, exists := s.games[req.GameId]
	if !exists {
		return nil, fmt.Errorf("game with ID %s not found", req.GameId)
	}
	// Check if game is in waiting state
	if game.State != protos.State_WAITING {
		return nil, fmt.Errorf("game with ID %s not currently joinable", req.GameId)
	}
	// Check if player is already in the game
	for _, p := range game.Players {
		if p.PlayerId == req.Player.PlayerId {
			return nil, fmt.Errorf("player %s is already in the game", req.Player.PlayerId)
		}
	}

	// Add player to the game
	game.Players = append(game.Players, req.Player)

	// Return updated game info
	response := &protos.GameInfo{
		GameId:  game.ID,
		Players: game.Players,
		State:   game.State,
	}
	return response, nil

}
