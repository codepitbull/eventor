package eventor_test

import (
	"errors"
	"testing"
	"time"

	"github.com/codepitbull/eventor/eventor"
)

func TestEventorAllOk(t *testing.T) {
	str, start := NewTesteEventor()
	err := eventor.EventProcessor{ToRun: str, NextExecution: start}.Run()
	if err != nil {
		t.Fatal(err)
	}
}

func TestEventorWithMissingFunc(t *testing.T) {
	str, start := NewEventorWithMissingFunc()
	err := eventor.EventProcessor{ToRun: str, NextExecution: start}.Run()
	if err == nil {
		t.Fatal("This should have failed")
	}
	if err != nil && !errors.Is(err, &eventor.FunctionNotInInterfaceError{}) {
		t.Fatalf("Expected a FunctionNotInInterfaceError but got %s", err)
	}
}

func TestEventorWithBrokenReturn(t *testing.T) {
	str, start := NewEventorWithBrokenReturn()
	err := eventor.EventProcessor{ToRun: str, NextExecution: start}.Run()
	if err == nil {
		t.Fatal("This should have failed")
	}
	if err != nil && !errors.Is(err, &eventor.DurationWithoutEventFuncError{}) {
		t.Fatalf("Expected a DurationWithoutEventFuncError but got %s", err)
	}
}

func NewTesteEventor() (TestEvents, eventor.EventFunc) {
	dem := TesteEventor{}
	return dem, dem.Start
}

type TesteEventor struct {
}

type TestEvents interface {
	Start() (eventor.EventFunc, *time.Duration, error)
	Step2() (eventor.EventFunc, *time.Duration, error)
	End() (eventor.EventFunc, *time.Duration, error)
}

func (f TesteEventor) Start() (eventor.EventFunc, *time.Duration, error) {
	return f.Step2, nil, nil
}

func (f TesteEventor) Step2() (eventor.EventFunc, *time.Duration, error) {
	return f.End, nil, nil
}

func (f TesteEventor) End() (eventor.EventFunc, *time.Duration, error) {
	return nil, nil, nil
}

func NewEventorWithMissingFunc() (EventorWithMissingFuncStruct, eventor.EventFunc) {
	dem := EventorWithMissingFuncStruct{}
	return dem, dem.Start
}

type EventorWithMissingFuncStruct struct {
}

type EventorWithMissingFunc interface {
	Start() (eventor.EventFunc, *time.Duration, error)
}

func funcNotInInterface() (eventor.EventFunc, *time.Duration, error) {
	return nil, nil, nil
}

func (f EventorWithMissingFuncStruct) Start() (eventor.EventFunc, *time.Duration, error) {
	return funcNotInInterface, nil, nil
}

func NewEventorWithBrokenReturn() (EventorWithBrokenReturnStruct, eventor.EventFunc) {
	dem := EventorWithBrokenReturnStruct{}
	return dem, dem.Start
}

type EventorWithBrokenReturnStruct struct {
}

type EventorWithBrokenReturn interface {
	Start() (eventor.EventFunc, *time.Duration, error)
}

func (f EventorWithBrokenReturnStruct) Start() (eventor.EventFunc, *time.Duration, error) {
	duration := 500 * time.Millisecond
	return nil, &duration, nil
}
