package main

import (
	"fmt"
	"os"
	"time"
)

func findDuplicates(dirA string, dirB string) error {
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
	start := time.Now()
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Missing required arguments <directoryA> <directoryB>")
		return
	}

	dirA := args[0]
	dirB := args[1]

	err := findDuplicates(dirA, dirB)

	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	elapsed := time.Since(start)
	fmt.Println(fmt.Sprintf("Execution time: %s", elapsed))
}
