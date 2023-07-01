package main

import (
	"fmt"
	"os"
	"strings"
)

func getProjectName(installation_path string) string {
	path_slices := strings.Split(installation_path, "/")

	app_name := path_slices[len(path_slices)-1]

	if app_name == "." {
		return "app"
	}

	return app_name
}

func getAppName(installation_path string) string {
	path_slices := strings.Split(installation_path, "/")

	path := strings.Join(path_slices[:len(path_slices)-1], "/")

	if strings.Trim(path, " ") == "" {
		return "."
	}

	return path
}

func main() {
	args := os.Args[1:]

	err := checkArgs(args)

	if err != nil {
		fmt.Println(colorRed(err.Error()))
		fmt.Println()
	}

	if len(args) == 0 {
		fmt.Print(
			"\ncreate-fudge-app is a project builder for Typescript ecosystem.\n\n" +
				"Usage:\n\n" +
				"\tcreate-fudge-app <command> [arguments]\n\n\n" +
				"Use \"create-fudge-app help\" for more information.\n",
		)

		return
	}

	if args[0] == "--help" {
		fmt.Print("fudge... you are by yourself in this. Sorry mate")
	}

	argsMap := make(map[string]string)

	for index, arg := range args {
		if index%2 != 0 {
			continue
		}

		argsMap[arg] = args[index+1]
	}

	for key, item := range argsMap {
		fmt.Println(key, "->", item)
	}

	args_length := len(args)

	if args_length%2 != 1 {
		fmt.Println(colorRed("app name not informed"))
		return
	}

	installation_path := args[len(args)-1]

	path := getAppName(installation_path)

	app_name := getProjectName(installation_path)

	buildFramework(argsMap["--framework"], app_name, path)
}
