package main

import (
	"context"
	"fmt"
	"log"
	"werewolf/protos"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverAddr = "localhost:8080"

func main() {
	// Connect to the gRPC server
	conn, err := grpc.NewClient(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := protos.NewGameClient(conn)

	// Host a new game
	hostPlayer := &protos.PlayerInfo{
		PlayerId:     "1",
		PlayerName:   "Host Player",
		StartingRole: protos.Role_NONE,
	}
	gameResponse, err := client.HostGame(context.Background(), hostPlayer)
	if err != nil {
		log.Fatalf("Error hosting game: %v", err)
	}
	fmt.Printf("Game hosted with ID: %s\n", gameResponse.GameId)

	testGameId := gameResponse.GameId

	// Join the game
	joinPlayer := &protos.PlayerInfo{
		PlayerId:     "456",
		PlayerName:   "Joining Player",
		StartingRole: protos.Role_NONE,
	}

	joinReq := &protos.JoinRequest{
		GameId: testGameId,
		Player: joinPlayer,
	}

	joinResponse, err := client.JoinGame(context.Background(), joinReq)
	if err != nil {
		log.Fatalf("Error joining game: %v", err)
	}
	fmt.Printf("Joined game: %s, Players: %d\n", joinResponse.GameId, len(joinResponse.Players))
}
