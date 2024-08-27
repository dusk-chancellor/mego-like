package services

import (
	"log"

	"github.com/dusk-chancellor/mego-like/internal/models"
	"github.com/dusk-chancellor/mego-like/pkg/utils"
)

const element = "like_service"

func (s *likeService) Exists(like models.Like) (bool, error) {
	exists, err := s.likeRepo.Exists(like)
	if err != nil {
		log.Printf("Element: %s | Failed to check if like exists: %v", element, err)
		return false, err
	}

	return exists, nil
}

func (s *likeService) Like(like models.Like) (string, string, error) {
	userId, postId, err := s.likeRepo.Like(like)
	if err != nil {
		log.Printf("Element: %s | Failed to like: %v", element, err)
		return "", "", err
	}

	return userId, postId, nil
}

func (s *likeService) Count(postId string) (int32, error) {
	count, err := s.likeRepo.Count(postId)
	if err != nil {
		log.Printf("Element: %s | Failed to count likes: %v", element, err)
		return 0, err
	}

	return count, nil
}

func (s *likeService) Find(pageSize int, pageToken string) ([]*models.Like, string, error) {
	var err error
	if pageSize < 1 {
		pageSize = 10
	}
	startIndex := 0
	if pageToken != "" {
		startIndex, err = utils.DecodePageToken(pageToken)
		if err != nil {
			log.Printf("Element: %s | Failed to decode page token: %v", element, err)
			return nil, "", err
		}
	}

	likes, err := s.likeRepo.Find(startIndex, pageSize+1)
	if err != nil {
		log.Printf("Element: %s | Failed to find likes: %v", element, err)
		return nil, "", err
	}

	var nextPageToken string
	if len(likes) > pageSize {
		nextPageToken = utils.EncodePageToken(startIndex + pageSize)
		likes = likes[:pageSize]
	}

	return likes, nextPageToken, nil
}
