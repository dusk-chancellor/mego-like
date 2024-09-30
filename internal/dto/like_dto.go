package dto

import (
	pb "github.com/antibomberman/mego-protos/gen/go/like"
	"github.com/dusk-chancellor/mego-like/internal/models"
)

func ToPbLikes(model []models.Like) []*pb.Like {
	pbLikes := make([]*pb.Like, 0, len(model))

	for _, modelLike := range model {
		pbLikes = append(pbLikes, &pb.Like{
			UserId:    modelLike.UserId,
			PostId:    modelLike.PostId.Int64,
			CommentId: modelLike.CommentId.Int64,
		})
	}
	return pbLikes
}
