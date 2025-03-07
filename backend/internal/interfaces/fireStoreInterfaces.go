package interfaces

import (
	"context"

	"cloud.google.com/go/firestore"
	"werewolf/internal/structs"
)

// FirestoreClientInterface allows us to mock Firestore
type FirestoreClientInterface interface {
	Collection(name string) FirestoreCollectionInterface
}

// FirestoreCollectionInterface for mocking Firestore collections
type FirestoreCollectionInterface interface {
	Doc(id string) FirestoreDocInterface
}

// FirestoreDocInterface for mocking Firestore documents
type FirestoreDocInterface interface {
	Get(ctx context.Context) (*firestore.DocumentSnapshot, error)
	Set(ctx context.Context, game *structs.Game) (*firestore.DocumentSnapshot, error)
}
