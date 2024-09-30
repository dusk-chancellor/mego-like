package services

import (
	"context"
	"log"

	"github.com/dusk-chancellor/mego-like/internal/models"
	"github.com/dusk-chancellor/mego-like/pkg/utils"
)

const element = "like_service"

func (s *likeService) Exists(ctx context.Context, userId, postId, commentId int64) (bool, error) {
	if postId != 0 {
		exists, err := s.likeRepo.PostExists(ctx, userId, postId)
		if err != nil {
			log.Printf("Element: %s | Failed to check if like exists: %v", element, err)
			return false, err
		}

		return exists, nil
	}
	if commentId != 0 {
		exists, err := s.likeRepo.CommentExists(ctx, userId, commentId)
		if err != nil {
			log.Printf("Element: %s | Failed to check if like exists: %v", element, err)
			return false, err
		}

		return exists, nil
	}
	return false, nil
}

func (s *likeService) AddLike(ctx context.Context, userId, postId, commentId int64) error {
	if postId != 0 {
		return s.likeRepo.PostAddLike(ctx, userId, postId)
	}
	if commentId != 0 {
		return s.likeRepo.CommentAddLike(ctx, userId, commentId)
	}

	return nil
}
func (s *likeService) DeleteLike(ctx context.Context, userId, postId, commentId int64) error {
	if postId != 0 {
		return s.likeRepo.PostDeleteLike(ctx, userId, postId)
	}
	if commentId != 0 {
		return s.likeRepo.CommentDeleteLike(ctx, userId, commentId)
	}

	return nil
}

func (s *likeService) Count(ctx context.Context, postId, commentId int64) (int32, error) {
	if postId != 0 {
		return s.likeRepo.PostCount(ctx, postId)
	}
	if commentId != 0 {
		return s.likeRepo.CommentCount(ctx, commentId)
	}
	return 0, nil
}

func (s *likeService) FindByPosts(ctx context.Context, pageSize int, pageToken string) (likes []models.Like, token string, err error) {
	if pageSize < 1 {
		pageSize = 10
	}
	startIndex := 0
	if pageToken != "" {
		startIndex, err = utils.DecodePageToken(pageToken)
		if err != nil {
			log.Printf("Element: %s | Failed to decode page token: %v", element, err)
			return likes, token, err
		}
	}

	likes, err = s.likeRepo.PostFind(ctx, startIndex, pageSize+1)
	if err != nil {
		log.Printf("Element: %s | Failed to find likes: %v", element, err)
		return likes, token, err
	}

	if len(likes) > pageSize {
		token = utils.EncodePageToken(startIndex + pageSize)
		likes = likes[:pageSize]
	}

	return likes, token, err
}
func (s *likeService) FindByComments(ctx context.Context, pageSize int, pageToken string) (likes []models.Like, token string, err error) {
	if pageSize < 1 {
		pageSize = 10
	}
	startIndex := 0
	if pageToken != "" {
		startIndex, err = utils.DecodePageToken(pageToken)
		if err != nil {
			log.Printf("Element: %s | Failed to decode page token: %v", element, err)
			return likes, token, err
		}
	}

	likes, err = s.likeRepo.CommentFind(ctx, startIndex, pageSize+1)
	if err != nil {
		log.Printf("Element: %s | Failed to find likes: %v", element, err)
		return likes, token, err
	}

	if len(likes) > pageSize {
		token = utils.EncodePageToken(startIndex + pageSize)
		likes = likes[:pageSize]
	}

	return likes, token, err
}
