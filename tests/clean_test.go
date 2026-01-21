package tests

import (
	"runtime"
	"testing"

	"github.com/igorarthur/macleaner/cmd"
	"github.com/igorarthur/macleaner/internal/paths"
)

func TestClean(t *testing.T) {
	testPath := "/var/lib/docker"
	if p, ok := paths.DockerPaths[runtime.GOOS]; ok && len(p) > 0 {
		testPath = p[0]
	}

	tests := []struct {
		name          string
		existingFiles []string
		isDryRun      bool
		isAssumeYes   bool
		expectRemoved int
	}{
		{
			name:          "Normal deletion with confirmation skipped",
			existingFiles: []string{testPath},
			isDryRun:      false,
			isAssumeYes:   true,
			expectRemoved: 1,
		},
		{
			name:          "Dry run should not remove anything",
			existingFiles: []string{testPath},
			isDryRun:      true,
			isAssumeYes:   true,
			expectRemoved: 0,
		},
		{
			name:          "No files exist",
			existingFiles: []string{},
			isDryRun:      false,
			isAssumeYes:   true,
			expectRemoved: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := &MockFS{existingPaths: tt.existingFiles}
			dryRun := tt.isDryRun
			assumeYes := tt.isAssumeYes

			err := cmd.Clean(mockFS, dryRun, assumeYes)

			if err != nil {
				t.Errorf("clean() produced unexpected error: %v", err)
			}
			if len(mockFS.removedPaths) != tt.expectRemoved {
				t.Errorf("expected %d removals, got %d", tt.expectRemoved, len(mockFS.removedPaths))
			}
		})
	}
}
