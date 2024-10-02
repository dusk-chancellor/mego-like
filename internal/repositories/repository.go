package repositories

import (
	"context"
	"github.com/dusk-chancellor/mego-like/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type LikeRepository interface {
	PostExists(ctx context.Context, userId, postId int64) (bool, error)
	PostAddLike(ctx context.Context, userId, postId int64) error
	PostDeleteLike(ctx context.Context, userId, postId int64) error
	PostCount(ctx context.Context, postId int64) (int32, error)
	PostFind(ctx context.Context, offset, limit int) ([]models.Like, error)

	CommentExists(ctx context.Context, userId, commentId int64) (bool, error)
	CommentAddLike(ctx context.Context, userId, commentId int64) error
	CommentDeleteLike(ctx context.Context, userId, commentId int64) error
	CommentCount(ctx context.Context, commentId int64) (int32, error)
	CommentFind(ctx context.Context, startIndex, pageSize int) ([]models.Like, error)
}

type likeRepository struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewLikeRepository(db *sqlx.DB, rdb *redis.Client) LikeRepository {
	return &likeRepository{
		db:    db,
		redis: rdb,
	}
}
