package errors

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

// ChainedError a chained error definition
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

var runtimeInfoEnabled bool
var mx sync.RWMutex

func init() {
	runtimeInfoEnabled = true
}

// EnableRuntimeInfo if runtime info enabled, it will collect infomation about where the error occured
func EnableRuntimeInfo() {
	mx.Lock()
	defer mx.Unlock()
	runtimeInfoEnabled = true
}

// DisableRuntimeInfo disable collect runtime info for performance issue
func DisableRuntimeInfo() {
	mx.Lock()
	defer mx.Unlock()
	runtimeInfoEnabled = false
}

// Ensure make sure the error is the first arguement typed errs.ChainedError
// if err := DoSomething(); err != nil {
//     return errs.Ensure(err, errs.New("do something error"))
// }
func Ensure(err error, causedBy ChainedError) ChainedError {
	if errors.Is(causedBy, err) {
		return New(err)
	}
	return causedBy.Cause(err)
}

// New new a ChainedError, this func implement with reflect to make it easy to use but have performance issue
func New(msg interface{}) ChainedError {
	return newWithMsg(msg, nil)
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

// New new a ChainedError with msg string
func NewChainedError(msg string, causedBy ChainedError) *chainedError {
	e := &chainedError{msg: msg, causedBy: causedBy}
	if runtimeInfoEnabled {
		mx.RLock()
		defer mx.RUnlock()
		pc := make([]uintptr, 6)
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
	return e
}

// Error return the error message not the whole chained error message
func (e *chainedError) Error() string {
	return e.msg
}

// CausedBy give who cause the error
func (e *chainedError) CausedBy() ChainedError {
	return e.causedBy
}

// At get where the error occured
func (e *chainedError) At() (file string, line int, method string) {
	return e.inFile, e.inLine, e.inMethod
}

// Cause give which error will be caused
func (e *chainedError) Cause(msg interface{}, v ...interface{}) ChainedError {
	return newWithMsg(msg, e, v...)
}

// Print print the chained error detail
func (e *chainedError) Print() {
	fmt.Print(e.String())
}

// String stringify the whole error chain
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
		if inMethod != "" {
			ret += fmt.Sprintf("\"%s\" in %s\n", t.Error(), inMethod)
		} else {
			ret += fmt.Sprintf("\"%s\"\n", t.Error())
		}
		t = t.CausedBy()
		if t == nil {
			break
		}
		ret += "\tCaused by\n"
	}
	return ret
}
