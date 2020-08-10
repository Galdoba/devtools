package cli

import (
	"fmt"
	"os"
	"strings"
)

func ArgsAll() []string {
	args := os.Args
	fmt.Println(args)
	return args
}

func ArgsFlagged() []string {
	allArgs := os.Args
	fmt.Println(allArgs)
	var fa []string
	for i := range allArgs {
		if string(allArgs[i][0]) == "-" {
			fa = append(fa, allArgs[i])
		}
	}
	return fa
}

func ArgsMap() map[string][]string {
	args := os.Args
	argMap := make(map[string][]string)
	var mapKey string
	var mapSlice []string
	for i, val := range args {
		val = strings.ToLower(val)
		if string(val[0]) == "-" {
			argMap[mapKey] = mapSlice
			mapKey = args[i]
			mapSlice = nil
			continue
		}
		mapSlice = append(mapSlice, args[i])
	}
	argMap[mapKey] = mapSlice
	delete(argMap, "")
	return argMap
}

func ArgExist(arg string) bool {
	args := os.Args
	for i := range args {
		if arg == args[i] {
			return true
		}
	}
	return false
}
