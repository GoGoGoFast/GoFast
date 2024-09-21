// Package fileutil provides utility functions for handling file names and paths
// across different operating systems. This package ensures that file names are
// valid and cleans invalid characters based on the operating system's rules.
//
// Package fileutil 提供了处理文件名和路径的实用函数，
// 适用于不同的操作系统。该包确保文件名有效，并根据操作系统的规则清理无效字符。
package fileutil

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Touch creates an empty file if it doesn't exist or updates the modified time if it does.
//
// Touch 如果文件不存在则创建一个空文件，如果文件存在则更新修改时间。
//
// Parameters:
// - path: the path of the file to touch (要创建或更新的文件路径)
//
// Returns:
// - error: if an error occurs (如果发生错误)
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
// Rm 删除一个文件或目录。如果是目录，它会递归删除所有内容。
//
// Parameters:
// - path: the path of the file or directory to remove (要删除的文件或目录的路径)
//
// Returns:
// - error: if an error occurs (如果发生错误)
func Rm(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("failed to remove file or directory at %s: %w", path, err)
	}
	return nil
}

// Cp copies a file or directory. If the source is a directory, it copies recursively.
//
// Cp 复制一个文件或目录。如果源是目录，它会递归复制。
//
// Parameters:
// - src: the source file or directory (源文件或目录)
// - dst: the destination file or directory (目标文件或目录)
//
// Returns:
// - error: if an error occurs (如果发生错误)
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

// Mv 重命名或移动一个文件/目录从 src 到 dst。
//
// Parameters:
// - src: the source file or directory (源文件或目录)
// - dst: the destination file or directory (目标文件或目录)
//
// Returns:
// - error: if an error occurs (如果发生错误)
func Mv(src, dst string) error {
	if err := os.Rename(src, dst); err != nil {
		return fmt.Errorf("failed to move file or directory from %s to %s: %w", src, dst, err)
	}
	return nil
}

// Mkdir creates a directory and all necessary parent directories.
//
// Mkdir 创建一个目录和所有必要的父目录。
//
// Parameters:
// - path: the path of the directory to create (要创建的目录路径)
//
// Returns:
// - error: if an error occurs (如果发生错误)
func Mkdir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory at %s: %w", path, err)
	}
	return nil
}

// IsEmpty checks if a file or directory is empty.
//
// IsEmpty 检查文件或目录是否为空。
//
// Parameters:
// - path: the path of the file or directory (要检查的文件或目录路径)
//
// Returns:
// - bool: true if the file or directory is empty, false otherwise (如果文件或目录为空则返回 true，否则返回 false)
// - error: if an error occurs (如果发生错误)
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
// IsFile 检查给定路径是否是文件。
//
// Parameters:
// - path: the path to check (要检查的路径)
//
// Returns:
// - bool: true if the path is a file, false otherwise (如果路径是文件则返回 true，否则返回 false)
func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// IsDir checks if the given path is a directory.
//
// IsDir 检查给定路径是否是目录。
//
// Parameters:
// - path: the path to check (要检查的路径)
//
// Returns:
// - bool: true if the path is a directory, false otherwise (如果路径是目录则返回 true，否则返回 false)
func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
