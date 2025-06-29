package check

import (
	"fmt"
	"os"
)

// MustExists checks if the given file or directory path exists.
// Returns true if the path exists (file or directory), otherwise false.
func MustExists(fileOrPath string) error {
	_, err := os.Stat(fileOrPath)
	if err != nil {
		return err
	}
	return nil
}

// MustNotExists checks if the given file or directory path does NOT exist.
// Returns nil if the path does not exist, otherwise returns an error.
func MustNotExists(fileOrPath string) error {
	_, err := os.Stat(fileOrPath)
	if err == nil {
		return fmt.Errorf("path exists: %s", fileOrPath)
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}
