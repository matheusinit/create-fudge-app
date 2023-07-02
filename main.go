package main

import (
	"fmt"
	"os"
)

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
		fmt.Print("fudge... you are by yourself in this. Sorry mate.")
	}

	argsMap := make(map[string]string)

	for index, arg := range args {
		if index%2 != 0 {
			continue
		}

		if index+1 >= len(args) {
			break
		}

		argsMap[arg] = args[index+1]
	}

	for key, item := range argsMap {
		fmt.Println(key, "->", item)
	}

	app_name, path_name := getProjectOptions(args)

	buildProject(argsMap["--fe-framework"], app_name, path_name)
}
