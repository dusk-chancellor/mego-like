package services

import (
	"log"

	"github.com/dusk-chancellor/mego-like/internal/models"
)

const element = "like_service"

func (s *likeService) Exists(like models.Like) bool {
	exists, err := s.likeRepo.Exists(like)
	if err != nil {
		log.Printf("Element: %s | Failed to check if like exists: %v", element, err)
		return false
	}

	return exists
}

func (s *likeService) Like(like models.Like) (string, string, error) {
	userId, postId, err := s.likeRepo.Like(like)
	if err != nil {
		log.Printf("Element: %s | Failed to like: %v", element, err)
		return "", "", err
	}

	return userId, postId, nil
}

func (s *likeService) Count(postId string) int32 {
	count, err := s.likeRepo.Count(postId)
	if err != nil {
		log.Printf("Element: %s | Failed to count likes: %v", element, err)
		return 0
	}

	return count
}
