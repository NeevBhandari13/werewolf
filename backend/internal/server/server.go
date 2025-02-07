package server

import (
	"context"
	"fmt"
	"log"
	"werewolf/internal/game"
	"werewolf/protos"

	"cloud.google.com/go/firestore"
)

type GameServer struct {
	protos.UnimplementedGameServer
	db *firestore.Client
}

func NewGameServer(db *firestore.Client) *GameServer {
	return &GameServer{db: db}
}

func convertGameToProtoResponse(game *game.Game) (*protos.GameInfo, error) {
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

func (s *GameServer) HostGame(ctx context.Context, player *protos.PlayerInfo) (*protos.GameInfo, error) {

	newGame, err := game.CreateGame(ctx, s.db, player)
	if err != nil {
		return nil, err
	}

	// Create the game into a response object and send to client
	response, err := convertGameToProtoResponse(newGame)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *GameServer) JoinGame(ctx context.Context, req *protos.JoinRequest) (*protos.GameInfo, error) {
	// Validate player info
	err := game.ValidatePlayerInfo(req.Player)
	if err != nil {
		return nil, err
	}

	// create reference to game firestore document for games
	// reference created outside the transaction
	gameRef := s.db.Collection("games").Doc(req.GameId)

	err = s.db.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		// get game's document snapshot from firestore
		gameDoc, err := tx.Get(gameRef)
		// Check if the game exists
		if err != nil {
			log.Printf("Error retrieving game: %v", err)
			return fmt.Errorf("error getting game with ID: %s", req.GameId)
		}

		// decode game document into game struct
		var game game.Game
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

	var updatedGame game.Game
	if err := latestGameDoc.DataTo(&updatedGame); err != nil {
		return nil, fmt.Errorf("error decoding updated game data: %s", err)
	}
	updatedGameResponse, err := convertGameToProtoResponse(&updatedGame)
	if err != nil {
		return nil, err
	}
	return updatedGameResponse, nil

}
