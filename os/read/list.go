package read

import (
	"fmt"
	"os"
)

// ListDir lists all files and directories inside the given path.
// Returns an error if the path does not exist, is not a directory, or on any read error.
func ListDir(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("cannot access path: %w", err)
	}
	if !stat.IsDir() {
		return fmt.Errorf("path is not a directory: %s", path)
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}
	return nil
}
