// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.20.3
// source: protos/game.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// enum for different roles
type Role int32

const (
	Role_NONE         Role = 0
	Role_WEREWOLF     Role = 1
	Role_TROUBLEMAKER Role = 2
	Role_ROBBER       Role = 3
	Role_VILLAGER     Role = 4
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "NONE",
		1: "WEREWOLF",
		2: "TROUBLEMAKER",
		3: "ROBBER",
		4: "VILLAGER",
	}
	Role_value = map[string]int32{
		"NONE":         0,
		"WEREWOLF":     1,
		"TROUBLEMAKER": 2,
		"ROBBER":       3,
		"VILLAGER":     4,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_game_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_protos_game_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_protos_game_proto_rawDescGZIP(), []int{0}
}

// Game states
type State int32

const (
	State_WAITING     State = 0 // Waiting for players to join
	State_NIGHT_TIME  State = 1 // Role players doing actions
	State_IN_PROGRESS State = 2 // Game is currently being played
	State_VOTING      State = 3 // Players are voting
	State_COMPLETED   State = 4 // Game has ended
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "WAITING",
		1: "NIGHT_TIME",
		2: "IN_PROGRESS",
		3: "VOTING",
		4: "COMPLETED",
	}
	State_value = map[string]int32{
		"WAITING":     0,
		"NIGHT_TIME":  1,
		"IN_PROGRESS": 2,
		"VOTING":      3,
		"COMPLETED":   4,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_game_proto_enumTypes[1].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_protos_game_proto_enumTypes[1]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_protos_game_proto_rawDescGZIP(), []int{1}
}

// info about player
type PlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId     string `protobuf:"bytes,1,opt,name=playerId,proto3" json:"playerId,omitempty"`                         // unique ID to identify player
	PlayerName   string `protobuf:"bytes,2,opt,name=playerName,proto3" json:"playerName,omitempty"`                     // player's display name
	StartingRole Role   `protobuf:"varint,3,opt,name=startingRole,proto3,enum=game.Role" json:"startingRole,omitempty"` // player's role at start of round
	EndingRole   Role   `protobuf:"varint,4,opt,name=endingRole,proto3,enum=game.Role" json:"endingRole,omitempty"`     // player's role at end of round
	Voted        bool   `protobuf:"varint,5,opt,name=voted,proto3" json:"voted,omitempty"`                              // boolean indicating whether player has voted
}

func (x *PlayerInfo) Reset() {
	*x = PlayerInfo{}
	mi := &file_protos_game_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerInfo) ProtoMessage() {}

func (x *PlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_game_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerInfo.ProtoReflect.Descriptor instead.
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return file_protos_game_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerInfo) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *PlayerInfo) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *PlayerInfo) GetStartingRole() Role {
	if x != nil {
		return x.StartingRole
	}
	return Role_NONE
}

func (x *PlayerInfo) GetEndingRole() Role {
	if x != nil {
		return x.EndingRole
	}
	return Role_NONE
}

func (x *PlayerInfo) GetVoted() bool {
	if x != nil {
		return x.Voted
	}
	return false
}

// game information
type GameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId  string        `protobuf:"bytes,1,opt,name=gameId,proto3" json:"gameId,omitempty"`                // Unique identifier for the game
	Players []*PlayerInfo `protobuf:"bytes,2,rep,name=players,proto3" json:"players,omitempty"`              // List of players in the game
	State   State         `protobuf:"varint,3,opt,name=state,proto3,enum=game.State" json:"state,omitempty"` // Current state of the game
}

func (x *GameInfo) Reset() {
	*x = GameInfo{}
	mi := &file_protos_game_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameInfo) ProtoMessage() {}

func (x *GameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_protos_game_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameInfo.ProtoReflect.Descriptor instead.
func (*GameInfo) Descriptor() ([]byte, []int) {
	return file_protos_game_proto_rawDescGZIP(), []int{1}
}

func (x *GameInfo) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *GameInfo) GetPlayers() []*PlayerInfo {
	if x != nil {
		return x.Players
	}
	return nil
}

