package api

import (
	"github.com/velibor7/XML/authentication_service/domain"
	pb "github.com/velibor7/XML/common/proto/authentication_service"
)

func mapUserCredential(userCredential *domain.UserCredential) *pb.UserCredential {
	userCredentialPb := &pb.UserCredential{
		Username: userCredential.Username,
		Password: userCredential.Password,
	}
	return userCredentialPb
}

func mapPbUserCredential(userCredential *pb.UserCredential) *domain.UserCredential {
	userCredentialPb := &domain.UserCredential{
		Username: userCredential.Username,
		Password: userCredential.Password,
	}
	return userCredentialPb
}
