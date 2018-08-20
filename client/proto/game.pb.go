// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game.proto

package elfgame

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import status "google.golang.org/genproto/googleapis/rpc/status"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Player struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_d600b10a232d86c5, []int{0}
}
func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (dst *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(dst, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Coordinate struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinate) Reset()         { *m = Coordinate{} }
func (m *Coordinate) String() string { return proto.CompactTextString(m) }
func (*Coordinate) ProtoMessage()    {}
func (*Coordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_d600b10a232d86c5, []int{1}
}
func (m *Coordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate.Unmarshal(m, b)
}
func (m *Coordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate.Marshal(b, m, deterministic)
}
func (dst *Coordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate.Merge(dst, src)
}
func (m *Coordinate) XXX_Size() int {
	return xxx_messageInfo_Coordinate.Size(m)
}
func (m *Coordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate proto.InternalMessageInfo

func (m *Coordinate) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Coordinate) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

type Step struct {
	Player               *Player     `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
	Coordinate           *Coordinate `protobuf:"bytes,2,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	Move                 string      `protobuf:"bytes,3,opt,name=move,proto3" json:"move,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Step) Reset()         { *m = Step{} }
func (m *Step) String() string { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()    {}
func (*Step) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_d600b10a232d86c5, []int{2}
}
func (m *Step) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Step.Unmarshal(m, b)
}
func (m *Step) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Step.Marshal(b, m, deterministic)
}
func (dst *Step) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Step.Merge(dst, src)
}
func (m *Step) XXX_Size() int {
	return xxx_messageInfo_Step.Size(m)
}
func (m *Step) XXX_DiscardUnknown() {
	xxx_messageInfo_Step.DiscardUnknown(m)
}

var xxx_messageInfo_Step proto.InternalMessageInfo

func (m *Step) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

func (m *Step) GetCoordinate() *Coordinate {
	if m != nil {
		return m.Coordinate
	}
	return nil
}

func (m *Step) GetMove() string {
	if m != nil {
		return m.Move
	}
	return ""
}