func (x *GameInfo) GetState() State {
	if x != nil {
		return x.State
	}
	return State_WAITING
}

// request to join game
type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameId string      `protobuf:"bytes,1,opt,name=gameId,proto3" json:"gameId,omitempty"` // The game ID or code to join
	Player *PlayerInfo `protobuf:"bytes,2,opt,name=player,proto3" json:"player,omitempty"` // Information about the joining player
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	mi := &file_protos_game_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_game_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_protos_game_proto_rawDescGZIP(), []int{2}
}

func (x *JoinRequest) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *JoinRequest) GetPlayer() *PlayerInfo {
	if x != nil {
		return x.Player
	}
	return nil
}

var File_protos_game_proto protoreflect.FileDescriptor

var file_protos_game_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65, 0x22, 0xba, 0x01, 0x0a, 0x0a, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x6f,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x6f, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x76, 0x6f, 0x74, 0x65, 0x64, 0x22, 0x71, 0x0a, 0x08, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x4f, 0x0a, 0x0b, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61, 0x6d, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x12, 0x28, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2a, 0x4a, 0x0a, 0x04, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x57, 0x45, 0x52, 0x45, 0x57, 0x4f, 0x4c, 0x46, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x52,
	0x4f, 0x55, 0x42, 0x4c, 0x45, 0x4d, 0x41, 0x4b, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06,
	0x52, 0x4f, 0x42, 0x42, 0x45, 0x52, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x56, 0x49, 0x4c, 0x4c,
	0x41, 0x47, 0x45, 0x52, 0x10, 0x04, 0x2a, 0x50, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a,
	0x4e, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a,
	0x06, 0x56, 0x4f, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x04, 0x32, 0x63, 0x0a, 0x04, 0x47, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2d,
	0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x11, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x09, 0x5a,
	0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_game_proto_rawDescOnce sync.Once
	file_protos_game_proto_rawDescData = file_protos_game_proto_rawDesc
)

func file_protos_game_proto_rawDescGZIP() []byte {
	file_protos_game_proto_rawDescOnce.Do(func() {
		file_protos_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_game_proto_rawDescData)
	})
	return file_protos_game_proto_rawDescData
}

var file_protos_game_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protos_game_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_game_proto_goTypes = []any{
	(Role)(0),           // 0: game.Role
	(State)(0),          // 1: game.State
	(*PlayerInfo)(nil),  // 2: game.PlayerInfo
	(*GameInfo)(nil),    // 3: game.GameInfo
	(*JoinRequest)(nil), // 4: game.JoinRequest
}
var file_protos_game_proto_depIdxs = []int32{
	0, // 0: game.PlayerInfo.startingRole:type_name -> game.Role
	0, // 1: game.PlayerInfo.endingRole:type_name -> game.Role
	2, // 2: game.GameInfo.players:type_name -> game.PlayerInfo
	1, // 3: game.GameInfo.state:type_name -> game.State
	2, // 4: game.JoinRequest.player:type_name -> game.PlayerInfo
	2, // 5: game.Game.HostGame:input_type -> game.PlayerInfo
	4, // 6: game.Game.JoinGame:input_type -> game.JoinRequest
	3, // 7: game.Game.HostGame:output_type -> game.GameInfo
	3, // 8: game.Game.JoinGame:output_type -> game.GameInfo
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_game_proto_init() }
func file_protos_game_proto_init() {
	if File_protos_game_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_game_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_game_proto_goTypes,
		DependencyIndexes: file_protos_game_proto_depIdxs,
		EnumInfos:         file_protos_game_proto_enumTypes,
		MessageInfos:      file_protos_game_proto_msgTypes,
	}.Build()
	File_protos_game_proto = out.File
	file_protos_game_proto_rawDesc = nil
	file_protos_game_proto_goTypes = nil
	file_protos_game_proto_depIdxs = nil
}
