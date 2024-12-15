package generating

import (
	"os"
	"testing"
)

func TestCreateDirIfNotExists(t *testing.T) {
	if err := createDirIfNotExists(prefix); err != nil {
		t.Errorf("error while creating directory: %v", err)
	}

	if _, err := os.Stat(prefix); os.IsNotExist(err) {
		t.Errorf("Directory %s must be exists, but it is not: %v", prefix, err)
	}
	os.Remove(prefix)
}

func TestClearGeneratedFiles(t *testing.T) {
	if err := createDirIfNotExists(prefix); err != nil {
		t.Fatalf("error while creating directory: %v", err)
	}

	if err := GenerateFile("test.txt", 100); err != nil {
		t.Fatalf("error while generating file: %v", err)
	}

	if err := ClearGeneratedFiles(); err != nil {
		t.Fatalf("error while clearing generated files: %v", err)
	}
}

func TestFileCreatingAndReading(t *testing.T) {
	var filename string = "test.txt"
	var amount int = 100
	if err := GenerateFile(filename, amount); err != nil {
		t.Fatalf("error while generating: %v", err)
	}
	nums, err := ReadDataFromFile(filename)
	if err != nil {
		t.Fatalf("error while reading: %v", err)
	}

	if len(nums) != amount {
		t.Fatalf("amount of written nums must be equals with reading ones")
	}

	if err := ClearGeneratedFiles(); err != nil {
		t.Errorf("error while clearing generated files: %v", err)
	}
}
