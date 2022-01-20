package errors

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type ChainedError interface {
	CausedBy() ChainedError
	// msg instanceof error or string
	Cause(msg interface{}, v ...interface{}) ChainedError
	At() (file string, line int, method string)
	Print()
	Error() string
	String() string
}

type chainedError struct {
	msg      string
	causedBy ChainedError
	inMethod string
	inFile   string
	inLine   int
}

func Ensure(err error, causedBy ChainedError) ChainedError {
	if errors.Is(causedBy, err) {
		return New(err, nil)
	}
	return causedBy.Cause(err)
}

func New(msg interface{}, causedBy ChainedError) ChainedError {
	return newWithMsg(msg, causedBy)
}

func newWithMsg(msg interface{}, e ChainedError, v ...interface{}) *chainedError {
	switch m := msg.(type) {
	case string:
		return NewChainedError(fmt.Sprintf(m, v...), e)
	case error:
		return NewChainedError(m.Error(), e)
	default:
		panic(fmt.Sprintf("Unsupported msg type [%s]", reflect.ValueOf(m).Elem().Kind()))
	}
}

func NewChainedError(msg string, causedBy ChainedError) *chainedError {
	e := &chainedError{msg: msg, causedBy: causedBy}
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	if n == 0 {
		return e
	}
	pc = pc[:n]
	frames := runtime.CallersFrames(pc)
	chainedErrorShowed := false
	for {
		frame, _ := frames.Next()
		if strings.Index(frame.File, "chained_error.go") == -1 {
			if chainedErrorShowed {
				e.inFile = frame.File
				e.inLine = frame.Line
				e.inMethod = frame.Function
				return e
			}
		} else if !chainedErrorShowed {
			chainedErrorShowed = true
		}
	}
}

func (e *chainedError) Error() string {
	return e.msg
}

func (e *chainedError) CausedBy() ChainedError {
	return e.causedBy
}

func (e *chainedError) At() (file string, line int, method string) {
	return e.inFile, e.inLine, e.inMethod
}

func (e *chainedError) Cause(msg interface{}, v ...interface{}) ChainedError {
	return newWithMsg(msg, e, v...)
}

func (e *chainedError) Print() {
	fmt.Print(e.String())
}

func (e *chainedError) String() string {
	ret := ""
	var t ChainedError = e
	for {
		if e == t {
			ret += "ERROR: "
		} else {
			ret += "\t  "
		}
		_, _, inMethod := t.At()
		ret += fmt.Sprintf("\"%s\" in %s\n", t.Error(), inMethod)
		t = t.CausedBy()
		if t == nil {
			break
		}
		ret += "\tCaused by\n"
	}
	return ret
}
