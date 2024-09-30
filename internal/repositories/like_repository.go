package repositories

import (
	"context"
	"fmt"
	"github.com/dusk-chancellor/mego-like/internal/models"
)

const element = "like_repository"

func (r *likeRepository) PostExists(ctx context.Context, userId, postId int64) (bool, error) {
	query := `
        SELECT COUNT(*) 
        FROM likes 
        WHERE post_id = $1 AND user_id = $2
    `

	var count int
	err := r.db.QueryRowContext(ctx, query, postId, userId).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check post existence: %w", err)
	}

	return count > 0, nil
}

func (r *likeRepository) PostAddLike(ctx context.Context, userId, postId int64) error {
	query := `INSERT INTO likes (user_id, post_id) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, userId, postId)
	if err != nil {
		return fmt.Errorf("failed to like post: %w", err)
	}

	return nil
}

func (r *likeRepository) PostDeleteLike(ctx context.Context, userId, postId int64) error {
	query := `DELETE FROM likes WHERE user_id = $1 AND post_id = $2`

	_, err := r.db.ExecContext(ctx, query, userId, postId)
	if err != nil {
		return fmt.Errorf("failed to unlike post: %w", err)
	}

	return nil
}

func (r *likeRepository) PostFind(ctx context.Context, offset, limit int) (likes []models.Like, err error) {
	query := `
        SELECT *
        FROM likes
        WHERE post_id IS NOT NULL
        ORDER BY id DESC
        OFFSET $1 LIMIT $2
    `
	err = r.db.SelectContext(ctx, &likes, query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query likes: %w", err)
	}

	return likes, nil
}

func (r *likeRepository) PostCount(ctx context.Context, postId int64) (count int32, err error) {
	query := `
        SELECT COUNT(*) 
        FROM likes 
        WHERE post_id = $1
    `
	err = r.db.QueryRowContext(ctx, query, postId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count likes: %w", err)
	}

	return count, nil
}

func (r *likeRepository) CommentExists(ctx context.Context, userId, commentId int64) (bool, error) {
	query := `
        SELECT COUNT(*) 
        FROM likes 
        WHERE comment_id = $1 AND user_id = $2
    `

	var count int
	err := r.db.QueryRowContext(ctx, query, commentId, userId).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check comment existence: %w", err)
	}

	return count > 0, nil
}

func (r *likeRepository) CommentAddLike(ctx context.Context, userId, commentId int64) error {
	query := `INSERT INTO likes (user_id, comment_id) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, userId, commentId)
	if err != nil {
		return fmt.Errorf("failed to like comment: %w", err)
	}

	return nil
}

func (r *likeRepository) CommentDeleteLike(ctx context.Context, userId, commentId int64) error {
	query := `DELETE FROM likes WHERE user_id = $1 AND comment_id = $2`

	_, err := r.db.ExecContext(ctx, query, userId, commentId)
	if err != nil {
		return fmt.Errorf("failed to unlike comment: %w", err)
	}

	return nil
}

func (r *likeRepository) CommentFind(ctx context.Context, offset, limit int) (likes []models.Like, err error) {
	query := `
        SELECT *
        FROM likes
        WHERE comment_id IS NOT NULL
        ORDER BY id DESC
        OFFSET $1 LIMIT $2
    `
	err = r.db.SelectContext(ctx, &likes, query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query likes: %w", err)
	}

	return likes, nil
}

func (r *likeRepository) CommentCount(ctx context.Context, commentId int64) (count int32, err error) {
	query := `
        SELECT COUNT(*) 
        FROM likes 
        WHERE comment_id = $1
    `
	err = r.db.QueryRowContext(ctx, query, commentId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count likes: %w", err)
	}

	return count, nil
}
