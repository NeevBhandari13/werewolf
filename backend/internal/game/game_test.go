package game

import (
	"testing"
	"werewolf/protos"
)

func TestValidatePlayerInfo(t *testing.T) {
	// define a struct to hold all the tests
	tests := []struct {
		name      string             // label for test
		player    *protos.PlayerInfo // player input
		expectErr bool               // whether or not the test should return an error
	}{
		{"Good player", &protos.PlayerInfo{PlayerName: "Bob", PlayerId: "1"}, false},
		{"Nil Player", nil, true},
		{"Missing Player ID", &protos.PlayerInfo{PlayerName: "Bob"}, true},
		{"Missing Player Name", &protos.PlayerInfo{PlayerId: "456"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePlayerInfo(tt.player)
			if tt.expectErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}

}
