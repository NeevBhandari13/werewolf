package mocks

import (
	"context"
	"werewolf/internal/interfaces"
	"werewolf/internal/structs"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/mock"
)

// MockFirestoreClient mocks FirestoreClientInterface
type MockFirestoreClient struct {
	mock.Mock
}

func (m *MockFirestoreClient) Collection(name string) interfaces.FirestoreCollectionInterface {
	args := m.Called(name)
	return args.Get(0).(interfaces.FirestoreCollectionInterface)
}

// MockCollection mocks FirestoreCollectionInterface
type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) Doc(id string) interfaces.FirestoreDocInterface {
	args := m.Called(id)
	return args.Get(0).(interfaces.FirestoreDocInterface)
}

// MockDocument mocks FirestoreDocInterface
type MockDocument struct {
	mock.Mock
}

func (m *MockDocument) Get(ctx context.Context) (*firestore.DocumentSnapshot, error) {
	args := m.Called(ctx)
	return nil, args.Error(1)
}

func (m *MockDocument) Set(ctx context.Context, game *structs.Game) (*firestore.DocumentSnapshot, error) {
	args := m.Called(ctx)
	return nil, args.Error(1)
}
