package clients

import (
	pb "github.com/antibomberman/mego-protos/gen/go/comment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CommentClient struct {
	pb.CommentServiceClient
}

func NewCommentClient(address string) (*CommentClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &CommentClient{pb.NewCommentServiceClient(conn)}, nil
}
