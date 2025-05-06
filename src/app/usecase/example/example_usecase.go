package example

import "github.com/oktopriima/marvel/src/app/usecase/example/dto"

type exampleUsecase struct {
}

type UsecaseContract interface {
	Execute(request dto.Request) (dto.Response, error)
}

func NewUsecase() UsecaseContract {
	return &exampleUsecase{}
}

func (e *exampleUsecase) Execute(request dto.Request) (dto.Response, error) {
	// business logic happens here

	resp := dto.CreateResponse(request.Messages)
	return resp, nil
}
