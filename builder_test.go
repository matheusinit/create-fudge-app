package main

import "testing"
import "strings"

func TestIfGetAppNameDefaultValue(t *testing.T) {
	installation_path := ""

	expected := "app"

	if got := getAppName(installation_path); got != expected {
		t.Errorf("getAppName(installation_path) = %q, want %q", got, expected)
	}
}

func TestIfGetAppNameReturnsProperValueAtCurrentPath(t *testing.T) {
	installation_path := "./nextjs"

	expected := "nextjs"

	if got := getAppName(installation_path); got != expected {
		t.Errorf("getAppName(installation_path) = %q, want %q", got, expected)
	}
}

func TestIfGetAppNameReturnsProperValueAtRelativePath(t *testing.T) {
	installation_path := "../../nextjs"

	expected := "nextjs"

	if got := getAppName(installation_path); got != expected {
		t.Errorf("getAppName(installation_path) = %q, want %q", got, expected)
	}
}

func TestIfGetAppNameReturnsProperValueAtDefaultValue(t *testing.T) {
	installation_path := "../../app"

	expected := "app"

	if got := getAppName(installation_path); got != expected {
		t.Errorf("getAppName(installation_path) = %q, want %q", got, expected)
	}
}

func TestIfGetAppNameReturnsProperValueAtOnlyRelativePath(t *testing.T) {
	installation_path := "../.."

	expected := "app"

	if got := getAppName(installation_path); got != expected {
		t.Errorf("getAppName(installation_path) = %q, want %q", got, expected)
	}
}

func TestIfGetAppNameReturnsProperValueAtOnlyRelativePathEndingWithSlash(t *testing.T) {
	installation_path := "../../"

	expected := "app"

	if got := getAppName(installation_path); got != expected {
		t.Errorf("getAppName(installation_path) = %q, want %q", got, expected)
	}
}

func TestIfGetAppNameReturnsProperValueAtEmptyInstallationPath(t *testing.T) {
	installation_path := ""

	expected := "app"

	if got := getAppName(installation_path); got != expected {
		t.Errorf("getAppName(installation_path) = %q, want %q", got, expected)
	}
}

func TestIfGetAppNameReturnsProperValueAtInstallationPathOnlyWithSpaces(t *testing.T) {
	installation_path := strings.Repeat(" ", 10)

	expected := "app"

	if got := getAppName(installation_path); got != expected {
		t.Errorf("getAppName(installation_path) = %q, want %q", got, expected)
	}
}
