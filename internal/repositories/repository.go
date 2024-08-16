package repositories

import (
	"github.com/dusk-chancellor/mego-like/internal/models"
	"github.com/jmoiron/sqlx"
)

type LikeRepository interface {
	Exists(like models.Like) (bool, error)
	Like(like models.Like) (string, string, error)
	Count(postId string) (int32, error)
}

type likeRepository struct {
	db *sqlx.DB
}

func NewLikeRepository(db *sqlx.DB) LikeRepository {
	return &likeRepository{
		db: db,
	}
}
