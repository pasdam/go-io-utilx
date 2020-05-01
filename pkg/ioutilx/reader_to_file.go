package ioutilx

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// ReaderToFile writes the content of Reader to the specified destination file.
// The folders structure will be created if it does not exist.
func ReaderToFile(reader io.Reader, dst string) error {
	err := os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	if err != nil {
		log.Printf("Error while creating parents folder of %s\n", dst)
		return err
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, reader)
	return err
}
