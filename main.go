package main

import (
	"fmt"
	"os"
)

func findDuplicates() error {
	dirA := os.Args[1:][0]
	dirB := os.Args[1:][1]

	finder, err := NewDuplicatesFinder(dirA, dirB)

	if err != nil {
		return err
	}

	if err = finder.find(); err != nil {
		return err
	}

	count := len(finder.duplicates)

	fmt.Println(fmt.Sprintf("Found %d files of %s duplicated on directory %s and subdirectories", count, dirA, dirB))
	return nil
}

func main() {
	err := findDuplicates()

	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
	}
}
