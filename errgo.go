package errgo

import (
	"fmt"
	"reflect"
	"sync"
)

const (
	NO_TYPE = -1
)

var (
	errs errMap
)

func init() {
	errs = errMap{
		m: make(map[int]Err),
	}
}

type errMap struct {
	// Map between error types, and their assosciated error
	m map[int]Err
	sync.RWMutex
}

type Err struct {
	// Error type
	errType int
	// error string
	str string
}

func New(e interface{}, params ...interface{}) Err {
	if reflect.ValueOf(e).Type().String() == "string" {
		return Err{errType: NO_TYPE, str: e.(string)}
	} else if reflect.ValueOf(e).Type().String() == "int" {
		errs.RLock()
		defer errs.RUnlock()
		er := errs.m[e.(int)]
		er.str = fmt.Sprintf(er.str, params...)
		return er
	}
	panic("Cannot create error as defined")
}

func Register(errType int, str string) {
	errs.Lock()
	defer errs.Unlock()
	errs.m[errType] = Err{
		errType: errType,
		str:     str,
	}
}

func (e Err) String() string {
	return e.str
}

func (e Err) IsType(errType int) bool {
	return e.errType == errType
}

func (e Err) Type() int {
	return e.errType
}
