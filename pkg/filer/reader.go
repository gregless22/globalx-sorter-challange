package filer

import (
	"errors"
	"os"
)

func ReadFile(filename string, container []string) error {
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		return errors.New("error opening the file")
	}
	defer file.Close()

	return nil
}
