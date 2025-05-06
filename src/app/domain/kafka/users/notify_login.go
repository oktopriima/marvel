package users

import (
	"encoding/json"
	"github.com/oktopriima/marvel/src/app/usecase/users"
	"github.com/oktopriima/marvel/src/app/usecase/users/dto"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/messages"
	"log"
)

type authHandler struct {
	uc users.UserUsecaseContract
}

type LoginNotificationHandler interface {
	Serve(data *messages.MessageDecoder) error
}

func NewNotifyLoginHandler(uc users.UserUsecaseContract) LoginNotificationHandler {
	return &authHandler{uc: uc}
}

func (a *authHandler) Serve(data *messages.MessageDecoder) error {
	var req dto.NotifyLoginRequest

	if err := json.Unmarshal(data.Body, &req); err != nil {
		log.Fatalf("failed to unmarshal request: %v", err)
		return nil
	}

	resp, err := a.uc.NotifyLogin(data.Context, &req)
	if err != nil {
		log.Fatalf("failed to execute usecase: %v", err)
		return nil
	}

	log.Printf("usecase execute from topic :%s response: %v", data.Topic, resp)
	return nil
}
