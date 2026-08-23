package filehandling_test

import (
	"os"
	"path/filepath"
	"testing"

	filehandling "github.com/vinayakgaud/schemabridge/core/filehandling"
)

func TestReadFile(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "sample.json")
	expected := []byte(`{"name":"SchemaBridge"}`)

	if err := os.WriteFile(filePath, expected, 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	actual, err := filehandling.ReadFile(filePath)

	if err != nil {
		t.Fatalf("ReadFile() returned an error: %v", err)
	}

	if string(actual) != string(expected) {
		t.Fatalf("ReadFile() = %q, want %q", actual, expected)
	}
}

func TestReadFile_FileNotFound(t *testing.T) {
	tempDir := t.TempDir()
	filePath := filepath.Join(tempDir, "does-not-exist.json")
	_, err := filehandling.ReadFile(filePath)

	if err == nil {
		t.Fatal("ReadFile() expected an error, got nil")
	}
}
