package repositories

// redis: key structure -> user_id:post_id

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/dusk-chancellor/mego-like/internal/models"
)

const element = "like_repository"
var ctx = context.Background()

func (r *likeRepository) Exists(like models.Like) (bool, error) {
	key := fmt.Sprintf("%s:%s", like.UserId, like.PostId)
	exists, _ := r.redis.Get(ctx, key).Bool()
	if exists {
		return true, nil
	}

	q := `SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = $1 AND user_id = $2);`

	if err := r.db.QueryRow(q, like.PostId, like.UserId).Scan(&exists); err != nil {
		log.Printf("Element: %s | Failed to check if like exists in db: %v", element, err)
		return false, err
	}

	return exists, nil
}

func (r *likeRepository) Like(like models.Like) (string, string, error) {
	q := `INSERT INTO likes (user_id, post_id) VALUES ($1, $2) RETURNING user_id, post_id;`

	var userId, postId string
	if err := r.db.QueryRow(q, like.UserId, like.PostId).Scan(&userId, &postId); err != nil {
		log.Printf("Element: %s | Failed to like in db: %v", element, err)
		return "", "", err
	}

	key := fmt.Sprintf("%s:%s", like.UserId, like.PostId)
	_, err := r.redis.Set(ctx, key, 1, 24*time.Hour).Result()
	if err != nil {
		log.Printf("Element: %s | Failed to like in redis: %v", element, err)
	}

	return userId, postId, nil
}

func (r *likeRepository) Find(startIndex, pageSize int) ([]*models.Like, error) {
	q := `SELECT * FROM likes LIMIT $1 OFFSET $2;`

	var likes []*models.Like
	if err := r.db.Select(&likes, q, startIndex, pageSize); err != nil {
		log.Printf("Element: %s | Failed to find likes in db: %v", element, err)
		return nil, err
	}
	if len(likes) == 0 {
		return []*models.Like{}, nil
	}
	return likes, nil
}

func (r *likeRepository) Count(postId string) (int32, error) {
/*	exists, _ := r.redis.Get(context.Background(), "*:"+postId).Bool() 
	if exists {
		count, err := r.countRedis(postId)
		if err != nil {
			log.Printf("Element: %s | Failed to count likes in redis: %v", element, err)
		}
		return count, nil
	}
*/
	q := `SELECT COUNT(*) FROM likes WHERE post_id = $1;`

	var count int32
	if err := r.db.QueryRow(q, postId).Scan(&count); err != nil {
		log.Printf("Element: %s | Failed to count likes in db: %v", element, err)
		return 0, err
	}

	return count, nil
}
/*
func (r *likeRepository) countRedis(postId string) (int32, error) {
	var totalCount int64
	var cursor uint64

	for {
		keys, _, err := r.redis.Scan(context.Background(), cursor, "*:"+postId, 0).Result()
		if err != nil {
			return 0, err
		}

		if len(keys) == 0 {
			break
		}

		for _, key := range keys {
			val, err := r.redis.Get(context.Background(), key).Int64()
			if err != nil {
				return 0, err
			}
			totalCount += val
		}

		cursor = uint64(len(keys))
	}
	return int32(totalCount), nil
}
*/