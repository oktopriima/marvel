package example

import (
	"fmt"
	"github.com/oktopriima/marvel/src/cmd/kafka/consumer/dto"
)

type exampleUsecase struct {
}

type Usecase interface {
	Serve(data *dto.MessageDecoder) error
}

func NewExampleUsecase() Usecase {
	return &exampleUsecase{}
}

func (e *exampleUsecase) Serve(data *dto.MessageDecoder) error {
	fmt.Printf("incoming message %v", data.Cast(data.Body))

	return nil
}
