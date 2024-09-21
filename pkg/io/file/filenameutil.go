// Package fileutil provides utility functions for handling file names and paths
// across different operating systems. This package ensures that file names are
// valid and cleans invalid characters based on the operating system's rules.
//
// Package fileutil 提供了处理文件名和路径的实用函数，
// 适用于不同的操作系统。该包确保文件名有效，并根据操作系统的规则清理无效字符。
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
// GetBaseName 返回没有扩展名的文件基本名。
// 例如，给定文件路径 "example/test/file.txt"，它将返回 "file"。
//
// Parameters:
// - filePath: the path of the file (文件路径)
//
// Returns:
// - string: the base name of the file (文件基本名)
func GetBaseName(filePath string) string {
	return strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
}

// GetExtension returns the file extension.
// For example, given the file path "example/test/file.txt", it will return ".txt".
//
// GetExtension 返回文件的扩展名。
// 例如，给定文件路径 "example/test/file.txt"，它将返回 ".txt"。
//
// Parameters:
// - filePath: the path of the file (文件路径)
//
// Returns:
// - string: the file extension (文件扩展名)
func GetExtension(filePath string) string {
	return filepath.Ext(filePath)
}

// GetAbsolutePath returns the absolute path for a given relative or classpath file.
// It resolves the absolute path based on the current working directory.
// If the operation fails, it returns an error.
//
// GetAbsolutePath 返回给定相对或类路径文件的绝对路径。
// 它根据当前工作目录解析绝对路径。
// 如果操作失败，它将返回一个错误。
//
// Parameters:
// - path: the relative or classpath file (相对或类路径文件)
//
// Returns:
// - string: the absolute path (绝对路径)
// - error: if an error occurs (如果发生错误)
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
// CleanInvalidChars 根据当前操作系统清理文件名中的无效字符。
// 对于 Windows，它将 \ / : * ? " < > | 替换为 _。
// 对于 Linux 和 macOS，它将 / 替换为 _。
// 对于未知系统，它默认为 Linux 类行为。
//
// Parameters:
// - fileName: the name of the file to be cleaned (要清理的文件名)
//
// Returns:
// - string: the cleaned file name (清理后的文件名)
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

// IsFileNameValid 根据当前操作系统检查文件名是否有效。
// 对于 Windows，它检查是否存在 \ / : * ? " < > |。
// 对于 Linux 和 macOS，它检查是否存在 /。
// 对于未知系统，它默认为 Linux 类行为。
//
// Parameters:
// - fileName: the name of the file to be checked (要检查的文件名)
//
// Returns:
// - bool: true if the file name is valid, false otherwise (如果文件名有效则返回 true，否则返回 false)
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
