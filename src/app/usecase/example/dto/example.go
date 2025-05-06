package dto

type Request struct {
	Messages string `json:"messages"`
}

type response struct {
	Messages string `json:"messages"`
}

type Response interface {
	GetObject() *response
}

func (r *response) GetObject() *response {
	return r
}

func CreateResponse(message string) Response {
	return &response{
		Messages: message,
	}
}
