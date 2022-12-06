package util

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

func ReadInput(filesystem fs.FS, filename string) (string, error) {
	file, err := filesystem.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	return string(content), nil
}
