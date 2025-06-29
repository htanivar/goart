package read

import (
	"fmt"
	"os"
)

// ListDir lists all files and directories inside the given path.
// Returns an error if the path does not exist, is not a directory, or on any read error.
func ListDir(path string) ([]string, error) {
	dl := []string{}
	stat, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("cannot access path: %w", err)
	}
	if !stat.IsDir() {
		return nil, fmt.Errorf("path is not a directory: %s", path)
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}
	for _, entry := range entries {
		dl = append(dl, entry.Name())
	}
	return dl, nil
}
