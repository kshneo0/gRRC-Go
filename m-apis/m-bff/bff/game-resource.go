package bff

import (
	pbgameengine "github.com/gRPC/m-apis/m-game-engine/v1"
	pbhighscore "github.com/gRPC/m-apis/m-highscore/v1"
)

type gameResource struct {
	gameClient       pbhighscore.GameClient
	gameEngineClient pbgameengine.GameEngineClient
}

func NewGameResource(gameClient pbhighscore.GameClient, gameEngineClient pbgameengine.GameEngineClient) *gameResource {
	return &gameResource{
		gameClient:       gameClient,
		gameEngineClient: gameEngineClient,
	}
}
