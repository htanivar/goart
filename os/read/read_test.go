package read

import (
	"os"
	"path/filepath"
	"testing"
)

type sampleStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestReadFileAsByte(t *testing.T) {
	// Create a temp file with known content
	dir := t.TempDir()
	filepath := dir + "/testfile.txt"
	content := []byte("hello world")
	if err := os.WriteFile(filepath, content, 0644); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	tests := []struct {
		name      string
		filename  string
		want      []byte
		wantError bool
	}{
		{
			name:      "existing file",
			filename:  filepath,
			want:      content,
			wantError: false,
		},
		{
			name:      "non-existent file",
			filename:  dir + "/doesnotexist.txt",
			want:      nil,
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileAsByte(tt.filename)
			if (err != nil) != tt.wantError {
				t.Errorf("ReadFileAsByte() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if err == nil && string(got) != string(tt.want) {
				t.Errorf("ReadFileAsByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFileAsString(t *testing.T) {
	dir := t.TempDir()
	filepath := dir + "/testfile.txt"
	content := "hello gopher"
	if err := os.WriteFile(filepath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	tests := []struct {
		name      string
		filename  string
		want      string
		wantError bool
	}{
		{
			name:      "existing file",
			filename:  filepath,
			want:      content,
			wantError: false,
		},
		{
			name:      "non-existent file",
			filename:  dir + "/doesnotexist.txt",
			want:      "",
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileAsString(tt.filename)
			if (err != nil) != tt.wantError {
				t.Errorf("ReadFileAsString() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if got != tt.want {
				t.Errorf("ReadFileAsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFileAsObj(t *testing.T) {
	dir := t.TempDir()

	// Valid JSON file
	validPath := filepath.Join(dir, "valid.json")
	validContent := `{"name":"Alice","age":30}`
	if err := os.WriteFile(validPath, []byte(validContent), 0644); err != nil {
		t.Fatalf("Failed to write valid JSON file: %v", err)
	}

	// Invalid JSON file
	invalidPath := filepath.Join(dir, "invalid.json")
	invalidContent := `{"name": "Bob", "age":}` // Invalid JSON
	if err := os.WriteFile(invalidPath, []byte(invalidContent), 0644); err != nil {
		t.Fatalf("Failed to write invalid JSON file: %v", err)
	}

	// Non-existent file
	nonexistentPath := filepath.Join(dir, "no.json")

	tests := []struct {
		name     string
		file     string
		wantErr  bool
		wantName string
		wantAge  int
	}{
		{
			name:     "valid JSON file",
			file:     validPath,
			wantErr:  false,
			wantName: "Alice",
			wantAge:  30,
		},
		{
			name:    "invalid JSON file",
			file:    invalidPath,
			wantErr: true,
		},
		{
			name:    "file does not exist",
			file:    nonexistentPath,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result sampleStruct
			err := ReadFileAsObj(tt.file, &result)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileAsObj() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if result.Name != tt.wantName || result.Age != tt.wantAge {
					t.Errorf("Unmarshaled struct = %+v, want Name=%q Age=%d", result, tt.wantName, tt.wantAge)
				}
			}
		})
	}
}
