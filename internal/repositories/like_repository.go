package repositories

import (
	"log"

	"github.com/dusk-chancellor/mego-like/internal/models"
)

const element = "like_repository"


func (r *likeRepository) Exists(like models.Like) (bool, error) {
	q := `SELECT EXISTS(SELECT 1 FROM likes WHERE post_id = $1 AND user_id = $2);`

	var exists bool
	if err := r.db.QueryRow(q, like.PostId, like.UserId).Scan(&exists); err != nil {
		log.Printf("Element: %s | Failed to check if like exists: %v", element, err)
		return false, err
	}

	return exists, nil
}

func (r *likeRepository) Like(like models.Like) (string, string, error) {
	q := `INSERT INTO likes (user_id, post_id) VALUES ($1, $2) RETURNING user_id, post_id;`

	var userId, postId string
	if err := r.db.QueryRow(q, like.UserId, like.PostId).Scan(&userId, &postId); err != nil {
		log.Printf("Element: %s | Failed to like: %v", element, err)
		return "", "", err
	}

	return userId, postId, nil
}

func (r *likeRepository) Count(postId string) (int32, error) {
	q := `SELECT COUNT(*) FROM likes WHERE post_id = $1;`

	var count int32
	if err := r.db.QueryRow(q, postId).Scan(&count); err != nil {
		log.Printf("Element: %s | Failed to count likes: %v", element, err)
		return 0, err
	}

	return count, nil
}
