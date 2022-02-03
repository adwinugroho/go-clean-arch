package models

import "encoding/json"

type Error struct {
	Status       bool   `json:"status"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"errMessage"`
}

func NewError(code int, message string) *Error {
	return &Error{Status: false, Code: code, ErrorMessage: message}
}

func (e *Error) Error() string {
	errBytes, _ := json.Marshal(e)
	return string(errBytes)
}
