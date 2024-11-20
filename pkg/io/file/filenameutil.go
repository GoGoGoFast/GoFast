// Package fileutil provides utility functions for handling file names and paths
// across different operating systems. This package ensures that file names are
// valid and cleans invalid characters based on the operating system's rules.
package fileutil

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// GetBaseName returns the base name of the file without the extension.
// For example, given the file path "example/test/file.txt", it will return "file".
//
// Parameters:
// - filePath: the path of the file
//
// Returns:
// - string: the base name of the file
func GetBaseName(filePath string) string {
	return strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
}

// GetExtension returns the file extension.
// For example, given the file path "example/test/file.txt", it will return ".txt".
//
// Parameters:
// - filePath: the path of the file
//
// Returns:
// - string: the file extension
func GetExtension(filePath string) string {
	return filepath.Ext(filePath)
}

// GetAbsolutePath returns the absolute path for a given relative or classpath file.
// It resolves the absolute path based on the current working directory.
// If the operation fails, it returns an error.
//
// Parameters:
// - path: the relative or classpath file
//
// Returns:
// - string: the absolute path
// - error: if an error occurs
func GetAbsolutePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path of %s: %w", path, err)
	}
	return absPath, nil
}

// CleanInvalidChars cleans invalid characters from the file name based on the current operating system.
// For Windows, it replaces \ / : * ? " < > | with _.
// For Linux and macOS, it replaces / with _.
// For unknown systems, it defaults to Linux-like behavior.
//
// Parameters:
// - fileName: the name of the file to be cleaned
//
// Returns:
// - string: the cleaned file name
func CleanInvalidChars(fileName string) string {
	switch runtime.GOOS {
	case "windows":
		// Windows invalid characters: \ / : * ? " < > |
		invalidChars := []string{"\\", "/", ":", "*", "?", "\"", "<", ">", "|"}
		for _, char := range invalidChars {
			fileName = strings.ReplaceAll(fileName, char, "_")
		}
	case "linux", "darwin": // "darwin" is the GOOS value for macOS
		// Linux and macOS invalid characters: /
		fileName = strings.ReplaceAll(fileName, "/", "_")
	default:
		// Default to Linux-like behavior for unknown systems (e.g., some IoT systems)
		fileName = strings.ReplaceAll(fileName, "/", "_")
	}
	return fileName
}

// IsFileNameValid checks if the file name is valid based on the current operating system.
// For Windows, it checks for the presence of \ / : * ? " < > |.
// For Linux and macOS, it checks for the presence of /.
// For unknown systems, it defaults to Linux-like behavior.
//
// Parameters:
// - fileName: the name of the file to be checked
//
// Returns:
// - bool: true if the file name is valid, false otherwise
func IsFileNameValid(fileName string) bool {
	switch runtime.GOOS {
	case "windows":
		// Windows invalid characters: \ / : * ? " < > |
		invalidChars := []string{"\\", "/", ":", "*", "?", "\"", "<", ">", "|"}
		for _, char := range invalidChars {
			if strings.Contains(fileName, char) {
				return false
			}
		}
	case "linux", "darwin": // "darwin" is the GOOS value for macOS
		// Linux and macOS invalid characters: /
		if strings.Contains(fileName, "/") {
			return false
		}
	default:
		// Default to Linux-like behavior for unknown systems (e.g., some IoT systems)
		if strings.Contains(fileName, "/") {
			return false
		}
	}
	return true
}
