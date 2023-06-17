package main

import "errors"

func isANoValueArg(arg string) bool {
	noValueArgs := []string{"--help", "--version", "-v", "-h"}

	for _, argItem := range noValueArgs {
		if argItem == arg {
			return true
		}
	}

	return false
}

func checkArgs(args []string) error {
	for index, arg := range args {
		if isANoValue := isANoValueArg(arg); isANoValue {
			continue
		}

		if arg[0] == '-' && len(args)-1 == index {
			return errors.New("no value was passed for arg " + "\"" + arg + "\"")
		}

		if arg[0] == '-' && arg[1] != '-' {
			if len(arg) > 2 {
				return errors.New("grouped shorthand is not permitted")
			}

			if nextArg := args[index+1]; nextArg[0] == '-' {
				return errors.New("no value was passed for arg " + "\"" + arg + "\"")
			}
		}

		if arg[0] == '-' && arg[1] == '-' {
			nextArg := args[index+1]

			if nextArg == "" || nextArg[0] == '-' {
				return errors.New("no value was passed for arg " + "\"" + arg + "\"")
			}
		}
	}

	return nil
}
