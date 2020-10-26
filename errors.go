package errors

import (
	"errors"
	"fmt"
)

// 保留errType 定义
const (
	// NoErr 无err
	NoErr = 0
	// UnknowErr 未知err
	UnknowErr = -1
)

// Typer 带有业务错误类型的 error
type Typer interface {
	error
	Type() int32
}

// Error 业务错误
type Error struct {
	typ int32  //错误类型
	msg string //错误信息

	cause error //要包装的原因
}

// New 创建一个新的Error
func New(typ int32, msg string) error {
	return &Error{
		typ: typ,
		msg: msg,
	}
}

// Errorf 创建Error
func Errorf(typ int32, format string, a ...interface{}) error {
	return &Error{
		typ: typ,
		msg: fmt.Sprintf(format, a...),
	}
}

// Wrap 返回一个Error 并嵌套原有error
func Wrap(typ int32, err error, msg string) error {
	return &Error{
		typ: typ,
		msg: msg,

		cause: err,
	}
}

// WithMsg 仅包装信息，不改变下游的error type
func WithMsg(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// Wrapf 返回一个Error 并嵌套原有error
func Wrapf(typ int32, err error, format string, a ...interface{}) error {
	return &Error{
		typ: typ,
		msg: fmt.Sprintf(format, a...),

		cause: err,
	}
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}

	if e.cause == nil {
		return e.msg
	}
	return fmt.Sprintf("%s: %s", e.msg, e.cause.Error())
}

// Type 返回错误类型
func (e *Error) Type() int32 {
	if e == nil {
		return NoErr
	}
	return e.typ
}

func (e *Error) Unwrap() error {
	return e.cause
}

// Type 返回error 的业务类型type
func Type(err error) int32 {
	if err == nil { //error 为空认为无错误
		return NoErr
	}

	if e := Typer(nil); errors.As(err, &e) {
		return e.Type()
	}

	return UnknowErr
}
