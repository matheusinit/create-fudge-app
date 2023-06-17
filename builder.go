package main

import (
	"log"
	"os/exec"
)

func buildFramework(framework, projectName, path string) {
	if framework == "react" {
		cmd := exec.Command("pnpm", "create", "vite", projectName, "--template", framework)

		cmd.Dir = path

		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}
	}
}
