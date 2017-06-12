package main

import (
	"errors"
)

const (
	userNameInvalidLenErrorMessage = "User name length is invalid"
	userNameTakenErrorMsg = "User name has been already takle"
)

func main() {
}

func newMember(existingNames[] string, newName string) (error, string) {

	if validateLen(newName) == false {
		return 0, errors.New(userNameInvalidLenErrorMessage)
	}

	return 0;
}

func validateLen(name string) bool {
	length := len(name)
	if length < 1 || length > 50 {
		return false
	}
	return true
}

func validate