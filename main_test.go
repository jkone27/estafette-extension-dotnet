package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSearchAndDeleteFile(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Create a test file named "nuget.config" in the temporary directory
	testFilePath := filepath.Join(tempDir, "nuget.config")
	_, err := os.Create(testFilePath)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer os.Remove(testFilePath)

	// Call the function being tested
	err = searchAndDeleteFile("nuget.config")
	if err != nil {
		t.Fatalf("Error deleting file: %v", err)
	}

	// Check if the file is deleted
	if _, err := os.Stat(testFilePath); !os.IsNotExist(err) {
		t.Errorf("File was not deleted")
	}
}
