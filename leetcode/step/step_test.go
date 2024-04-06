package step

import "testing"

func Test_Step(t *testing.T) {
	n := walkStep(10)
	t.Log(n)
}

func Test_StepDP(t *testing.T) {
	n := stepDp(10)
	t.Log(n)
}

func Test_StepDpUp(t *testing.T) {
	n := stepDpUP(10)
	t.Log(n)
}