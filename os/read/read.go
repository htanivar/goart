package read

import (
	"encoding/json"
	"fmt"
	ch "github.com/htanivar/goart/os/check"
	"os"
)

// ReadFileAsByte reads the contents of a file and returns it as a byte slice.
// Returns an empty slice if the file does not exist or cannot be read.
func ReadFileAsByte(filename string) ([]byte, error) {
	// 1. Make sure the file must exist
	if !ch.IsFileExists(filename) {
		return nil, fmt.Errorf("file does not exist: %s", filename)
	}

	// 2. Read the file as []byte
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}
	return data, nil
}

// ReadFileAsString reads the contents of a file and returns it as a string.
// Returns an empty string if the file does not exist or cannot be read.
func ReadFileAsString(filename string) (string, error) {
	// 1. Use ReadFileAsByte
	data, err := ReadFileAsByte(filename)
	if err != nil {
		return "", err
	}
	// 2. Return the string
	return string(data), nil
}

// ReadFileAsObj reads the contents of a JSON file and unmarshals it into modelIn (should be a pointer).
// Returns an error if the file does not exist, can't be read, or the JSON is invalid.
func ReadFileAsObj(filename string, modelIn any) error {
	// 1. Check if the file exists
	if !ch.IsFileExists(filename) {
		return fmt.Errorf("file does not exist: %s", filename)
	}

	// 2. Read file contents as bytes
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// 3. Unmarshal JSON data into the provided model
	if err := json.Unmarshal(data, modelIn); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}
