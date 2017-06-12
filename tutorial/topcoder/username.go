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
}

func validateUserName(name string) bool {
	length := len(name)

	if length < 1 || length > 50 {
		return false, errors.New(errorUseNameLen)
	}

	
}