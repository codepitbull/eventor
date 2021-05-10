package eventor

import (
	"reflect"
	"runtime"
	"strings"
	"time"
)

func GetFunctionName(i interface{}) string {
	funcname := strings.TrimSuffix(runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name(), "-fm")
	nameComps := strings.Split(funcname, ".")
	return nameComps[len(nameComps)-1]
}

func CallFunctionByName(i interface{}, name string) (EventFunc, error) {
	t := reflect.ValueOf(i)
	m := t.MethodByName(name)
	if !m.IsValid() {
		return nil, &FunctionNotInInterfaceError{Interface: t.Type(), FuncName: name}
	}
	args := []reflect.Value{}
	result := m.Call(args)

	functioneValue := result[0]
	durationValue := result[1]
	errorValue := result[2]

	if !errorValue.IsNil() {
		return nil, result[2].Interface().(error)
	} else if !durationValue.IsNil() {
		if !functioneValue.IsNil() {
			return functioneValue.Interface().(EventFunc), nil
		} else {
			return nil, &DurationWithoutEventFuncError{Interface: t.Type(), FuncName: name}
		}
	} else if !functioneValue.IsNil() {
		return functioneValue.Interface().(EventFunc), nil
	}
	//We reached the end
	return nil, nil
}

type EventFunc func() (EventFunc, *time.Duration, error)

type EventProcessor struct {
	ToRun         interface{}
	NextExecution EventFunc
}

func (f EventProcessor) Run() error {
	for f.NextExecution != nil {
		next, err := CallFunctionByName(f.ToRun, GetFunctionName(f.NextExecution))
		if err != nil {
			return err
		}
		f.NextExecution = next
	}
	return nil
}
