package os

import (
	"os"
)

// IsPathExists checks if the given path exists and is a directory.
// Returns true if the path exists and is a directory, otherwise false.
func IsPathExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		// Path does not exist or error in stat
		return false
	}
	return stat.IsDir()
}

// IsFileExists checks if the given path exists and is a regular file.
// Returns true if the path exists and is a file, otherwise false.
func IsFileExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		// Path does not exist or error in stat
		return false
	}
	return stat.Mode().IsRegular()
}
