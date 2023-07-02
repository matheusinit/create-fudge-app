package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func getAppName(installation_path string) string {
	path_slices := strings.Split(installation_path, "/")

	app_name := path_slices[len(path_slices)-1]

	if app_name == "." || strings.Trim(app_name, " ") == "" {
		return "app"
	}

	return app_name
}

func getPath(installation_path string) string {
	path_slices := strings.Split(installation_path, "/")

	path := strings.Join(path_slices[:len(path_slices)-1], "/")

	if strings.Trim(path, " ") == "" {
		return "."
	}

	return path
}

func getProjectOptions(args []string) (string, string) {
	var installation_path string

	if len(args)%2 == 1 {
		installation_path = ""
	} else {
		installation_path = args[len(args)-1]
	}

	app_name := getAppName(installation_path)

	path_name := getPath(installation_path)

	return app_name, path_name
}

func buildProject(framework, app_name, path_name string) {
	fmt.Println("framework", "->", framework)
	fmt.Println("app_name", "->", app_name)
	fmt.Println("path_name", "->", path_name)

	if framework == "nextjs" {
		// TO-DO

		// [x] ensure installation_path is getting nil when not passed OR get default value (./app)
		// [x] get stdout to print
		// [ ] do a complete front-end boilerplate com react, tailwind, eslint

		installation_path := path_name + "/" + app_name

		fmt.Println("installation_path", "->", installation_path)

		cmd := exec.Command("pnpm", "create", "next-app", "--example", "next-css", installation_path)

		cmd.Dir = path_name

		// STD

		var stdout bytes.Buffer
		var stderr bytes.Buffer

		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()

		if err != nil {
			log.Fatal(colorRed(err.Error() + ":" + stderr.String()))
			return
		}

		return
	}

	if framework == "reactjs" {
		framework = "react"
	}

	cmd := exec.Command("pnpm", "create", "vite", app_name, "--template", framework)

	cmd.Dir = path_name

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}
