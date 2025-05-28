package io

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCsvFileReader_ReadEntriesOfFiles(t *testing.T) {
	t.Run("read single CSV file", func(t *testing.T) {
		// Create temporary CSV file
		tempDir := t.TempDir()
		csvFile := filepath.Join(tempDir, "test.csv")

		csvContent := `name,age,city
John,25,New York
Jane,30,Los Angeles
Bob,35,Chicago`

		err := os.WriteFile(csvFile, []byte(csvContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		reader := &CsvFileReader{}
		result := reader.ReadEntriesOfFiles([]string{csvFile})

		expected := [][]string{
			{"John", "25", "New York"},
			{"Jane", "30", "Los Angeles"},
			{"Bob", "35", "Chicago"},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("read multiple CSV files", func(t *testing.T) {
		tempDir := t.TempDir()

		// Create first CSV file
		csvFile1 := filepath.Join(tempDir, "test1.csv")
		csvContent1 := `name,age
Alice,28
Charlie,32`
		err := os.WriteFile(csvFile1, []byte(csvContent1), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file 1: %v", err)
		}

		// Create second CSV file
		csvFile2 := filepath.Join(tempDir, "test2.csv")
		csvContent2 := `name,age
David,45
Eve,38`
		err = os.WriteFile(csvFile2, []byte(csvContent2), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file 2: %v", err)
		}

		reader := &CsvFileReader{}
		result := reader.ReadEntriesOfFiles([]string{csvFile1, csvFile2})

		// Since files are read concurrently, we need to check that all expected entries are present
		expectedEntries := map[string]bool{
			"Alice,28":   false,
			"Charlie,32": false,
			"David,45":   false,
			"Eve,38":     false,
		}

		if len(result) != 4 {
			t.Errorf("Expected 4 entries, got %d", len(result))
		}

		for _, entry := range result {
			if len(entry) == 2 {
				key := entry[0] + "," + entry[1]
				if _, exists := expectedEntries[key]; exists {
					expectedEntries[key] = true
				}
			}
		}

		for key, found := range expectedEntries {
			if !found {
				t.Errorf("Expected entry %s not found in result", key)
			}
		}
	})

	t.Run("read from non-existent file", func(t *testing.T) {
		reader := &CsvFileReader{}
		result := reader.ReadEntriesOfFiles([]string{"non-existent-file.csv"})

		// Should return empty slice when file doesn't exist
		if len(result) != 0 {
			t.Errorf("Expected empty result for non-existent file, got %v", result)
		}
	})

	t.Run("read empty CSV file with header only", func(t *testing.T) {
		tempDir := t.TempDir()
		csvFile := filepath.Join(tempDir, "empty.csv")

		csvContent := `name,age,city`
		err := os.WriteFile(csvFile, []byte(csvContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		reader := &CsvFileReader{}
		result := reader.ReadEntriesOfFiles([]string{csvFile})

		if len(result) != 0 {
			t.Errorf("Expected empty result for file with header only, got %v", result)
		}
	})

	t.Run("read CSV with different column counts", func(t *testing.T) {
		tempDir := t.TempDir()
		csvFile := filepath.Join(tempDir, "irregular.csv")

		// CSV with inconsistent column counts will cause an error
		// The readFile function will return partial results when it encounters the error
		csvContent := `name,age,city
John,25,New York
Jane,30
Bob,35,Chicago,Extra`

		err := os.WriteFile(csvFile, []byte(csvContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		reader := &CsvFileReader{}
		result := reader.ReadEntriesOfFiles([]string{csvFile})

		// The CSV reader will stop at the first line with wrong field count
		// So we should only get the first valid line
		expected := [][]string{
			{"John", "25", "New York"},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("read CSV with quoted fields", func(t *testing.T) {
		tempDir := t.TempDir()
		csvFile := filepath.Join(tempDir, "quoted.csv")

		csvContent := `name,description,value
"John Doe","Software Engineer, Senior",100000
"Jane Smith","Manager, Team Lead",120000`

		err := os.WriteFile(csvFile, []byte(csvContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		reader := &CsvFileReader{}
		result := reader.ReadEntriesOfFiles([]string{csvFile})

		expected := [][]string{
			{"John Doe", "Software Engineer, Senior", "100000"},
			{"Jane Smith", "Manager, Team Lead", "120000"},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("read empty file list", func(t *testing.T) {
		reader := &CsvFileReader{}
		result := reader.ReadEntriesOfFiles([]string{})

		if len(result) != 0 {
			t.Errorf("Expected empty result for empty file list, got %v", result)
		}
	})

	t.Run("read CSV with special characters", func(t *testing.T) {
		tempDir := t.TempDir()
		csvFile := filepath.Join(tempDir, "special.csv")

		csvContent := `name,email,notes
José García,jose@email.com,"Contains áéíóú"
山田太郎,yamada@email.com,"日本語"`

		err := os.WriteFile(csvFile, []byte(csvContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		reader := &CsvFileReader{}
		result := reader.ReadEntriesOfFiles([]string{csvFile})

		expected := [][]string{
			{"José García", "jose@email.com", "Contains áéíóú"},
			{"山田太郎", "yamada@email.com", "日本語"},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestReadFile(t *testing.T) {
	t.Run("read valid CSV file", func(t *testing.T) {
		tempDir := t.TempDir()
		csvFile := filepath.Join(tempDir, "test.csv")

		csvContent := `header1,header2
value1,value2
value3,value4`

		err := os.WriteFile(csvFile, []byte(csvContent), 0644)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}

		ch := make(chan [][]string, 1)
		readFile(csvFile, ch)
		result := <-ch

		expected := [][]string{
			{"value1", "value2"},
			{"value3", "value4"},
		}

		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("read non-existent file", func(t *testing.T) {
		ch := make(chan [][]string, 1)
		readFile("non-existent-file.csv", ch)
		result := <-ch

		if len(result) != 0 {
			t.Errorf("Expected empty result for non-existent file, got %v", result)
		}
	})

	t.Run("read file with permission error", func(t *testing.T) {
		// This test might not work on all systems, so we'll skip it if we can't create the scenario
		tempDir := t.TempDir()
		csvFile := filepath.Join(tempDir, "restricted.csv")

		err := os.WriteFile(csvFile, []byte("test"), 0000) // No permissions
		if err != nil {
			t.Skip("Cannot create file with restricted permissions")
		}

		ch := make(chan [][]string, 1)
		readFile(csvFile, ch)
		result := <-ch

		// Should return empty slice when file can't be opened
		if len(result) != 0 {
			t.Errorf("Expected empty result for restricted file, got %v", result)
		}

		// Cleanup - restore permissions so temp dir can be cleaned up
		_ = os.Chmod(csvFile, 0644)
	})
}

// Test the FileReader interface compliance
func TestFileReaderInterface(t *testing.T) {
	var _ FileReader = &CsvFileReader{}

	// This test ensures that CsvFileReader implements the FileReader interface
	t.Run("CsvFileReader implements FileReader", func(t *testing.T) {
		reader := &CsvFileReader{}
		if reader == nil {
			t.Error("CsvFileReader should not be nil")
		}
	})
}
