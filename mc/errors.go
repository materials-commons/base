package mc

import (
	"errors"
)

var (
	ErrNotFound = errors.New("Not found")
	ErrInvalid  = errors.New("Invalid")
	ErrExists   = errors.New("Exists")
)

type ErrorCode int

const (
	ErrorCodeSuccess ErrorCode = iota
	ErrorCodeNotFound
	ErrorCodeInvalid
	ErrorCodeExists
	ErrorCodeUnknown
)

var errorCodeMapping = map[ErrorCode]error{
	ErrorCodeSuccess:  nil,
	ErrorCodeNotFound: ErrNotFound,
	ErrorCodeInvalid:  ErrInvalid,
	ErrorCodeExists:   ErrExists,
}

func ErrorCodeToError(ec ErrorCode) error {
	return errorCodeMapping[ec]
}

var errorMapping = map[string]ErrorCode{
	ErrNotFound.Error(): ErrorCodeNotFound,
	ErrInvalid.Error():  ErrorCodeInvalid,
	ErrExists.Error():   ErrorCodeInvalid,
}

func ErrorToErrorCode(err error) ErrorCode {
	return errorMapping[err.Error()]
}
