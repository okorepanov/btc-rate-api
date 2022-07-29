package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isFileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}
