package api

import (
	"context"

	"github.com/velibor7/XML/common/loggers"
	pbComment "github.com/velibor7/XML/common/proto/comment_service"
	pb "github.com/velibor7/XML/common/proto/profile_service"
	"github.com/velibor7/XML/common/tracer"
	"github.com/velibor7/XML/profile_service/application"
)

var log = loggers.NewProfileLogger()

type ProfileHandler struct {
	pb.UnimplementedProfileServiceServer
	service       *application.ProfileService
	commentClient pbComment.CommentServiceClient
}

func NewProfileHandler(service *application.ProfileService, commentClient pbComment.CommentServiceClient) *ProfileHandler {
	return &ProfileHandler{
		service:       service,
		commentClient: commentClient,
	}
}

func (handler *ProfileHandler) Get(ctx context.Context, request *pb.GetRequest) (*pb.GetResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "Get")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)

	Profile, err := handler.service.Get(request.Id)
	if err != nil {
		log.WithField("profileId", request.Id).Errorf("Cannot get profile: ", err)
		return nil, err
	}
	ProfilePb := mapProfileToPb(Profile)

	response := &pb.GetResponse{
		Profile: ProfilePb,
	}
	return response, nil
}

func (handler *ProfileHandler) GetAll(ctx context.Context, request *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "GetAll")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	Profiles, err := handler.service.GetAll()
	if err != nil {
		log.Errorf("Can't get all profiles: ", err)
		return nil, err
	}
	response := &pb.GetAllResponse{
		Profile: []*pb.Profile{},
	}
	for _, Profile := range Profiles {
		current := mapProfileToPb(Profile)
		response.Profile = append(response.Profile, current)
	}
	return response, nil
}

func (handler ProfileHandler) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "Create")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	profile := mapPbToProfile(request.Profile)
	err := handler.service.Create(profile)
	if err != nil {
		log.Errorf("Can't create profile: ", err)
		return nil, err
	}
	log.Info("Profile created: ", profile.Username)
	return &pb.CreateResponse{
		Profile: mapProfileToPb(profile),
	}, nil
}

func (handler ProfileHandler) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	span := tracer.StartSpanFromContext(ctx, "Update")
	defer span.Finish()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	profile := mapPbToProfile(request.Profile)
	err := handler.service.Update(request.Id, profile)
	if err != nil {
		log.Errorf("Can't update profile: ", err)
		return nil, err
	}
	log.Info("Profile updated: ", profile.Username)
	return &pb.UpdateResponse{
		Profile: mapProfileToPb(profile),
	}, nil
}
