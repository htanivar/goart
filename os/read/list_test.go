package read

import (
	"os"
	"os/user"
	"runtime"
	"testing"
)

func TestListDir(t *testing.T) {
	type args struct {
		path string
	}
	dir := t.TempDir()
	file1, err := os.CreateTemp(dir, "file1")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	file1.Close()
	subdir := dir + "/subdir"
	if err := os.Mkdir(subdir, 0755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}
	nonExistent := dir + "/doesnotexist"

	noread := dir + "/noreaddir"
	if err := os.Mkdir(noread, 0755); err != nil {
		t.Fatalf("Failed to create noreaddir: %v", err)
	}

	// Remove all permissions for noread directory (Unix only)
	skipPermissionCase := false
	if runtime.GOOS == "windows" {
		skipPermissionCase = true
	}
	if !skipPermissionCase {
		currentUser, err := user.Current()
		if err == nil && currentUser.Uid == "0" {
			// Running as root, skip permission test
			skipPermissionCase = true
		}
	}
	if !skipPermissionCase {
		if err := os.Chmod(noread, 0); err != nil {
			t.Fatalf("Failed to remove permissions from noreaddir: %v", err)
		}
		defer os.Chmod(noread, 0755) // Restore permissions after test
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		skip    bool
	}{
		{
			name:    "valid directory",
			args:    args{path: dir},
			wantErr: false,
		},
		{
			name:    "file path",
			args:    args{path: file1.Name()},
			wantErr: true,
		},
		{
			name:    "non-existent path",
			args:    args{path: nonExistent},
			wantErr: true,
		},
		{
			name:    "permission denied",
			args:    args{path: noread},
			wantErr: true,
			skip:    skipPermissionCase,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			if tt.skip {
				t.Skipf("Skipping %q test due to OS or root user", tt.name)
			}
			if err := ListDir(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("ListDir() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
