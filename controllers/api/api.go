package api

import (
	"DMAPI/logger"
	"math/rand"
	"time"
)

func ResponseError(err string, code int64) Response {
	logger.Error.Println(err, code)
	return Response{
		Error: &Error{
			Message: err,
			Code:    code,
		},
	}
}

type Response struct {
	Result interface{} `json:"result"`
	Error  *Error      `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Code    int64  `json:"code"`
	Obj     any    `json:"obj,omitempty"`
}

func RandString(i int64) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, i)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
