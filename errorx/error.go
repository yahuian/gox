package errorx

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
)

type ErrorX struct {
	cause error
	msg   string
	trace string
}

func New(msg string) error {
	return &ErrorX{
		msg:   msg,
		trace: trace(),
		cause: nil,
	}
}

func (e *ErrorX) Error() string {
	if e.cause == nil {
		return e.msg
	}
	if e.msg != "" {
		return e.msg + ": " + e.cause.Error()
	}
	return e.cause.Error()
}

func (e *ErrorX) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			if e.cause != nil {
				fmt.Fprintf(s, "%+v\n", e.cause)
			}
			_, _ = io.WriteString(s, e.trace)
			_, _ = io.WriteString(s, "\n\t")
			_, _ = io.WriteString(s, e.msg)
			return
		}
		fallthrough
	case 's':
		_, _ = io.WriteString(s, e.Error())
	case 'q':
		fmt.Fprintf(s, "%q", e.Error())
	}
}

func (e *ErrorX) Unwrap() error { return e.cause }

func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return &ErrorX{
		cause: err,
		msg:   "",
		trace: trace(),
	}
}

func WrapMsg(msg string, err error) error {
	if err == nil {
		return nil
	}
	return &ErrorX{
		cause: err,
		msg:   msg,
		trace: trace(),
	}
}

func trace() string {
	skip := 2

	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}

	funcName := "unknown"
	pc, _, _, ok := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		funcName = details.Name()
	}

	return funcName + " " + file + ":" + strconv.Itoa(line)
}
