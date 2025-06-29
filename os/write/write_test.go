package write

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type sampleOut struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestWriteFile(t *testing.T) {
	tmpDir := t.TempDir()
	validFile := filepath.Join(tmpDir, "testfile.txt")
	invalidFile := filepath.Join(tmpDir, "nonexistentdir", "file.txt")

	tests := []struct {
		name     string
		fileName string
		wantErr  bool
	}{
		{
			name:     "valid path",
			fileName: validFile,
			wantErr:  false,
		},
		{
			name:     "non-existent directory",
			fileName: invalidFile,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteFile(tt.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			// For the valid case, verify that the file was actually created
			if !tt.wantErr {
				if _, err := os.Stat(tt.fileName); os.IsNotExist(err) {
					t.Errorf("Expected file %s to be created, but it does not exist", tt.fileName)
				}
			}
		})
	}
}

func TestWriteObjAsJson(t *testing.T) {
	dir := t.TempDir()
	validPath := filepath.Join(dir, "out.json")
	invalidPath := filepath.Join(dir, "doesnotexist", "fail.json")

	tests := []struct {
		name      string
		file      string
		input     sampleOut
		wantErr   bool
		checkFile bool
	}{
		{
			name:      "valid write",
			file:      validPath,
			input:     sampleOut{Name: "Ravi", Age: 42},
			wantErr:   false,
			checkFile: true,
		},
		{
			name:      "non-existent directory",
			file:      invalidPath,
			input:     sampleOut{Name: "Jagan", Age: 30},
			wantErr:   true,
			checkFile: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WriteObjAsJson(tt.file, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteObjAsJson() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.checkFile {
				// Read the file and check its contents
				data, err := os.ReadFile(tt.file)
				if err != nil {
					t.Fatalf("Expected file %s to be created, but failed to read: %v", tt.file, err)
				}
				var out sampleOut
				if err := json.Unmarshal(data, &out); err != nil {
					t.Fatalf("Failed to unmarshal written file: %v", err)
				}
				if out != tt.input {
					t.Errorf("File contents = %+v, want %+v", out, tt.input)
				}
			}
		})
	}
}
