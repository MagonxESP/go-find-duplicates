package main

import (
	"errors"
	"fmt"
	"os"
)

type DuplicatesFinder struct {
	directoryA []os.DirEntry
	directoryB []os.DirEntry
	duplicates []os.File
}

func (d DuplicatesFinder) find() {
	for _, file := range d.directoryA {

	}
}

func NewDuplicatesFinder(directoryA string, directoryB string) (DuplicatesFinder, error) {
	filesA, err := Read(directoryA)

	if err != nil {
		return DuplicatesFinder{}, err
	}

	filesB, err := Read(directoryB)

	if err != nil {
		return DuplicatesFinder{}, err
	}

	instance := DuplicatesFinder{}
	instance.directoryA = filesA
	instance.directoryB = filesB

	return instance, nil
}

func Read(directory string) ([]os.DirEntry, error) {
	err := IsValid(directory)

	if err != nil {
		return []os.DirEntry{}, err
	}

	files, err := os.ReadDir(directory)

	if err != nil {
		return []os.DirEntry{}, err
	}

	return files, nil
}

func IsValid(directory string) error {
	info, err := os.Stat(directory)

	if err != nil {
		return err
	}

	if !info.IsDir() {
		return errors.New(fmt.Sprintf("The path %s is not a directory", directory))
	}

	return nil
}