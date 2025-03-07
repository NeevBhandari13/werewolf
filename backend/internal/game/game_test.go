package game

import (
	"context"
	"testing"
	"werewolf/protos"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestValidatePlayerInfo(t *testing.T) {

	// define a struct to hold all the tests
	tests := []struct {
		name      string             // label for test
		player    *protos.PlayerInfo // player input
		expectErr bool               // whether or not the test should return an error
	}{
		// arrange
		{"Good player", &protos.PlayerInfo{PlayerName: "Bob", PlayerId: "1"}, false},
		{"Nil Player", nil, true},
		{"Missing Player ID", &protos.PlayerInfo{PlayerName: "Bob"}, true},
		{"Missing Player Name", &protos.PlayerInfo{PlayerId: "456"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// act
			err := ValidatePlayerInfo(tt.player)
			// assert
			if tt.expectErr {
				// if i use requires instead of assert, testing will stop at first failure
				assert.NotNil(t, err)
			}
			if !tt.expectErr {
				assert.Nil(t, err)
			}
		})
	}

}

// MockFirestoreClient mocks FirestoreClientInterface
type MockFirestoreClient struct {
	mock.Mock
}

func (m *MockFirestoreClient) Collection(name string) FirestoreCollectionInterface {
	args := m.Called(name)
	return args.Get(0).(FirestoreCollectionInterface)
}

// MockCollection mocks FirestoreCollectionInterface
type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) Doc(id string) FirestoreDocInterface {
	args := m.Called(id)
	return args.Get(0).(FirestoreDocInterface)
}

// MockDocument mocks FirestoreDocInterface
type MockDocument struct {
	mock.Mock
}

func (m *MockDocument) Get(ctx context.Context) (*firestore.DocumentSnapshot, error) {
	args := m.Called(ctx)
	return nil, args.Error(1)
}

func (m *MockDocument) Set(ctx context.Context, game *Game) (*firestore.DocumentSnapshot, error) {
	args := m.Called(ctx)
	return nil, args.Error(1)
}

func TestGenerateGameId(t *testing.T) {
	ctx := context.Background()

	mockDB := new(MockFirestoreClient)
	mockCollection := new(MockCollection)
	mockDocument := new(MockDocument)

	mockDB.On("Collection", "games").Return(mockCollection)
	mockCollection.On("Doc", mock.Anything).Return(mockDocument)

	// Simulate the first call: Game ID exists (i.e., Get returns a valid document) so we have to retry
	mockDocument.On("Get", ctx).Return(nil, nil).Once() // No error means the ID exists

	// Simulate the second call: Game ID does not exist (i.e., Get returns NotFound error)
	mockDocument.On("Get", ctx).Return(nil, status.Error(codes.NotFound, "not found"))

	// Call generateGameId to check if it retries generating a new ID
	gameID, err := generateGameId(ctx, mockDB)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(gameID) != 6 {
		t.Fatalf("expected game ID of length 6, got %v", gameID)
	}

	// Verify that the mock calls were made as expected
	mockDocument.AssertExpectations(t)
}

func TestCreateGame(t *testing.T) {
	ctx := context.Background()

	// Mock dependencies
	mockDB := new(MockFirestoreClient)
	mockCollection := new(MockCollection)
	mockDocument := new(MockDocument)

	// Set up expectations
	mockDB.On("Collection", "games").Return(mockCollection)
	mockCollection.On("Doc", mock.Anything).Return(mockDocument)

	// ensure generation of game runs smoothly
	// Simulate Firestore behavior: Game does not exist
	mockDocument.On("Get", ctx).Return(nil, status.Error(codes.NotFound, "not found"))
	mockDocument.On("Set", ctx, mock.Anything).Return(nil, nil) // can return nil as does not matter for this function

	// Test case: valid player info
	validPlayer := &protos.PlayerInfo{
		PlayerName: "John",
		PlayerId:   "24",
	}

	// Call CreateGame for valid case
	game, err := CreateGame(ctx, mockDB, validPlayer)
	assert.Nil(t, err)
	assert.NotNil(t, game)
	assert.Equal(t, "24", game.Host.PlayerId)

	// Test case: Invalid player info
	invalidPlayer := &protos.PlayerInfo{PlayerName: "Bob"}

	// Call CreateGame for invalid case
	_, err = CreateGame(ctx, mockDB, invalidPlayer)
	assert.NotNil(t, err) // Should return an error because PlayerId is missing
}
