package read

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"testing"
)

func TestListDir(t *testing.T) {
	type args struct {
		path string
	}
	dir := t.TempDir()
	// Create test files and directories
	f1, err := os.CreateTemp(dir, "file1")
	if err != nil {
		t.Fatalf("Failed to create file1: %v", err)
	}
	f1.Close()
	f2, err := os.CreateTemp(dir, "file2")
	if err != nil {
		t.Fatalf("Failed to create file2: %v", err)
	}
	f2.Close()
	subdir := filepath.Join(dir, "subdir")
	if err := os.Mkdir(subdir, 0755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}
	nonExistent := filepath.Join(dir, "notfound")

	// Permission denied dir (only for non-Windows and non-root)
	permDeniedDir := filepath.Join(dir, "noread")
	if err := os.Mkdir(permDeniedDir, 0755); err != nil {
		t.Fatalf("Failed to create noread: %v", err)
	}
	skipPermTest := false
	if runtime.GOOS == "windows" {
		skipPermTest = true
	}
	if !skipPermTest {
		curUser, err := user.Current()
		if err == nil && curUser.Uid == "0" {
			skipPermTest = true
		}
	}
	if !skipPermTest {
		if err := os.Chmod(permDeniedDir, 0); err != nil {
			t.Fatalf("Failed to chmod noread: %v", err)
		}
		defer os.Chmod(permDeniedDir, 0755)
	}

	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
		skip    bool
	}{
		{
			name:    "valid directory",
			args:    args{path: dir},
			want:    nil, // Order not guaranteed, so we don't test values
			wantErr: false,
		},
		{
			name:    "file path",
			args:    args{path: f1.Name()},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "non-existent path",
			args:    args{path: nonExistent},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "permission denied directory",
			args:    args{path: permDeniedDir},
			want:    nil,
			wantErr: true,
			skip:    skipPermTest,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skip("Skipping permission denied test on this OS or as root user")
			}
			got, err := ListDir(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListDir() error = %v, wantErr %v", err, tt.wantErr)
			}
			// If no error, check contents for the valid dir test
			if tt.name == "valid directory" && err == nil {
				// Should contain file1, file2, subdir, noread (entries can be unordered)
				expectedEntries := map[string]bool{
					filepath.Base(f1.Name()): true,
					filepath.Base(f2.Name()): true,
					"subdir":                 true,
					"noread":                 true,
				}
				gotMap := map[string]bool{}
				for _, entry := range got {
					gotMap[entry] = true
				}
				for entry := range expectedEntries {
					if !gotMap[entry] {
						t.Errorf("Expected directory listing to contain %q", entry)
					}
				}
			}
		})
	}
}
