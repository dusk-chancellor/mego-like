package services

import (
	"context"
	"github.com/dusk-chancellor/mego-like/internal/clients"
	"github.com/dusk-chancellor/mego-like/internal/models"
	"github.com/dusk-chancellor/mego-like/internal/repositories"
)

type LikeService interface {
	Exists(ctx context.Context, userId, postId, commentId int64) (bool, error)
	AddLike(ctx context.Context, userId, postId, commentId int64) error
	DeleteLike(ctx context.Context, userId, postId, commentId int64) error
	Count(ctx context.Context, postId, commentId int64) (int32, error)
	FindByPosts(ctx context.Context, pageSize int, pageToken string) ([]models.Like, string, error)
	FindByComments(ctx context.Context, pageSize int, pageToken string) ([]models.Like, string, error)
}

type likeService struct {
	likeRepo      repositories.LikeRepository
	postClient    *clients.PostClient
	commentClient *clients.CommentClient
}

func NewLikeService(likeRepo repositories.LikeRepository, postClient *clients.PostClient, commentClient *clients.CommentClient) LikeService {
	return &likeService{
		likeRepo:      likeRepo,
		postClient:    postClient,
		commentClient: commentClient,
	}
}
