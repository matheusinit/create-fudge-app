package main

import (
	"log"
	"os/exec"
)

func buildProject(framework, installation_path string, app_name_informed bool) {
	path := getAppName(installation_path)

	app_name := getProjectName(installation_path, app_name_informed)

	if framework == "nextjs" {
		// TO-DO

		// [ ] ensure installation_path is getting nil when not passed OR get default value (./app)

		installation_path = path + "/" + app_name

		cmd := exec.Command("pnpm", "create", "next-app", "--example", "next-css", installation_path)

		cmd.Dir = path

		err := cmd.Run()

		if err != nil {
			log.Fatal(err)
		}

		return
	}

	if framework == "reactjs" {
		framework = "react"
	}

	cmd := exec.Command("pnpm", "create", "vite", app_name, "--template", framework)

	cmd.Dir = path

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
