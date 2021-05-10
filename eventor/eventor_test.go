package eventor_test

import (
	"fmt"
	"github.com/codepitbull/eventor/eventor"
	"testing"
	"time"
)

func TestEventorAllOk(t *testing.T) {
	te := TestEventor{}
	processor := eventor.NewEventProcessor("start", map[eventor.Step]eventor.EventFunc{
		"start": te.Start,
		"step2": te.Step2,
		"end": te.End,
	},
	)
	err := processor.Run()
	if err != nil {
		t.Fatal(err)
	}
}

type TestEventor struct {}

func (t *TestEventor) Start() (eventor.Step, *time.Duration, error) {
	fmt.Println("start")
	return "step2", nil, nil
}

func (t *TestEventor) Step2() (eventor.Step, *time.Duration, error) {
	fmt.Println("step2")
	return "end", nil, nil
}

func (t *TestEventor) End() (eventor.Step, *time.Duration, error) {
	fmt.Println("end")
	return "", nil, nil
}
