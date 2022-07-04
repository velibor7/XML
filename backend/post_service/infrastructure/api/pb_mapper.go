package api

import (
	"fmt"
	"time"

	pb "github.com/velibor7/XML/common/proto/post_service"
	"github.com/velibor7/XML/post_service/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapPost(post *domain.Post) *pb.Post {
	postPb := &pb.Post{
		Id:      post.Id.Hex(),
		Text:    post.Text,
		Created: timestamppb.New(post.Created),
		UserId:  post.UserId.Hex(),
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

	return postPb
}

func mapCreatePost(post *pb.Post) *domain.Post {
	id, _ := primitive.ObjectIDFromHex(post.Id)
	userId, _ := primitive.ObjectIDFromHex(post.UserId)
	postPb := &domain.Post{
		Id:      id,
		Text:    post.Text,
		UserId:  userId,
		Images:  post.Images,
		Links:   post.Links,
		Created: time.Now(),
	}

	return postPb
}

func mapUpdatePost(oldData *pb.Post, newData *pb.Post) *domain.Post {
	id, _ := primitive.ObjectIDFromHex(oldData.Id)
	userId, _ := primitive.ObjectIDFromHex(oldData.UserId)
	fmt.Println(id)

	postPb := &domain.Post{
		Id:      id,
		Text:    newData.Text,
		UserId:  userId,
		Images:  newData.Images,
		Links:   newData.Links,
		Created: oldData.Created.AsTime(),
	}

	for _, like := range newData.Likes {
		if like.Id == "" {
			if likedPostByUser(oldData, like.UserId) == false {
				like_id := primitive.NewObjectID()
				postPb.Likes = append(postPb.Likes, domain.Like{
					Id:     like_id,
					UserId: like.UserId,
				})
			}
		} else {
			like_id, _ := primitive.ObjectIDFromHex(like.Id)
			postPb.Likes = append(postPb.Likes, domain.Like{
				Id:     like_id,
				UserId: like.UserId,
			})
		}
	}

	for _, dislike := range newData.Dislikes {
		if dislike.Id == "" {
			if dislikedPostByUser(oldData, dislike.UserId) == false {
				dislike_id := primitive.NewObjectID()
				postPb.Dislikes = append(postPb.Dislikes, domain.Dislike{
					Id:     dislike_id,
					UserId: dislike.UserId,
				})
			}
		} else {
			dislike_id, _ := primitive.ObjectIDFromHex(dislike.Id)
			postPb.Dislikes = append(postPb.Dislikes, domain.Dislike{
				Id:     dislike_id,
				UserId: dislike.UserId,
			})
		}
	}

	return postPb
}

func likedPostByUser(post *pb.Post, userId string) bool {
	for _, like := range post.Likes {
		if like.UserId == userId {
			fmt.Println("Postoji duplikat - like")
			fmt.Println("ISPIS:", like.UserId, userId)
			return true
		}
	}
	return false
}

func dislikedPostByUser(post *pb.Post, userId string) bool {
	for _, dislike := range post.Dislikes {
		if dislike.UserId == userId {
			fmt.Println("Postoji duplikat - dislike")
			fmt.Println("ISPIS:", dislike.UserId, userId)
			return true
		}
	}
	return false
}
