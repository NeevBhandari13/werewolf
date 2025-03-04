package game

import (
	"testing"
	"werewolf/protos"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type FirestoreMock struct {
	mock.Mock
}

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

func TestGenerateGameId(t *testing.T) {

}
