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
		m: make(map[int]err),
	}
}

type errMap struct {
	// Map between error types, and their assosciated error
	m map[int]err
	sync.RWMutex
}

type err struct {
	// Error type
	errType int
	// error string
	str string
}

func New(e interface{}, params ...interface{}) err {
	if reflect.ValueOf(e).Type().String() == "string" {
		return err{errType: NO_TYPE, str: e.(string)}
	} else if reflect.ValueOf(e).Type().String() == "int" {
		errs.RLock()
		defer errs.RUnlock()
		er := errs.m[e.(int)]
		er.str = fmt.Sprintf(er.str, params...)
		return er
	}
	panic("Cannot create error as defined")
}

/*func NewType(errType int, params ...interface{}) err {
	errs.RLock()
	defer errs.RUnlock()
	e := errs.m[errType]
	e.str = fmt.Sprintf(e.str, params...)
	return e
}*/

func Register(errType int, str string) {
	errs.Lock()
	defer errs.Unlock()
	errs.m[errType] = err{
		errType: errType,
		str:     str,
	}
}

func (e err) String() string {
	return e.str
}

func (e err) IsType(errType int) bool {
	return e.errType == errType
}

func (e err) Type() int {
	return e.errType
}
