package dto

import (
	"github.com/antibomberman/mego-protos/gen/go/like"
	"github.com/dusk-chancellor/mego-like/internal/models"
)

func ToPbLikes(model []*models.Like) (pbLikes []*like.Like) {
	for _, modelLike := range model {
		pbLike := &like.Like{
			UserId: modelLike.UserId,
			PostId: modelLike.PostId,
		}
		pbLikes = append(pbLikes, pbLike)
	}
	return
}
