package repositories

import (
	"github.com/dusk-chancellor/mego-like/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type LikeRepository interface {
	Exists(like models.Like) (bool, error)
	Like(like models.Like) (string, string, error)
	Count(postId string) (int32, error)
}

type likeRepository struct {
	db *sqlx.DB
	redis *redis.Client
}

func NewLikeRepository(db *sqlx.DB, rdb *redis.Client) LikeRepository {
	return &likeRepository{
		db: db,
		redis: rdb,
	}
}
