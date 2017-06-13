package main

import (
	"errors"
	"regexp"
	"fmt"
)

const (
	userNameInvalidLenErrorMessage = "User name length is invalid"
	userNameInvalidContentErrorMessage = "User name contains illigal symbols"
)

func main() {

	printNewName("test")

	printNewName("test2")

	printNewName("")

	var shortName string
	for i := 0; i < 50; i++ {
		shortName += "a";
	}
	printNewName(shortName)

	var longName string
	for i := 0; i < 51; i++ {
		longName += "a";
	}
	printNewName(longName)

	buildDB([]string{"MasterOfDisaster", "DingBat", "Orpheus", "WolfMan", "MrKnowItAll"})
}

func printNewName(newName string) {
	r, err := newMember(newName)
	if err != nil {
		fmt.Printf("\"%s\" is not available: \"%s\"\n", newName, err)
		return
	}
	fmt.Printf("\"%s\" is available\n", r)
}

func newMember(newName string) (string, error) {

	length := len(newName)
	if length < 1 || length > 50 {
		return "", errors.New(userNameInvalidLenErrorMessage)
	}

	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(newName) {
		return "", errors.New(userNameInvalidContentErrorMessage)
	}

	return newName, nil;
}

func buildDB(existingNames []string) map[string]int {

	db := make(map[string]int)

	for _, name := range existingNames {
		counter, ok := db[name]
		if (!ok) {
			counter = 0
		}
		db[name] = counter + 1
	}

	fmt.Println(db)

	return db
}