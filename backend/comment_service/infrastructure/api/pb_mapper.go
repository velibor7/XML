package api

import (
	"github.com/velibor7/XML/comment_service/domain"
	pb "github.com/velibor7/XML/common/proto/comment_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func mapComment(comment *domain.Comment) *pb.Comment {
	commentPb := &pb.Comment{
		Id:       comment.Id.Hex(),
		PostId:   comment.PostId.Hex(),
		Username: comment.Username,
		Content:  comment.Content,
	}
	return commentPb
}

func mapPb(pbComent *pb.Comment) *domain.Comment {
	comment := &domain.Comment{
		Id:       getObjectId(pbComent.Id),
		PostId:   getObjectId(pbComent.PostId),
		Username: pbComent.Username,
		Content:  pbComent.Content,
	}
	return comment
}

func getObjectId(id string) primitive.ObjectID {
	if objectId, err := primitive.ObjectIDFromHex(id); err == nil {
		return objectId
	}
	return primitive.NewObjectID()
}
