package handlers

import "errors"

var (
	ErrDevelopmentParametersNotEnough = errors.New("development parameters not enough")
	ErrDevelopmentParametersInvalid   = errors.New("development parameters invalid")
)
