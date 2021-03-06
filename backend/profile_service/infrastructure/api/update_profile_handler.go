package api

import (
	saga "github.com/velibor7/XML/common/saga/messaging"
	events "github.com/velibor7/XML/common/saga/update_profile"
	"github.com/velibor7/XML/profile_service/application"
)

type UpdateProfileCommandHandler struct {
	profileService    *application.ProfileService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateProfileCommandHandler(profileService *application.ProfileService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileCommandHandler, error) {
	o := &UpdateProfileCommandHandler{
		profileService:    profileService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *UpdateProfileCommandHandler) handle(command *events.UpdateProfileCommand) {
	reply := &events.UpdateProfileReply{
		Profile:      command.Profile,
		Type:         events.UnknownReply,
		OldUsername:  command.OldUsername,
		OldFirstName: command.OldFirstName,
		OldLastName:  command.OldLastName,
		OldIsPrivate: command.OldIsPrivate,
	}
	switch command.Type {
	case events.RollbackUpdatedProfile:
		oldProfile := command.Profile
		oldProfile.Username = command.OldUsername
		oldProfile.FirstName = command.OldFirstName
		oldProfile.LastName = command.OldLastName
		oldProfile.IsPrivate = command.OldIsPrivate
		oldProfile.FullName = command.OldFirstName + " " + command.OldLastName
		err := handler.profileService.RollbackUpdate(mapAuthProfileToProfile(&oldProfile))
		if err != nil {
			return
		}
		reply.Type = events.ProfileUpdateRolledBack
	default:
		reply.Type = events.UnknownReply
	}
	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}
