package main

import (
	"errors"
)

const (
	errorUseNameLen = "User name is incorrect";
)

func main() {
}

func newMember(existingNames[] string, newName string) string  {
	validateUserName(newName);

	return 0;
}

func validateUserName(name string) bool {
	length := len(name)

	if length < 1 || length > 50 {
		return false, errors.New(errorUseNameLen)
	}

	return true
}