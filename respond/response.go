package respond

import (
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"Message"`
	Data    any    `json:"Data"`
	Time    int64  `json:"Time"`
}

func Message(message string) (code int, response *Response) {
	code = http.StatusOK
	response = &Response{
		Message: message,
		Time:    time.Now().Unix(),
	}
	return
}

func Data(data any) (code int, response *Response) {
	code = http.StatusOK
	response = &Response{
		Data: data,
		Time: time.Now().Unix(),
	}
	return
}
