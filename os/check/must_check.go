package check

import "os"

// MustExists checks if the given file or directory path exists.
// Returns true if the path exists (file or directory), otherwise false.
func MustExists(fileOrPath string) bool {
	_, err := os.Stat(fileOrPath)
	return err == nil
}
