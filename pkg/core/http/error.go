package http

import "errors"

var (
	ErrConnect      = errors.New("connect failed")
	ErrUnauthorized = errors.New("unauthorized")
	ErrInvalid      = errors.New("invalid user")

	ErrAccessKeyUnauthorized = errors.New("access key unauthorized")
	ErrConnectApi            = errors.New("api connect err")
)
