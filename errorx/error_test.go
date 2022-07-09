package errorx_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/yahuian/gox/errorx"
)

func foo() error {
	return errorx.WrapMsg("call bar err", bar())
}

func bar() error {
	return errorx.Wrap(read())
}

var errNotFound = errorx.New("not found")

func read() error {
	return errNotFound
}

func TestErrorx(t *testing.T) {
	err := foo()
	simple := `call bar err: not found`
	if fmt.Sprintf("%s", err) != simple {
		t.Errorf("want '%+v' get '%+v'\n", simple, err.Error())
	}

	// NOTE: when your add imports must change line number
	detail := `github.com/yahuian/gox/errorx_test.init github.com/yahuian/gox/errorx/error_test.go:20
not found
github.com/yahuian/gox/errorx_test.bar github.com/yahuian/gox/errorx/error_test.go:17

github.com/yahuian/gox/errorx_test.foo github.com/yahuian/gox/errorx/error_test.go:13
call bar err`
	get := strings.ReplaceAll(fmt.Sprintf("%+v", err), "\t", "")
	if get != detail {
		t.Errorf("want \n%s\n get \n%+v\n", detail, get)
	}

	if errors.Is(err, errNotFound) != true {
		t.Error("errors Is failed")
	}

	var cause *errorx.ErrorX
	if errors.As(err, &cause) != true {
		t.Error("errors As failed")
	}

	if errors.Unwrap(errors.Unwrap(err)).Error() != "not found" {
		t.Error("errors Unwrap failed")
	}
}

func BenchmarkWrap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = foo()
	}
}

func BenchmarkStd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fooStd()
	}
}

func fooStd() error {
	return fmt.Errorf("%w", barStd())
}

func barStd() error {
	return fmt.Errorf("%w", readStd())
}

var errStdNotFound = errors.New("not found")

func readStd() error {
	return errStdNotFound
}
