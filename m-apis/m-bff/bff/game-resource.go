package bff

import (
	"context"
	"strconv"

	pbgameengine "github.com/gRPC/m-apis/m-game-engine/v1"
	pbhighscore "github.com/gRPC/m-apis/m-highscore/v1"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
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

func NewGrpcGameServiceClient(serverAddr string) (pbhighscore.GameClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal().Msgf("Failed to dial: %v", err)
		return nil, err
	} else {
		log.Info().Msgf("Successfully connected to [%s]", serverAddr)
	}

	if conn == nil {
		log.Info().Msg("m-highscore connection is nil in m-bff")
	}

	client := pbhighscore.NewGameClient(conn)

	return client, nil
}

func NewGrpcGameEngineServiceClient(serverAddr string) (pbgameengine.GameEngineClient, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())

	if err != nil {
		log.Fatal().Msgf("Failed to dial: %v", err)
		return nil, err
	} else {
		log.Info().Msgf("Successfully connected to [%s]", serverAddr)
	}

	if conn == nil {
		log.Info().Msg("m-game-engine connection is nil in m-bff")
	}

	client := pbgameengine.NewGameEngineClient(conn)

	return client, nil
}

func (gr *gameResource) SetHighScore(c *gin.Context) {
	highScoreString := c.Param("hs")
	highScoreFloat64, err := strconv.ParseFloat(highScoreString, 64)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert highscore to float")
	}
	gr.gameClient.SetHighScore(context.Background(), &pbhighscore.SetHighScoreRequest{
		HighScore: highScoreFloat64,
	})
}

func (gr *gameResource) GetHighScore(c *gin.Context) {

	highScoreResponse, err := gr.gameClient.GetHighScore(context.Background(), &pbhighscore.GetHighScoreRequest{})
	if err != nil {
		log.Error().Err(err).Msg("Error while getting highscore")
		return
	}
	hsString := strconv.FormatFloat(highScoreResponse.HighScore, 'e', -1, 64)

	c.JSONP(200, gin.H{
		"hs": hsString,
	})

}
