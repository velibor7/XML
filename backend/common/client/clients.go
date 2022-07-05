package client

import (
	"context"
	"time"

	pbComment "github.com/velibor7/XML/common/proto/comment_service"
	pbProfile "github.com/velibor7/XML/common/proto/profile_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCommentClient(address string) (pbComment.CommentServiceClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbComment.NewCommentServiceClient(conn)
	return client, nil
}

func NewProfileClient(address string) (pbProfile.ProfileServiceClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*30)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, err
	}
	client := pbProfile.NewProfileServiceClient(conn)
	return client, nil
}
