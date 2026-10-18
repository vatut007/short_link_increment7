package service

import "errors"

var (
	ErrNoType     = errors.New("undefined content type")
	ErrEmptyUrl   = errors.New("Url is empty")
	ErrInvalidUrl = errors.New("Invalid url")
)
