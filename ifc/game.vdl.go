// This file was auto-generated by the vanadium vdl tool.
// Source: game.vdl

package ifc

import (
	// VDL system imports
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/vdl"
)

type Player struct {
	Id int32
}

func (Player) __VDLReflect(struct {
	Name string `vdl:"github.com/monopole/volley/ifc.Player"`
}) {
}

type MasterCommand struct {
	Name string
}

func (MasterCommand) __VDLReflect(struct {
	Name string `vdl:"github.com/monopole/volley/ifc.MasterCommand"`
}) {
}

type Ball struct {
	Owner Player
	X     float32
	Y     float32
	Dx    float32
	Dy    float32
}

func (Ball) __VDLReflect(struct {
	Name string `vdl:"github.com/monopole/volley/ifc.Ball"`
}) {
}

func init() {
	vdl.Register((*Player)(nil))
	vdl.Register((*MasterCommand)(nil))
	vdl.Register((*Ball)(nil))
}

// GameServiceClientMethods is the client interface
// containing GameService methods.
type GameServiceClientMethods interface {
	// Receiver adds the player p to list of known players and
	// concomitantly promises to inform p of game state changes.
	Recognize(ctx *context.T, p Player, opts ...rpc.CallOpt) error
	// Receiver forgets player p, because player p has quit
	// or has been ejected from the game.
	Forget(ctx *context.T, p Player, opts ...rpc.CallOpt) error
	// Accept a ball.
	Accept(ctx *context.T, b Ball, opts ...rpc.CallOpt) error
	// Quit
	Quit(*context.T, ...rpc.CallOpt) error
	// Master command
	DoMasterCommand(ctx *context.T, c MasterCommand, opts ...rpc.CallOpt) error
	// Change value of pause duration.
	SetPauseDuration(ctx *context.T, p float32, opts ...rpc.CallOpt) error
	// Change value of gravity
	SetGravity(ctx *context.T, p float32, opts ...rpc.CallOpt) error
}

// GameServiceClientStub adds universal methods to GameServiceClientMethods.
type GameServiceClientStub interface {
	GameServiceClientMethods
	rpc.UniversalServiceMethods
}

// GameServiceClient returns a client stub for GameService.
func GameServiceClient(name string) GameServiceClientStub {
	return implGameServiceClientStub{name}
}

type implGameServiceClientStub struct {
	name string
}

func (c implGameServiceClientStub) Recognize(ctx *context.T, i0 Player, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Recognize", []interface{}{i0}, nil, opts...)
	return
}

func (c implGameServiceClientStub) Forget(ctx *context.T, i0 Player, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Forget", []interface{}{i0}, nil, opts...)
	return
}

func (c implGameServiceClientStub) Accept(ctx *context.T, i0 Ball, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Accept", []interface{}{i0}, nil, opts...)
	return
}

func (c implGameServiceClientStub) Quit(ctx *context.T, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Quit", nil, nil, opts...)
	return
}

func (c implGameServiceClientStub) DoMasterCommand(ctx *context.T, i0 MasterCommand, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "DoMasterCommand", []interface{}{i0}, nil, opts...)
	return
}

func (c implGameServiceClientStub) SetPauseDuration(ctx *context.T, i0 float32, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "SetPauseDuration", []interface{}{i0}, nil, opts...)
	return
}

func (c implGameServiceClientStub) SetGravity(ctx *context.T, i0 float32, opts ...rpc.CallOpt) (err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "SetGravity", []interface{}{i0}, nil, opts...)
	return
}

// GameServiceServerMethods is the interface a server writer
// implements for GameService.
type GameServiceServerMethods interface {
	// Receiver adds the player p to list of known players and
	// concomitantly promises to inform p of game state changes.
	Recognize(ctx *context.T, call rpc.ServerCall, p Player) error
	// Receiver forgets player p, because player p has quit
	// or has been ejected from the game.
	Forget(ctx *context.T, call rpc.ServerCall, p Player) error
	// Accept a ball.
	Accept(ctx *context.T, call rpc.ServerCall, b Ball) error
	// Quit
	Quit(*context.T, rpc.ServerCall) error
	// Master command
	DoMasterCommand(ctx *context.T, call rpc.ServerCall, c MasterCommand) error
	// Change value of pause duration.
	SetPauseDuration(ctx *context.T, call rpc.ServerCall, p float32) error
	// Change value of gravity
	SetGravity(ctx *context.T, call rpc.ServerCall, p float32) error
}

