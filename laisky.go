package errors

import (
	"fmt"
	"log"
	"sync/atomic"
)

var skipCallers int32 = 3

func panicErr(err error) {
	if err != nil {
		log.Panicf("%+v", err)
	}
}

// SetSkipCallers set depth of skip callers
func SetSkipCallers(skip int) {
	atomic.StoreInt32(&skipCallers, int32(skip))
}

// GetSkipCallers get depth of skip callers
func GetSkipCallers() int {
	return int(atomic.LoadInt32(&skipCallers))
}

// WithStackWithSkip like WithStackm but you can skip arbitrary stacks
func WithStackWithSkip(skip int, err error) error {
	if err == nil {
		return nil
	}

	return &withStack{
		err,
		callers(skip),
	}
}

// WrapWithSkip same as Wrap, but you can skip arbitrary stacks
func WrapWithSkip(skip int, err error, message string) error {
	if err == nil {
		return nil
	}
	err = &withMessage{
		cause: err,
		msg:   message,
	}
	return &withStack{
		err,
		callers(skip),
	}
}

// WrapfWithSkip same as Wrapf, but you can skip arbitrary stacks
func WrapfWithSkip(skip int, err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	err = &withMessage{
		cause: err,
		msg:   fmt.Sprintf(format, args...),
	}
	return &withStack{
		err,
		callers(skip),
	}
}

// ErrorfWithStack same as Errorf, but you can skip arbitrary stacks
func ErrorfWithStack(skip int, format string, args ...interface{}) error {
	return &fundamental{
		msg:   fmt.Sprintf(format, args...),
		stack: callers(skip),
	}
}
