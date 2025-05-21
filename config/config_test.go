package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig_Success(t *testing.T) {
	content := []byte("pack_sizes: [250, 500, 1000]\n")
	tmpFile := createTempFile(t, "success", content)
	defer os.Remove(tmpFile)

	cfg, err := NewConfig(tmpFile)
	assert.NoError(t, err)
	assert.Equal(t, []int{250, 500, 1000}, cfg.PackSize)
}

func TestNewConfig_Fail_EmptyPackSize(t *testing.T) {
	tmpFile := createTempFile(t, "empty", []byte("pack_sizes: []\n"))
	defer os.Remove(tmpFile)

	_, err := NewConfig(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "pack_size is empty or not configured")
}

func TestNewConfig_Fail_ReadFile(t *testing.T) {
	tmpFile := "readFileError.yaml"
	content := []byte("pack_sizes: [250, 500, NOEXIST]\n")
	err := os.WriteFile(tmpFile, content, 0644)
	assert.NoError(t, err)
	defer os.Remove(tmpFile)

	_, err = NewConfig(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse config file")
}

func TestNewConfig_Fail_WrongKey(t *testing.T) {
	tmpFile := createTempFile(t, "wrong_key", []byte("wrong_key: [250, 500, 1000]\n"))
	defer os.Remove(tmpFile)

	_, err := NewConfig(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "pack_size is empty or not configured")
}

func TestNewConfig_Fail_EmptyFile(t *testing.T) {
	tmpFile := createTempFile(t, "empty_file", []byte(""))
	defer os.Remove(tmpFile)

	_, err := NewConfig(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "pack_size is empty or not configured")
}

func TestNewConfig_Fail_InvalidYAML(t *testing.T) {
	tmpFile := createTempFile(t, "invalid_yaml", []byte("pack_sizes: [250, 500, 1000\n"))
	defer os.Remove(tmpFile)

	_, err := NewConfig(tmpFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse config file")
}

func createTempFile(t *testing.T, name string, content []byte) string {
	fileName := fmt.Sprintf("test_config_%s.yaml", name)
	if err := os.WriteFile(fileName, content, 0644); err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	return fileName
}