// GameServiceServerStubMethods is the server interface containing
// GameService methods, as expected by rpc.Server.
// There is no difference between this interface and GameServiceServerMethods
// since there are no streaming methods.
type GameServiceServerStubMethods GameServiceServerMethods

// GameServiceServerStub adds universal methods to GameServiceServerStubMethods.
type GameServiceServerStub interface {
	GameServiceServerStubMethods
	// Describe the GameService interfaces.
	Describe__() []rpc.InterfaceDesc
}

// GameServiceServer returns a server stub for GameService.
// It converts an implementation of GameServiceServerMethods into
// an object that may be used by rpc.Server.
func GameServiceServer(impl GameServiceServerMethods) GameServiceServerStub {
	stub := implGameServiceServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implGameServiceServerStub struct {
	impl GameServiceServerMethods
	gs   *rpc.GlobState
}

func (s implGameServiceServerStub) Recognize(ctx *context.T, call rpc.ServerCall, i0 Player) error {
	return s.impl.Recognize(ctx, call, i0)
}

func (s implGameServiceServerStub) Forget(ctx *context.T, call rpc.ServerCall, i0 Player) error {
	return s.impl.Forget(ctx, call, i0)
}

func (s implGameServiceServerStub) Accept(ctx *context.T, call rpc.ServerCall, i0 Ball) error {
	return s.impl.Accept(ctx, call, i0)
}

func (s implGameServiceServerStub) Quit(ctx *context.T, call rpc.ServerCall) error {
	return s.impl.Quit(ctx, call)
}

func (s implGameServiceServerStub) DoMasterCommand(ctx *context.T, call rpc.ServerCall, i0 MasterCommand) error {
	return s.impl.DoMasterCommand(ctx, call, i0)
}

func (s implGameServiceServerStub) SetPauseDuration(ctx *context.T, call rpc.ServerCall, i0 float32) error {
	return s.impl.SetPauseDuration(ctx, call, i0)
}

func (s implGameServiceServerStub) SetGravity(ctx *context.T, call rpc.ServerCall, i0 float32) error {
	return s.impl.SetGravity(ctx, call, i0)
}

func (s implGameServiceServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implGameServiceServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{GameServiceDesc}
}

// GameServiceDesc describes the GameService interface.
var GameServiceDesc rpc.InterfaceDesc = descGameService

// descGameService hides the desc to keep godoc clean.
var descGameService = rpc.InterfaceDesc{
	Name:    "GameService",
	PkgPath: "github.com/monopole/volley/ifc",
	Methods: []rpc.MethodDesc{
		{
			Name: "Recognize",
			Doc:  "// Receiver adds the player p to list of known players and\n// concomitantly promises to inform p of game state changes.",
			InArgs: []rpc.ArgDesc{
				{"p", ``}, // Player
			},
		},
		{
			Name: "Forget",
			Doc:  "// Receiver forgets player p, because player p has quit\n// or has been ejected from the game.",
			InArgs: []rpc.ArgDesc{
				{"p", ``}, // Player
			},
		},
		{
			Name: "Accept",
			Doc:  "// Accept a ball.",
			InArgs: []rpc.ArgDesc{
				{"b", ``}, // Ball
			},
		},
		{
			Name: "Quit",
			Doc:  "// Quit",
		},
		{
			Name: "DoMasterCommand",
			Doc:  "// Master command",
			InArgs: []rpc.ArgDesc{
				{"c", ``}, // MasterCommand
			},
		},
		{
			Name: "SetPauseDuration",
			Doc:  "// Change value of pause duration.",
			InArgs: []rpc.ArgDesc{
				{"p", ``}, // float32
			},
		},
		{
			Name: "SetGravity",
			Doc:  "// Change value of gravity",
			InArgs: []rpc.ArgDesc{
				{"p", ``}, // float32
			},
		},
	},
}
