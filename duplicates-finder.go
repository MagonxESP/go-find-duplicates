package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type DuplicatesFinder struct {
	directoryA string
	directoryB string
	duplicates []string
}

func (d *DuplicatesFinder) find() error {
	filesA, err := ScanDir(d.directoryA)

	if err != nil {
		return err
	}

	err = filepath.WalkDir(d.directoryB, func(path string, fileB fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		for _, file := range filesA {
			if !fileB.IsDir() && fileB.Name() == file.Name() {
				fmt.Println(fmt.Sprintf("%s duplicated on %s", file.Name(), path))
				d.duplicates = append(d.duplicates, path)
			}
		}

		return nil
	})

	return err
}

func NewDuplicatesFinder(directoryA string, directoryB string) (*DuplicatesFinder, error) {
	for _, path := range []string{directoryA, directoryB} {
		err := IsValidPath(path)

		if err != nil {
			return &DuplicatesFinder{}, err
		}
	}

	instance := DuplicatesFinder{}
	instance.directoryA = directoryA
	instance.directoryB = directoryB

	return &instance, nil
}

func ScanDir(directory string) ([]os.DirEntry, error) {
	if err := IsValidPath(directory); err != nil {
		return []os.DirEntry{}, err
	}

	files, err := os.ReadDir(directory)

	if err != nil {
		return []os.DirEntry{}, err
	}

	return files, nil
}

func IsValidPath(path string) error {
	info, err := os.Stat(path)

	if err != nil {
		return err
	}

	if !info.IsDir() {
		return errors.New(fmt.Sprintf("The path %s is not a directory", path))
	}

	return nil
}
