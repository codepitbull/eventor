package eventor_test

import (
	"fmt"
	"github.com/codepitbull/eventor/eventor"
	"testing"
	"time"
)

func TestEventorAllOk(t *testing.T) {
	te := TestEventor{}
	processor := eventor.NewEventProcessor(
		te.Start,
		te.Step2,
		te.End,
	)
	err := processor.Run()
	if err != nil {
		t.Fatal(err)
	}
}

type TestEventor struct {}

func (t *TestEventor) Start() (*time.Duration, error) {
	fmt.Println("start")
	return nil, nil
}

func (t *TestEventor) Step2() (*time.Duration, error) {
	fmt.Println("step2")
	return nil, nil
}

func (t *TestEventor) End() (*time.Duration, error) {
	fmt.Println("end")
	return nil, nil
}
