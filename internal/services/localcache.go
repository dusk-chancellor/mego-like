package services

import (
	"sync"

	"github.com/dusk-chancellor/mego-like/internal/models"
)

type likeLocalCache struct {
	likes map[string]string
	sync.Mutex
}

func NewLikeLocalCache() *likeLocalCache {
	return &likeLocalCache{
		likes: make(map[string]string),
	}
}

func (l *likeLocalCache) Like(like models.Like) {
	l.Lock()
	defer l.Unlock()

	l.likes[like.PostId] = like.UserId
}

func (l *likeLocalCache) Exists(like models.Like) bool {
	_, ok := l.likes[like.PostId]

	return ok
}

func (l *likeLocalCache) Count(postId string) int32 {
	count := len(l.likes)

	return int32(count)
}
