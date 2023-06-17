package main

import (
	"fmt"
	"os"
	"strings"
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
			"create-fudge-app is a project builder for Typescript ecosystem.\n\n" +
				"Usage:\n\n" +
				"\tcreate-fudge-app <command> [arguments]\n\n\n" +
				"Use \"create-fudge-app help\" for more information.\n\n",
		)

		return
	}

	if args[0] == "--help" {
		fmt.Print("fudge... you are by yourself in this. Sorry mate")
	}

	if args[0] == "--framework" {
		raw_path := "../../my-fudge-app"

		path_slices := strings.Split(raw_path, "/")

		path := strings.Join(path_slices[:len(path_slices)-1], "/")

		app_name := path_slices[len(path_slices)-1]

		buildFramework(args[1], app_name, path)
	}
}
