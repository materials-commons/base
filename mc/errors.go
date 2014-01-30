package mc

import (
	"errors"
)

var (
	ErrNotFound = errors.New("Not found")
	ErrInvalid  = errors.New("Invalid")
	ErrExists   = errors.New("Exists")
	ErrNoAccess = errors.New("No access")
	ErrCreate   = errors.New("Unable to create")
	ErrInternal = errors.New("Internal error")
)

type ErrorCode int

const (
	ErrorCodeSuccess ErrorCode = iota
	ErrorCodeNotFound
	ErrorCodeInvalid
	ErrorCodeExists
	ErrorCodeNoAccess
	ErrorCodeCreate
	ErrorCodeInternal
	ErrorCodeUnknown
)

var errorCodeMapping = map[ErrorCode]error{
	ErrorCodeSuccess:  nil,
	ErrorCodeNotFound: ErrNotFound,
	ErrorCodeInvalid:  ErrInvalid,
	ErrorCodeExists:   ErrExists,
	ErrorCodeNoAccess: ErrNoAccess,
	ErrorCodeCreate: ErrCreate,
	ErrorCodeInternal: ErrInternal,
}

func ErrorCodeToError(ec ErrorCode) error {
	return errorCodeMapping[ec]
}

var errorMapping = map[string]ErrorCode{
	ErrNotFound.Error(): ErrorCodeNotFound,
	ErrInvalid.Error():  ErrorCodeInvalid,
	ErrExists.Error():   ErrorCodeExists,
	ErrNoAccess.Error(): ErrorCodeNoAccess,
	ErrCreate.Error(): ErrorCodeCreate,
	ErrInternal.Error(): ErrorCodeInternal,
}

func ErrorToErrorCode(err error) ErrorCode {
	return errorMapping[err.Error()]
}
