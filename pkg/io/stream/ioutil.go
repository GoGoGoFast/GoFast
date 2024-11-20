// Package stream provides utility functions for handling file I/O operations.
// This package simplifies reading, writing, appending, copying files, and more.
package stream

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// ReadFile reads the content of a file and returns it as a string.
// It uses os.ReadFile to handle the file reading, which is the modern and recommended approach.
//
// Parameters:
// - path: the path of the file to read
//
// Returns:
// - string: the content of the file
// - error: if an error occurs
func ReadFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read file at %s: %w", path, err)
	}
	return string(content), nil
}

// WriteFile writes data to a file. It creates the file if it doesn't exist
// and overwrites it if it does. This uses os.WriteFile, which is a modern API for file writing.
//
// Parameters:
// - path: the path of the file to write to
// - data: the data to write to the file
//
// Returns:
// - error: if an error occurs
func WriteFile(path, data string) error {
	err := os.WriteFile(path, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file at %s: %w", path, err)
	}
	return nil
}

// AppendToFile appends data to an existing file. If the file does not exist, it will create it.
// Proper error handling is done for the deferred file close operation.
//
// Parameters:
// - path: the path of the file to append to
// - data: the data to append to the file
//
// Returns:
// - error: if an error occurs
func AppendToFile(path, data string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for appending at %s: %w", path, err)
	}

	defer func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Printf("warning: failed to close file %s properly: %v\n", path, cerr)
		}
	}()

	if _, err := f.WriteString(data); err != nil {
		return fmt.Errorf("failed to append to file at %s: %w", path, err)
	}
	return nil
}

// CopyFile copies a file from src to dst. If the destination file exists, it will be overwritten.
// This uses modern I/O methods to ensure optimal performance and error handling.
//
// Parameters:
// - src: the source file to copy
// - dst: the destination file
//
// Returns:
// - error: if an error occurs
func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer func() {
		if cerr := sourceFile.Close(); cerr != nil {
			fmt.Printf("warning: failed to close source file %s properly: %v\n", src, cerr)
		}
	}()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer func() {
		if cerr := destinationFile.Close(); cerr != nil {
			fmt.Printf("warning: failed to close destination file %s properly: %v\n", dst, cerr)
		}
	}()

	if _, err = io.Copy(destinationFile, sourceFile); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	if info, err := os.Stat(src); err == nil {
		if err = os.Chmod(dst, info.Mode()); err != nil {
			return fmt.Errorf("failed to set file mode on destination file: %w", err)
		}
	} else {
		return fmt.Errorf("failed to retrieve file info from source: %w", err)
	}

	return nil
}

// ReadDirFiles reads the contents of all files in a directory and returns them as a map
// with file names as keys and file content as values. It supports filtering files by extension.
//
// Parameters:
// - dirPath: the path of the directory to read
// - ext: the file extension to filter by
//
// Returns:
// - map[string]string: a map of file names to file contents
// - error: if an error occurs
func ReadDirFiles(dirPath string, ext string) (map[string]string, error) {
	files := make(map[string]string)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ext) {
			content, err := ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", path, err)
			}
			files[info.Name()] = content
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}
	return files, nil
}

// Exists checks whether a file or directory exists at the given path.
// It uses os.Stat, which is efficient and widely used for checking file existence.
//
// Parameters:
// - path: the path to check
//
// Returns:
// - bool: true if the file or directory exists, false otherwise
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ListFiles returns a list of all files (not directories) in the given directory.
// If the recursive flag is true, it will search through subdirectories as well.
//
// Parameters:
// - dir: the directory to list files from
// - recursive: whether to search through subdirectories
//
// Returns:
// - []string: a list of file paths
// - error: if an error occurs
func ListFiles(dir string, recursive bool) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		} else if !recursive {
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list files in directory %s: %w", dir, err)
	}
	return files, nil
}
