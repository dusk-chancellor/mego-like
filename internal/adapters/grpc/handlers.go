package grpc

import (
	"context"
	"log"

	pb "github.com/antibomberman/mego-protos/gen/go/like"
	"github.com/dusk-chancellor/mego-like/internal/dto"
)

const element = "like_handlers"

func (s *serverAPI) Exists(ctx context.Context, req *pb.ExistsRequest) (*pb.ExistsResponse, error) {
	exists, err := s.service.Exists(ctx, req.GetUserId(), req.GetPostId(), req.GetCommentId())
	if err != nil {
		log.Printf("Element: %s | Failed to check if like exists: %v", element, err)
		return nil, err
	}

	return &pb.ExistsResponse{Exists: exists}, nil
}

func (s *serverAPI) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {

	err := s.service.AddLike(ctx, req.GetUserId(), req.GetPostId(), req.GetCommentId())
	if err != nil {
		log.Printf("Element: %s | Failed to like: %v", element, err)
		return nil, err
	}

	return &pb.AddResponse{UserId: req.GetUserId(), PostId: req.GetPostId(), CommentId: req.GetCommentId()}, nil
}
func (s *serverAPI) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {

	err := s.service.DeleteLike(ctx, req.GetUserId(), req.GetPostId(), req.GetCommentId())
	if err != nil {
		log.Printf("Element: %s | Failed to like: %v", element, err)
		return nil, err
	}

	return &pb.DeleteResponse{UserId: req.GetUserId(), PostId: req.GetPostId(), CommentId: req.GetCommentId()}, nil
}

func (s *serverAPI) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountResponse, error) {

	count, err := s.service.Count(ctx, req.GetPostId(), req.GetCommentId())
	if err != nil {
		log.Printf("Element: %s | Failed to count: %v", element, err)
		return nil, err
	}

	return &pb.CountResponse{Count: count}, nil
}

func (s *serverAPI) FindByPosts(ctx context.Context, req *pb.FindByPostsRequest) (*pb.FindByPostsResponse, error) {
	likes, nextPageToken, err := s.service.FindByPosts(ctx, int(req.GetPageSize()), req.GetPageToken())
	if err != nil {
		log.Printf("Element: %s | Failed to find: %v", element, err)
		return nil, err
	}
	pbLikes := dto.ToPbLikes(likes)

	return &pb.FindByPostsResponse{
		Likes:         pbLikes,
		NextPageToken: nextPageToken,
	}, nil
}

func (s *serverAPI) FindByComments(ctx context.Context, req *pb.FindByCommentsRequest) (*pb.FindByCommentsResponse, error) {
	likes, nextPageToken, err := s.service.FindByPosts(ctx, int(req.GetPageSize()), req.GetPageToken())
	if err != nil {
		log.Printf("Element: %s | Failed to find: %v", element, err)
		return nil, err
	}
	pbLikes := dto.ToPbLikes(likes)

	return &pb.FindByCommentsResponse{
		Likes:         pbLikes,
		NextPageToken: nextPageToken,
	}, nil
}
