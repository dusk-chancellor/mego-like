package grpc

import (
	"context"
	"log"

	pb "github.com/antibomberman/mego-protos/gen/go/like"
	"github.com/dusk-chancellor/mego-like/internal/dto"
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
	exists, err := s.service.Exists(like)
	if err != nil {
		log.Printf("Element: %s | Failed to check if like exists: %v", element, err)
		return nil, err
	}

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

	count, err := s.service.Count(postId)
	if err != nil {
		log.Printf("Element: %s | Failed to count: %v", element, err)
		return nil, err
	}

	return &pb.CountResponse{Count: count}, nil
}

func (s *serverAPI) Find(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	pageSize := int(req.GetPageSize())
	pageToken := req.GetPageToken()
	likes, nextPageToken, err := s.service.Find(pageSize, pageToken)
	if err != nil {
		log.Printf("Element: %s | Failed to find: %v", element, err)
		return nil, err
	}
	pbLikes := dto.ToPbLikes(likes)	

	return &pb.FindResponse{
		Likes: pbLikes,
		NextPageToken: nextPageToken,
	}, nil
}
