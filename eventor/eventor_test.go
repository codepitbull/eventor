package eventor_test

import (
	"fmt"
	"github.com/codepitbull/eventor/eventor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestEventorAllOk(t *testing.T) {
	te := &TestEventor{}
	processor := eventor.NewEventProcessor("start", map[eventor.Step]eventor.EventFunc{
		"start": te.Start,
		"step2": te.Step2,
		"end":   te.End,
	},
	)
	got, err := processor.Run()
	require.NoError(t, err)
	assert.Equal(t, eventor.Step("end"), got)
	assert.Equal(t, []eventor.Step{"start", "step2"}, te.stepsRun)

	got, err = processor.RunFrom(got)
	require.NoError(t, err)
	assert.Equal(t, eventor.Step(""), got)
	assert.Equal(t, []eventor.Step{"start", "step2", "end"}, te.stepsRun)
}

type TestEventor struct {
	stepsRun []eventor.Step
}

func (t *TestEventor) Start() (eventor.Step, time.Duration, error) {
	step := eventor.Step("start")
	fmt.Println(step)
	t.stepsRun = append(t.stepsRun, step)
	return "step2", 0, nil
}

func (t *TestEventor) Step2() (eventor.Step, time.Duration, error) {
	step := eventor.Step("step2")
	fmt.Println(step)
	t.stepsRun = append(t.stepsRun, step)
	return "end", 3 * time.Second, nil
}

func (t *TestEventor) End() (eventor.Step, time.Duration, error) {
	step := eventor.Step("end")
	fmt.Println(step)
	t.stepsRun = append(t.stepsRun, step)
	return "", 0, nil
}
