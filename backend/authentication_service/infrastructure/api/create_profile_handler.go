package api

import (
	"github.com/velibor7/XML/authentication_service/application"
	events "github.com/velibor7/XML/common/saga/create_profile"
	saga "github.com/velibor7/XML/common/saga/messaging"
)

type CreateProfileCommandHandler struct {
	authService       *application.AuthService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewCreateProfileCommandHandler(authService *application.AuthService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateProfileCommandHandler, error) {
	o := &CreateProfileCommandHandler{
		authService:       authService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateProfileCommandHandler) handle(command *events.CreateProfileCommand) {
	reply := &events.CreateProfileReply{Type: events.UnknownReply}
	switch command.Type {
	case events.RollbackCreatedProfile:
		err := handler.authService.Delete(command.Profile.Id)
		if err != nil {
			return
		}
		reply.Type = events.ProfileCreationRolledBack
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
