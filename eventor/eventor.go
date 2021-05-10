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

	if !result[2].IsNil() {
		//we got an error
		return nil, result[2].Interface().(error)
	} else if !result[1].IsNil() {
		if !result[0].IsNil() {
			//TODO: add the delay (e.g. reschedule in k8s-eventloop)
			return result[0].Interface().(EventFunc), nil
		} else {
			//We got a duration without an EventFunc => Error
			return nil, &DurationWithoutEventFuncError{Interface: t.Type(), FuncName: name}
		}
	} else if !result[0].IsNil() {
		//We got a func without a duration to wait => Execute
		return result[0].Interface().(EventFunc), nil
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
