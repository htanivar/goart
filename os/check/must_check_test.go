package check

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMustExists(t *testing.T) {
	dir := t.TempDir()

	// Create a file inside temp dir
	file := filepath.Join(dir, "file.txt")
	if err := os.WriteFile(file, []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	nonExistent := filepath.Join(dir, "doesnotexist")

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "directory exists",
			path:    dir,
			wantErr: false,
		},
		{
			name:    "file exists",
			path:    file,
			wantErr: false,
		},
		{
			name:    "path does not exist",
			path:    nonExistent,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustExists(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("MustExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMustNotExists(t *testing.T) {
	dir := t.TempDir()

	// Create a file inside temp dir
	file := filepath.Join(dir, "file.txt")
	if err := os.WriteFile(file, []byte("data"), 0644); err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	nonExistent := filepath.Join(dir, "notexists")

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "directory exists",
			path:    dir,
			wantErr: true,
		},
		{
			name:    "file exists",
			path:    file,
			wantErr: true,
		},
		{
			name:    "path does not exist",
			path:    nonExistent,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MustNotExists(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("MustNotExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
