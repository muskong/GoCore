package respond

import (
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Time    int64  `json:"time"`
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
