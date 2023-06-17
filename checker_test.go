package main

import "testing"

func TestIfFailsWhenArgIsPassedAfterArg(t *testing.T) {
	args := []string{"--framework", "--styles"}

	if got := checkArgs(args); got == nil {
		t.Errorf("analyseArgs(args) = %q, want error", got)
	}
}

func TestIfFailsWhenNoValueIsGivenForArg(t *testing.T) {
	args := []string{"--framework", ""}

	if got := checkArgs(args); got == nil {
		t.Errorf("analyseArgs(args) = %q, want error", got)
	}
}

func TestIfFailsWhenGroupedShorthandIsUsed(t *testing.T) {
	args := []string{"-fs"}

	if got := checkArgs(args); got == nil {
		t.Errorf("analyseArgs(args) = %q, want error", got)
	}
}

func TestIfFailsWhenNoValueIsPassedToShorthand(t *testing.T) {
	args := []string{"-f"}

	if got := checkArgs(args); got == nil {
		t.Errorf("analyseArgs(args) = %q, want error", got)
	}
}

func TestIfFailsWhenTwoShorthandIsPassedAfterTheOther(t *testing.T) {
	args := []string{"-f", "-s"}

	if got := checkArgs(args); got == nil {
		t.Errorf("analyseArgs(args) = %q, want error", got)
	}
}

func TestIfFailsWhenShortHandIsPassedAfterArg(t *testing.T) {
	args := []string{"--framework", "-s", "tailwind"}

	if got := checkArgs(args); got == nil {
		t.Errorf("analyseArgs(args) = %q, want error", got)
	}
}

func TestIfSuccedsWhenAValueIsPassedToArg(t *testing.T) {
	args := []string{"--framework", "react"}

	if got := checkArgs(args); got != nil {
		t.Errorf("analyseArgs(args) = %q, want nil", got)
	}
}

func TestIfSuccedsWhenHelpArgIsPassedWithNoValue(t *testing.T) {
	args := []string{"--help"}

	if got := checkArgs(args); got != nil {
		t.Errorf("analyseArgs(args) = %q, want nil", got)
	}
}
