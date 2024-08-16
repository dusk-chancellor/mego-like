package services

import (
	"log"

	"github.com/dusk-chancellor/mego-like/internal/models"
)

const element = "like_service"

func (s *likeService) Exists(like models.Like) bool {
	return s.likeLC.Exists(like)
}

func (s *likeService) Like(like models.Like) (string, string, error) {
	userId, postId, err := s.likeRepo.Like(like)
	if err != nil {
		log.Printf("Element: %s | Failed to like: %v", element, err)
		return "", "", err
	}

	go s.likeLC.Like(like)

	return userId, postId, nil
}

func (s *likeService) Count(postId string) int32 {
	return s.likeLC.Count(postId)
}
