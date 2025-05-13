package example

import (
	"github.com/oktopriima/marvel/pkg/pubsubrouter"
	"github.com/oktopriima/marvel/src/app/usecase/example"
	"github.com/oktopriima/marvel/src/app/usecase/example/dto"
	"log"
)

type handlerExample struct {
	uc example.UsecaseContract
}

type EventProcessor interface {
	Serve(m *pubsubrouter.Message) error
}

func NewHandler(uc example.UsecaseContract) EventProcessor {
	return &handlerExample{
		uc: uc,
	}
}

func (h *handlerExample) Serve(m *pubsubrouter.Message) error {
	var req dto.Request
	req.Messages = string(m.Payload.Data)

	resp, err := h.uc.Execute(req)
	if err != nil {
		log.Fatalf("failed to execute usecase: %v", err)
		return nil
	}

	log.Printf("usecase response: %v", resp.GetObject())
	return nil
}
