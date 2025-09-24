package response

import (
	"fiber-crud-demo/dto"
	"net/http"
	"time"
)

type Response struct {
	Path      string       `json:"path,omitempty" example:"/v1/m_biodata"`
	Timestamp dto.JSONTime `json:"timestamp" example:"2024-02-16 10:33:10" swaggertype:"string"`
	Status    int          `json:"status" example:"200"`
	Message   string       `json:"message" example:"success"`
	Data      interface{}  `json:"data,omitempty" swaggertype:"string"`
	Error     interface{}  `json:"error,omitempty" swaggertype:"string"`
}

func (response *Response) Ok(path string, data interface{}) {
	response.Timestamp = dto.JSONTime{Time: time.Now()}
	response.Data = data
	response.Status = http.StatusOK
	response.Message = "success"
	response.Path = path
}

func (response *Response) Err(path string, error interface{}, status int) {
	response.Timestamp = dto.JSONTime{Time: time.Now()}
	response.Status = status
	response.Message = "error"
	response.Path = path
	response.Error = error
}

func (response *Response) ErrMessage(path string, status int, message string) {
	response.Timestamp = dto.JSONTime{Time: time.Now()}
	response.Status = status
	response.Message = message
	response.Path = path
}

func (response *Response) ErrMessagePayload(path string, status int, message string, error interface{}) {
	response.Timestamp = dto.JSONTime{Time: time.Now()}
	response.Status = status
	response.Message = message
	response.Path = path
	response.Error = error
}
