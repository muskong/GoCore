package helper

import (
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Time    int64  `json:"time"`
}

func Error(message string) (code int, response *Response) {
	code = http.StatusOK
	response = &Response{
		Message: message,
		Time:    time.Now().Unix(),
	}
	return
}

func Success(data any) (code int, response *Response) {
	code = http.StatusOK
	response = &Response{
		Data: data,
		Time: time.Now().Unix(),
	}
	return
}
