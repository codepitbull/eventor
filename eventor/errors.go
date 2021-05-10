package eventor

import (
	"fmt"
	"reflect"
)

type FunctionNotInInterfaceError struct {
	Interface reflect.Type
	FuncName  string
	error
}

func (e *FunctionNotInInterfaceError) Is(tgt error) bool {
	_, ok := tgt.(*FunctionNotInInterfaceError)
	if !ok {
		return false
	}
	return true
}

func (r *FunctionNotInInterfaceError) Error() string {
	return fmt.Sprintf("Type %v has no function %v", r.Interface, r.FuncName)
}

type DurationWithoutEventFuncError struct {
	Interface reflect.Type
	FuncName  string
	error
}

func (e *DurationWithoutEventFuncError) Is(tgt error) bool {
	_, ok := tgt.(*DurationWithoutEventFuncError)
	if !ok {
		return false
	}
	return true
}

func (r *DurationWithoutEventFuncError) Error() string {
	return fmt.Sprintf("Calling function %v on type %v returned a duration but no EventFunc", r.FuncName, r.Interface)
}
