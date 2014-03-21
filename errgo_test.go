package errgo

import (
	"testing"
)

const (
	custom_err_type = iota
	custom_err_type_2
)

func setup() {
	Register(custom_err_type, "Custom err type")
}

func cleanup() {
	delete(errs.m, custom_err_type)
}

func TestBaseErrStr(t *testing.T) {
	e := New("Test error")
	if e.String() != "Test error" {
		t.Errorf("Failed! Exptexted %v, got %v", "Test error", e.String())
	}
}

func TestBaseErrType(t *testing.T) {
	e := New("Test error")
	if !e.IsType(NO_TYPE) {
		t.Errorf("Failed! Expected %v, got %v", NO_TYPE, e.Type())
	}
}

func TestCustomErrStr(t *testing.T) {
	setup()
	defer cleanup()
	e := New(custom_err_type)
	if e.String() != "Custom err type" {
		t.Errorf("Failed: Expected %v, got %v", "Custom err type", e.String())
	}
}

func TestCustomErrType(t *testing.T) {
	setup()
	defer cleanup()
	e := New(custom_err_type)
	if e.Type() != custom_err_type {
		t.Errorf("Failed! Expected %v, got %v", custom_err_type, e.Type())
	}
}

func TestCustomErrTypeParams(t *testing.T) {
	Register(custom_err_type_2, "Custom err of type %v")
	defer func() { delete(errs.m, custom_err_type_2) }()
	e := New(custom_err_type_2, "2")
	if e.String() != "Custom err of type 2" {
		t.Errorf("Failed! Exptected %v, got %v", "Custom err of type 2", e.String())
	}
}
