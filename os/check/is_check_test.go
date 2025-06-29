package check

import (
	"os"
	"testing"
)

func TestIsPathExists(t *testing.T) {
	// Create a temp directory
	dir := t.TempDir()

	// Test: directory should exist
	if !IsPathExists(dir) {
		t.Errorf("Expected directory %s to exist", dir)
	}

	// Test: non-existent path
	nonExistent := dir + "/notexist"
	if IsPathExists(nonExistent) {
		t.Errorf("Expected path %s to NOT exist", nonExistent)
	}

	// Create a temp file
	file, err := os.CreateTemp(dir, "file")
	if err != nil {
		t.Fatalf("Unable to create temp file: %v", err)
	}
	file.Close()

	// Test: file path (should be false)
	if IsPathExists(file.Name()) {
		t.Errorf("Expected file %s to NOT be recognized as directory", file.Name())
	}
}

func TestIsFileExists(t *testing.T) {
	// Create a temp directory
	dir := t.TempDir()

	// Test: directory should not be a file
	if IsFileExists(dir) {
		t.Errorf("Expected directory %s to NOT be recognized as a file", dir)
	}

	// Test: non-existent path
	nonExistent := dir + "/notexist"
	if IsFileExists(nonExistent) {
		t.Errorf("Expected non-existent path %s to NOT be recognized as a file", nonExistent)
	}

	// Create a temp file
	file, err := os.CreateTemp(dir, "file")
	if err != nil {
		t.Fatalf("Unable to create temp file: %v", err)
	}
	file.Close()

	// Test: file should be recognized as file
	if !IsFileExists(file.Name()) {
		t.Errorf("Expected file %s to exist and be recognized as a file", file.Name())
	}
}
