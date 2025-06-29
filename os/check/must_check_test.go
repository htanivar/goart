package check

import (
	"os"
	"testing"
)

func TestMustExists(t *testing.T) {
	// Create a temp directory
	dir := t.TempDir()

	// Directory should exist
	if !MustExists(dir) {
		t.Errorf("Expected temp directory %s to exist", dir)
	}

	// Non-existent path should not exist
	nonExistent := dir + "/doesnotexist"
	if MustExists(nonExistent) {
		t.Errorf("Expected non-existent path %s to NOT exist", nonExistent)
	}

	// Create a temp file
	file, err := os.CreateTemp(dir, "testfile")
	if err != nil {
		t.Fatalf("Unable to create temp file: %v", err)
	}
	filePath := file.Name()
	file.Close()

	// File should exist
	if !MustExists(filePath) {
		t.Errorf("Expected file %s to exist", filePath)
	}

	// Remove file and check again
	if err := os.Remove(filePath); err != nil {
		t.Fatalf("Unable to remove temp file: %v", err)
	}
	if MustExists(filePath) {
		t.Errorf("Expected removed file %s to NOT exist", filePath)
	}
}
