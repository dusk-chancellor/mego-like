package grpc

import (
	"context"
	"log"

	pb "github.com/antibomberman/mego-protos/gen/go/like"
	"github.com/dusk-chancellor/mego-like/internal/models"
)

const element = "like_handlers"

func (s *serverAPI) Exists(ctx context.Context, req *pb.ExistsRequest) (*pb.ExistsResponse, error) {
	userId := req.GetUserId()
	postId := req.GetPostId()

	like := models.Like{
		UserId: userId,
		PostId: postId,
	}
	exists := s.service.Exists(like)

	return &pb.ExistsResponse{Exists: exists}, nil
}

func (s *serverAPI) Like(ctx context.Context, req *pb.LikeRequest) (*pb.LikeResponse, error) {
	userId := req.GetUserId()
	postId := req.GetPostId()

	like := models.Like{
		UserId: userId,
		PostId: postId,
	}

	userId, postId, err := s.service.Like(like)
	if err != nil {
		log.Printf("Element: %s | Failed to like: %v", element, err)
		return nil, err
	}

	return &pb.LikeResponse{UserId: userId, PostId: postId}, nil
}

func (s *serverAPI) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountResponse, error) {
	postId := req.GetPostId()

	count := s.service.Count(postId)

	return &pb.CountResponse{Count: count}, nil
}
