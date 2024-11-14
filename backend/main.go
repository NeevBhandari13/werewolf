package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var client *firestore.Client

func main() {
	// Initialize Firebase
	ctx := context.Background()
	opt := option.WithCredentialsFile("firebaseAdminConfig.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing Firebase app: %v", err)
	}

	// Initialize Firestore client
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error initializing Firestore: %v", err)
	}
	defer client.Close()

	// Set up a simple HTTP server
	http.HandleFunc("/test", testHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	_, _, err := client.Collection("testCollection").Add(ctx, map[string]interface{}{
		"message": "Hello, Firebase!",
	})
	if err != nil {
		http.Error(w, "Failed to write to Firestore", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Data added to Firestore!")
}
