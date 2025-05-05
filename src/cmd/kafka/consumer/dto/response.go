package dto

// Response ...
type Response struct {
	Name    string      `json:"name"`
	Message interface{} `json:"message,omitempty"`
	Error   error       `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Lang    string      `json:"-"`
}
