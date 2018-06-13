package main

import "os"

func GetEnv(args ...string) string {
	argsLen := len(args)
	if argsLen == 0 {
		return ""
	}
	if value, ok := os.LookupEnv(args[0]); ok {
		return value
	}
	if argsLen > 1 {
		return args[1]
	}
	return ""
}
