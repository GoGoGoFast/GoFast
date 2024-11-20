// Package fileutil provides utility functions for handling file names and paths
// across different operating systems. This package ensures that file names are
// valid and cleans invalid characters based on the operating system's rules.
package fileutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Touch creates an empty file if it doesn't exist or updates the modified time if it does.
//
// Parameters:
// - path: the path of the file to touch
//
// Returns:
// - error: if an error occurs
func Touch(path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to touch file at %s: %w", path, err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Printf("failed to close file at %s: %v\n", path, cerr)
		}
	}()
	return nil
}

// Rm removes a file or directory. If it's a directory, it removes all contents recursively.
//
// Parameters:
// - path: the path of the file or directory to remove
//
// Returns:
// - error: if an error occurs
func Rm(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to remove file or directory at %s: %w", path, err)
	}
	return nil
}

// Cp copies a file or directory. If the source is a directory, it copies recursively.
//
// Parameters:
// - src: the source file or directory
// - dst: the destination file or directory
//
// Returns:
// - error: if an error occurs
func Cp(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat source %s: %w", src, err)
	}

	if srcInfo.IsDir() {
		return copyDir(src, dst)
	} else {
		return copyFile(src, dst)
	}
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer func() {
		if cerr := sourceFile.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", dst, err)
	}
	defer func() {
		if cerr := destFile.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	if _, err := io.Copy(destFile, sourceFile); err != nil {
		return fmt.Errorf("failed to copy file %s to %s: %w", src, dst, err)
	}
	return nil
}

func copyDir(src, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read directory %s: %w", src, err)
	}

	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create destination directory %s: %w", dst, err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

// Mv renames or moves a file/directory from src to dst.
//
// Parameters:
// - src: the source file or directory
// - dst: the destination file or directory
//
// Returns:
// - error: if an error occurs
func Mv(src, dst string) error {
	if err := os.Rename(src, dst); err != nil {
		return fmt.Errorf("failed to move file or directory from %s to %s: %w", src, dst, err)
	}
	return nil
}

// Mkdir creates a directory and all necessary parent directories.
//
// Parameters:
// - path: the path of the directory to create
//
// Returns:
// - error: if an error occurs
func Mkdir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory at %s: %w", path, err)
	}
	return nil
}

// IsEmpty checks if a file or directory is empty.
//
// Parameters:
// - path: the path of the file or directory
//
// Returns:
// - bool: true if the file or directory is empty, false otherwise
// - error: if an error occurs
func IsEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("failed to open %s: %w", path, err)
	}
	defer func() {
		if cerr := f.Close(); cerr != nil {
			fmt.Printf("failed to close file at %s: %v\n", path, cerr)
		}
	}()

	stat, err := f.Stat()
	if err != nil {
		return false, fmt.Errorf("failed to stat %s: %w", path, err)
	}

	if stat.IsDir() {
		entries, err := f.ReadDir(1) // Read a single entry to check emptiness
		if err == io.EOF {
			return true, nil
		}
		if err != nil {
			return false, fmt.Errorf("failed to read directory %s: %w", path, err)
		}
		return len(entries) == 0, nil
	}

	return stat.Size() == 0, nil
}

// IsFile checks if the given path is a file.
//
// Parameters:
// - path: the path to check
//
// Returns:
// - bool: true if the path is a file, false otherwise
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsDir checks if the given path is a directory.
//
// Parameters:
// - path: the path to check
//
// Returns:
// - bool: true if the path is a directory, false otherwise
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
