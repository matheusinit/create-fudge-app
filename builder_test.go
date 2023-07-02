package main

import "testing"

func TestIfGetAppNameDefaultValue(t *testing.T) {
	installation_path := ""
	app_name_informed := false

	expected := "app"

	if got := getAppName(installation_path, app_name_informed); got != expected {
		t.Errorf("getAppName(installation_path, app_name_informed) = %q, want %q", got, expected)
	}
}
