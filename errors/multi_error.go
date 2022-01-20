package errors

import (
	"fmt"
	"sync"
)

type Errs interface {
	error
	Occurred(err error) Errs
	AnyOccurred() bool
}

type errs struct {
	errs []error
	mx   sync.RWMutex
}

func (errs *errs) Occurred(err error) Errs {
	errs.mx.Lock()
	defer errs.mx.Unlock()
	errs.errs = append(errs.errs, err)
	return errs
}

func (errs *errs) Error() string {
	errs.mx.RLock()
	defer errs.mx.RUnlock()
	msg := "errors occurred:\n"
	for _, e := range errs.errs {
		msg += fmt.Sprintf("\t%s\n", e.Error())
	}
	return msg
}

func (errs *errs) AnyOccurred() bool {
	errs.mx.RLock()
	defer errs.mx.RUnlock()
	return len(errs.errs) > 0
}

func Errors() *errs {
	return &errs{errs: []error{}}
}
