package tests

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/igorarthur/macleaner/cmd"
	"github.com/igorarthur/macleaner/internal/paths"
)

func TestScan(t *testing.T) {
	testPath := "/tmp/docker.log"
	if p, ok := paths.DockerPaths[runtime.GOOS]; ok && len(p) > 0 {
		testPath = p[0]
	}

	tests := []struct {
		name           string
		existingFiles  []string
		expectedOutput string
	}{
		{
			name:           "Found one file",
			existingFiles:  []string{testPath},
			expectedOutput: fmt.Sprintf("%s → 0.00 GB", testPath),
		},
		{
			name:           "Found nothing",
			existingFiles:  []string{},
			expectedOutput: "No Docker paths found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockFS := &MockFS{existingPaths: tt.existingFiles}

			buf := new(bytes.Buffer)

			err := cmd.Scan(mockFS, buf)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !strings.Contains(buf.String(), tt.expectedOutput) {
				t.Errorf("expected output %q, but got %q", tt.expectedOutput, buf.String())
			}
		})
	}
}
