package gopherpc

import "errors"

type ErrCode int

const (
	ErrParse          ErrCode = -32700
	ErrInvalidRequest ErrCode = -32600
	ErrMethodNotFound ErrCode = -32601
	ErrInvalidParams  ErrCode = -32602
	ErrInternalError  ErrCode = -32603
	//
	errorRegexString = `\S{0,}error\S{0,}:{\S{0,}code\S{0,}:-32[0-9][0-9][0-9],\S{0,}message\S{0,}:.{0,}}`
)

var (
	IdMustBeIntOrNil          error = errors.New("ID must be INT or NULL.")
	NotificationHasntResponse error = errors.New("Notification hasn't response.")
	NotSupportedProtoVersion  error = errors.New("Proto. version not supported.")
	NoneFindTemplate          error = errors.New("None find template.")
)
