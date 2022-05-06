package api

import (
	pb "github.com/velibor7/XML/common/proto/post_service"
	"github.com/velibor7/XML/post_service/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:      post.Id.Hex(),
		Text:    post.Text,
		Created: timestamppb.New(post.Created),
		UserId:  post.UserId,
		Images:  post.Images,
		Links:   post.Links,
	}

	for _, like := range post.Likes {
		postPb.Likes = append(postPb.Likes, &pb.Like{
			Id:     like.Id.Hex(),
			UserId: like.UserId,
		})
	}

	for _, dislike := range post.Dislikes {
		postPb.Dislikes = append(postPb.Dislikes, &pb.Dislike{
			Id:     dislike.Id.Hex(),
			UserId: dislike.UserId,
		})
	}

	for _, comment := range post.Comments {
		postPb.Comments = append(postPb.Comments, &pb.Comment{
			Id:     comment.Id.Hex(),
			UserId: comment.UserId,
			Text:   comment.Text,
		})
	}

	return postPb
}
