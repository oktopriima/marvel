package example

import (
	"github.com/oktopriima/marvel/src/app/usecase/example"
	"github.com/oktopriima/marvel/src/app/usecase/example/dto"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/messages"
	"log"
)

type handlerExample struct {
	uc example.UsecaseContract
}

type Handler interface {
	Serve(data *messages.MessageDecoder) error
}

func NewExampleHandler(uc example.UsecaseContract) Handler {
	return &handlerExample{
		uc: uc,
	}
}

func (h *handlerExample) Serve(data *messages.MessageDecoder) error {
	var req dto.Request
	req.Messages = string(data.Body)

	resp, err := h.uc.Execute(req)
	if err != nil {
		log.Fatalf("failed to execute usecase: %v", err)
		return nil
	}

	log.Printf("usecase response: %v", resp.GetObject())
	return nil
}
