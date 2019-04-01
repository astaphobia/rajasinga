package helpers

import "encoding/json"

type Response struct {
	StatusCode int
	Data       interface{}
	Message    interface{}
}

func (r Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		StatusCode int         `json:"status_code"`
		Data       interface{} `json:"data,omitempty"`
		Message    interface{} `json:"message,omitempty"`
	}{
		StatusCode: r.StatusCode,
		Data:       r.Data,
		Message:    r.Message,
	})
}

func NewResponse(statusCode int, data interface{}, message interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Data:       data,
		Message:    message,
	}
}
