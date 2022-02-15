package grpc

import (
	"context"

	"github.com/gRPC/m-apis/m-game-engine/internal/server/logic"
	pbgameengine "github.com/gRPC/m-apis/m-game-engine/v1"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) GetSize(ctx context.Context, input *pbgameengine.GetSizeRequest) (*pbgameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in m-game-engine called")
	return &pbgameengine.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}

func (g *Grpc) SetScore(ctx context.Context, input *pbgameengine.SetScoreRequest) (*pbgameengine.SetScoreResponse, error) {
	log.Info().Msg("SetScore in m-game-engine called")
	set := logic.SetScore(input.Score)
	return &pbgameengine.SetScoreResponse{
		Set: set,
	}, nil
}
