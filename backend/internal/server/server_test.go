package server_test

import (
	"context"
	"testing"
	"werewolf/protos"

	"cloud.google.com/go/firestore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"fmt"
	"log"
	"werewolf/internal/game"

	"cloud.google.com/go/firestore"
)

func TestCreateGame(t *testing.T) {
	ctx = context.Background()

	// Mock dependencies
	mockDB := new(MockFirestoreClient)
	mockCollection := new(MockCollection)
	mockDocument := new(MockDocument)
}
