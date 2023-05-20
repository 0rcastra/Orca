package utils

import (
	"fmt"
	"io"
	"os"
)

func CaptureOutput(fn func()) (string, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return "", fmt.Errorf("failed to create pipe: %w", err)
	}

	oldStdout := os.Stdout
	os.Stdout = w

	fn()

	os.Stdout = oldStdout
	if err := w.Close(); err != nil {
		return "", fmt.Errorf("failed to close pipe: %w", err)
	}

	outputBytes, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("failed to read captured output: %w", err)
	}

	return string(outputBytes), nil
}
