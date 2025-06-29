package write

import (
	"encoding/json"
	"fmt"
	ch "github.com/htanivar/goart/os/check"
	"os"
	"path/filepath"
)

// WriteFile creates or truncates a file at fileName.
// Returns error if the path does not exist or the file can't be written.
func WriteFile(fileName string) error {
	// 1. Make sure the path must exist (the parent directory)
	dir := filepath.Dir(fileName)
	if !ch.IsPathExists(dir) {
		return fmt.Errorf("directory does not exist: %s", dir)
	}

	// 2. Try to create or truncate the file
	f, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	defer f.Close()
	return nil
}

// WriteObjAsJson marshals modelIn to JSON and writes it to filename.
// Returns error if the parent directory doesn't exist or if writing fails.
func WriteObjAsJson(filename string, modelIn any) error {
	// 1. Make sure the parent directory exists
	dir := filepath.Dir(filename)
	if !ch.IsPathExists(dir) {
		return fmt.Errorf("directory does not exist: %s", dir)
	}

	// 2. Marshal the object to JSON
	data, err := json.MarshalIndent(modelIn, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// 3. Write JSON to file
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
