package api

import (
	"github.com/velibor7/XML/comment_service/application"
	saga "github.com/velibor7/XML/common/saga/messaging"
	events "github.com/velibor7/XML/common/saga/update_profile"
)

type UpdateProfileCommandHandler struct {
	commentService    *application.CommentService
	replyPublisher    saga.Publisher
	commandSubscriber saga.Subscriber
}

func NewUpdateProfileCommandHandler(commentService *application.CommentService, publisher saga.Publisher, subscriber saga.Subscriber) (*UpdateProfileCommandHandler, error) {
	o := &UpdateProfileCommandHandler{
		commentService:    commentService,
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
	case events.UpdateProfile:
		if command.Profile.Username == command.OldUsername {
			return
		}
		err := handler.commentService.UpdateUsername(command.OldUsername, command.Profile.Username)
		if err != nil {
			return
		}
		reply.Type = events.ProfileUpdated
		break
	case events.RollbackUpdatedProfile:
		err := handler.commentService.UpdateUsername(command.OldUsername, command.Profile.Username)
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
