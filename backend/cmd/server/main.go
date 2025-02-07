package main

import (
	"context"
	"log"
	"net"
	"werewolf/internal/server"
	"werewolf/protos"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"

	"google.golang.org/grpc"
)

const (
	port        = "localhost:8080"
	projectID   = "werewolf-51259"
	credentials = "/Users/neevbhandari/Projects/werewolf/backend/werewolf-51259-firebase-adminsdk-f27q8-5ec9287167.json"
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
	gameServer := server.NewGameServer(client)

	protos.RegisterGameServer(grpcServer, gameServer)

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
