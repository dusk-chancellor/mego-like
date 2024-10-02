package grpc

import (
	"github.com/dusk-chancellor/mego-like/internal/config"
	"github.com/dusk-chancellor/mego-like/internal/services"

	"github.com/antibomberman/mego-protos/gen/go/like"
	"google.golang.org/grpc"
)

type serverAPI struct {
	like.UnimplementedLikeServiceServer
	service services.LikeService
	cfg     *config.Config
}

func RegisterGRPC(grpc *grpc.Server, service services.LikeService, cfg *config.Config) {
	like.RegisterLikeServiceServer(grpc, &serverAPI{
		service: service,
		cfg:     cfg,
	})
}
