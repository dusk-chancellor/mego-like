package services

import (
	"github.com/dusk-chancellor/mego-like/internal/models"
	"github.com/dusk-chancellor/mego-like/internal/repositories"
)

type LikeService interface {
	Exists(like models.Like) (bool, error)
	Like(like models.Like) (string, string, error)
	Count(postId string) (int32, error)
	Find(pageSize int, pageToken string) ([]*models.Like, string, error)
}

type likeService struct {
	likeRepo repositories.LikeRepository
}

func NewLikeService(likeRepo repositories.LikeRepository) LikeService {
	return &likeService{
		likeRepo: likeRepo,
	}
}