type Reply struct {
	Coordinate           *Coordinate    `protobuf:"bytes,1,opt,name=coordinate,proto3" json:"coordinate,omitempty"`
	NextPlayer           string         `protobuf:"bytes,2,opt,name=next_player,json=nextPlayer,proto3" json:"next_player,omitempty"`
	FinalScore           float32        `protobuf:"fixed32,3,opt,name=final_score,json=finalScore,proto3" json:"final_score,omitempty"`
	Status               *status.Status `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Resigned             bool           `protobuf:"varint,5,opt,name=resigned,proto3" json:"resigned,omitempty"`
	LastMove             string         `protobuf:"bytes,6,opt,name=last_move,json=lastMove,proto3" json:"last_move,omitempty"`
	Board                string         `protobuf:"bytes,7,opt,name=board,proto3" json:"board,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Reply) Reset()         { *m = Reply{} }
func (m *Reply) String() string { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()    {}
func (*Reply) Descriptor() ([]byte, []int) {
	return fileDescriptor_game_d600b10a232d86c5, []int{3}
}
func (m *Reply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reply.Unmarshal(m, b)
}
func (m *Reply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reply.Marshal(b, m, deterministic)
}
func (dst *Reply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reply.Merge(dst, src)
}
func (m *Reply) XXX_Size() int {
	return xxx_messageInfo_Reply.Size(m)
}
func (m *Reply) XXX_DiscardUnknown() {
	xxx_messageInfo_Reply.DiscardUnknown(m)
}

var xxx_messageInfo_Reply proto.InternalMessageInfo

func (m *Reply) GetCoordinate() *Coordinate {
	if m != nil {
		return m.Coordinate
	}
	return nil
}

func (m *Reply) GetNextPlayer() string {
	if m != nil {
		return m.NextPlayer
	}
	return ""
}

func (m *Reply) GetFinalScore() float32 {
	if m != nil {
		return m.FinalScore
	}
	return 0
}

func (m *Reply) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Reply) GetResigned() bool {
	if m != nil {
		return m.Resigned
	}
	return false
}

func (m *Reply) GetLastMove() string {
	if m != nil {
		return m.LastMove
	}
	return ""
}

func (m *Reply) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func init() {
	proto.RegisterType((*Player)(nil), "pytorch.elf.game.v1.Player")
	proto.RegisterType((*Coordinate)(nil), "pytorch.elf.game.v1.Coordinate")
	proto.RegisterType((*Step)(nil), "pytorch.elf.game.v1.Step")
	proto.RegisterType((*Reply)(nil), "pytorch.elf.game.v1.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameClient is the client API for Game service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameClient interface {
	NewGC(ctx context.Context, in *Player, opts ...grpc.CallOption) (*status.Status, error)
	FreeGC(ctx context.Context, in *Player, opts ...grpc.CallOption) (*status.Status, error)
	ClearBoard(ctx context.Context, in *Player, opts ...grpc.CallOption) (*status.Status, error)
	Play(ctx context.Context, in *Step, opts ...grpc.CallOption) (*Reply, error)
	GenMove(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Reply, error)
	Pass(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Reply, error)
	Resign(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Reply, error)
}

type gameClient struct {
	cc *grpc.ClientConn
}

func NewGameClient(cc *grpc.ClientConn) GameClient {
	return &gameClient{cc}
}

func (c *gameClient) NewGC(ctx context.Context, in *Player, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/pytorch.elf.game.v1.Game/NewGC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) FreeGC(ctx context.Context, in *Player, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/pytorch.elf.game.v1.Game/FreeGC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) ClearBoard(ctx context.Context, in *Player, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/pytorch.elf.game.v1.Game/ClearBoard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Play(ctx context.Context, in *Step, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pytorch.elf.game.v1.Game/Play", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) GenMove(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pytorch.elf.game.v1.Game/GenMove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Pass(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pytorch.elf.game.v1.Game/Pass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameClient) Resign(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/pytorch.elf.game.v1.Game/Resign", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServer is the server API for Game service.
type GameServer interface {
	NewGC(context.Context, *Player) (*status.Status, error)
	FreeGC(context.Context, *Player) (*status.Status, error)
	ClearBoard(context.Context, *Player) (*status.Status, error)
	Play(context.Context, *Step) (*Reply, error)
	GenMove(context.Context, *Player) (*Reply, error)
	Pass(context.Context, *Player) (*Reply, error)
	Resign(context.Context, *Player) (*Reply, error)
}

func RegisterGameServer(s *grpc.Server, srv GameServer) {
	s.RegisterService(&_Game_serviceDesc, srv)
}

func _Game_NewGC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).NewGC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pytorch.elf.game.v1.Game/NewGC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).NewGC(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_FreeGC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).FreeGC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pytorch.elf.game.v1.Game/FreeGC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).FreeGC(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_ClearBoard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).ClearBoard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pytorch.elf.game.v1.Game/ClearBoard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).ClearBoard(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Step)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pytorch.elf.game.v1.Game/Play",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Play(ctx, req.(*Step))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_GenMove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).GenMove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pytorch.elf.game.v1.Game/GenMove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).GenMove(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Pass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Pass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pytorch.elf.game.v1.Game/Pass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Pass(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Game_Resign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServer).Resign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pytorch.elf.game.v1.Game/Resign",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServer).Resign(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

var _Game_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pytorch.elf.game.v1.Game",
	HandlerType: (*GameServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewGC",
			Handler:    _Game_NewGC_Handler,
		},
		{
			MethodName: "FreeGC",
			Handler:    _Game_FreeGC_Handler,
		},
		{
			MethodName: "ClearBoard",
			Handler:    _Game_ClearBoard_Handler,
		},
		{
			MethodName: "Play",
			Handler:    _Game_Play_Handler,
		},
		{
			MethodName: "GenMove",
			Handler:    _Game_GenMove_Handler,
		},
		{
			MethodName: "Pass",
			Handler:    _Game_Pass_Handler,
		},
		{
			MethodName: "Resign",
			Handler:    _Game_Resign_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "game.proto",
}

func init() { proto.RegisterFile("game.proto", fileDescriptor_game_d600b10a232d86c5) }

var fileDescriptor_game_d600b10a232d86c5 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xbd, 0x8e, 0xd3, 0x40,
	0x10, 0x96, 0x7d, 0xb6, 0x93, 0x4c, 0x10, 0xc5, 0x80, 0x84, 0xf1, 0x15, 0x39, 0xa5, 0x8a, 0x28,
	0x7c, 0x22, 0xd7, 0x40, 0x81, 0x4e, 0x24, 0x12, 0xa9, 0x40, 0xa7, 0x4d, 0x47, 0x13, 0xed, 0xd9,
	0x93, 0x60, 0xc9, 0xf1, 0x5a, 0xeb, 0x25, 0xc4, 0x0f, 0x40, 0xcf, 0x2b, 0xf1, 0x66, 0x68, 0x67,
	0xa3, 0x80, 0x90, 0x15, 0xe9, 0xdc, 0x65, 0xbe, 0x99, 0xef, 0x67, 0x67, 0x62, 0x80, 0x9d, 0xdc,
	0x53, 0x5a, 0x6b, 0x65, 0x14, 0xbe, 0xa8, 0x5b, 0xa3, 0x74, 0xf6, 0x2d, 0xa5, 0x72, 0x9b, 0x32,
	0x7e, 0x78, 0x9b, 0xbc, 0xda, 0x29, 0xb5, 0x2b, 0xe9, 0x56, 0xd7, 0xd9, 0x6d, 0x63, 0xa4, 0xf9,
	0xde, 0xb8, 0xe9, 0x69, 0x0c, 0xd1, 0x43, 0x29, 0x5b, 0xd2, 0xf8, 0x1c, 0xfc, 0x22, 0x8f, 0xbd,
	0x1b, 0x6f, 0x36, 0x12, 0x7e, 0x91, 0x4f, 0x67, 0x00, 0x4b, 0xa5, 0x74, 0x5e, 0x54, 0xd2, 0x10,
	0x3e, 0x03, 0xef, 0xc8, 0xcd, 0x50, 0x78, 0x47, 0x5b, 0xb5, 0xb1, 0xef, 0xaa, 0x76, 0xfa, 0xcb,
	0x83, 0x60, 0x6d, 0xa8, 0xc6, 0x3b, 0x88, 0x6a, 0x16, 0xe3, 0xc9, 0xf1, 0xfc, 0x3a, 0xed, 0xc8,
	0x92, 0x3a, 0x3f, 0x71, 0x1a, 0xc5, 0x7b, 0x80, 0xec, 0xec, 0xc3, 0xa2, 0xe3, 0xf9, 0xa4, 0x93,
	0xf8, 0x37, 0x8e, 0xf8, 0x87, 0x82, 0x08, 0xc1, 0x5e, 0x1d, 0x28, 0xbe, 0xe2, 0xe8, 0xfc, 0x7b,
	0xfa, 0xd3, 0x87, 0x50, 0x50, 0x5d, 0xb6, 0xff, 0xc9, 0x7b, 0x4f, 0x97, 0x9f, 0xc0, 0xb8, 0xa2,
	0xa3, 0xd9, 0x9c, 0x5e, 0xe6, 0xb3, 0x0b, 0x58, 0xe8, 0xb4, 0xb8, 0x09, 0x8c, 0xb7, 0x45, 0x25,
	0xcb, 0x4d, 0x93, 0x29, 0xed, 0x62, 0xf8, 0x02, 0x18, 0x5a, 0x5b, 0x04, 0xdf, 0x40, 0xe4, 0x76,
	0x1e, 0x07, 0x6c, 0x8f, 0xa9, 0xbb, 0x46, 0xaa, 0xeb, 0x2c, 0x5d, 0x73, 0x47, 0x9c, 0x26, 0x30,
	0x81, 0xa1, 0xa6, 0xa6, 0xd8, 0x55, 0x94, 0xc7, 0xe1, 0x8d, 0x37, 0x1b, 0x8a, 0x73, 0x8d, 0xd7,
	0x30, 0x2a, 0x65, 0x63, 0x36, 0xfc, 0xda, 0x88, 0x73, 0x0c, 0x2d, 0xf0, 0x59, 0x1d, 0x08, 0x5f,
	0x42, 0xf8, 0xa8, 0xa4, 0xce, 0xe3, 0x01, 0x37, 0x5c, 0x31, 0xff, 0x7d, 0x05, 0xc1, 0x4a, 0xee,
	0x09, 0xdf, 0x41, 0xf8, 0x85, 0x7e, 0xac, 0x96, 0x78, 0xe9, 0x26, 0x49, 0x47, 0x32, 0x7c, 0x0f,
	0xd1, 0x27, 0x4d, 0xd4, 0x87, 0xfa, 0x01, 0x60, 0x59, 0x92, 0xd4, 0x0b, 0x9b, 0xa5, 0x0f, 0x3d,
	0xb0, 0x5d, 0x7c, 0xdd, 0x49, 0xb4, 0xff, 0xb8, 0x24, 0xe9, 0x6c, 0xb9, 0xcb, 0x2f, 0x60, 0xb0,
	0xa2, 0x8a, 0x97, 0x73, 0xd1, 0xfa, 0x92, 0xc6, 0x3d, 0x04, 0x0f, 0xb2, 0x69, 0xfa, 0x0b, 0x7c,
	0x84, 0x48, 0xf0, 0xfd, 0x7a, 0x4b, 0x2c, 0x46, 0x5f, 0x07, 0x54, 0x6e, 0x2d, 0xf6, 0x18, 0xf1,
	0x47, 0x7b, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xc2, 0x4b, 0xa1, 0xf0, 0x03, 0x00, 0x00,
}
