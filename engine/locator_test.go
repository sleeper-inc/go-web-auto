package engine

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func createLocatorFile(t *testing.T, dir, page string, locators map[string]string) {
	t.Helper()
	data, err := json.Marshal(locators)
	if err != nil {
		t.Fatalf("failed to marshal locators: %v", err)
	}
	filePath := filepath.Join(dir, page+".json")
	err = os.WriteFile(filePath, data, 0644)
	if err != nil {
		t.Fatalf("failed to write locator file: %v", err)
	}
}

// Override os.ReadFile temporarily for testing
func TestLocator_Get(t *testing.T) {
	tmpDir := t.TempDir()
	originalWD, _ := os.Getwd()
	defer func(dir string) {
		err := os.Chdir(dir)
		if err != nil {
			t.Fatalf("failed to change directory: %v", err)
		}
	}(originalWD)
	err := os.Chdir(tmpDir)
	if err != nil {
		return
	}

	// Create `locators` subdirectory
	err = os.Mkdir("locators", 0755)
	if err != nil {
		t.Fatalf("failed to create locators directory: %v", err)
	}

	// Prepare test data
	page := "login"
	locators := map[string]string{
		"username": "id=username_field",
		"password": "id=password_field",
	}
	createLocatorFile(t, filepath.Join(tmpDir, "locators"), page, locators)

	loc := NewLocator()

	t.Run("valid key returns locator", func(t *testing.T) {
		got, err := loc.Get("login.username")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != "id=username_field" {
			t.Errorf("expected 'id=username_field', got '%s'", got)
		}
	})

	t.Run("invalid key format returns error", func(t *testing.T) {
		_, err := loc.Get("invalidkey")
		if err == nil || err.Error() != "invalid locator key: invalidkey" {
			t.Errorf("expected invalid key error, got %v", err)
		}
	})

	t.Run("non-existent locator name returns error", func(t *testing.T) {
		_, err := loc.Get("login.notexist")
		if err == nil || err.Error() != "locator not found: login.notexist" {
			t.Errorf("expected locator not found error, got %v", err)
		}
	})

	t.Run("non-existent page file returns error", func(t *testing.T) {
		_, err := loc.Get("signup.username")
		if err == nil {
			t.Errorf("expected file read error, got nil")
		}
	})
}
