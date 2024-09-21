// Package stream provides utility functions for handling file I/O operations.
// This package simplifies reading, writing, appending, copying files, and more.
//
// Package stream 提供了处理文件输入输出操作的实用函数。
// 该包简化了读取、写入、追加、复制文件等操作。
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
// ReadFile 读取文件内容并将其作为字符串返回。
// 它使用 os.ReadFile 来处理文件读取，这是现代且推荐的方法。
//
// Parameters:
// - path: the path of the file to read (要读取的文件路径)
//
// Returns:
// - string: the content of the file (文件内容)
// - error: if an error occurs (如果发生错误)
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
// WriteFile 将数据写入文件。如果文件不存在则创建文件，如果文件存在则覆盖文件。
// 它使用 os.WriteFile，这是现代的文件写入 API。
//
// Parameters:
// - path: the path of the file to write to (要写入的文件路径)
// - data: the data to write to the file (要写入文件的数据)
//
// Returns:
// - error: if an error occurs (如果发生错误)
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
// AppendToFile 将数据追加到现有文件。如果文件不存在则创建文件。
// 对延迟关闭文件操作进行了适当的错误处理。
//
// Parameters:
// - path: the path of the file to append to (要追加的文件路径)
// - data: the data to append to the file (要追加到文件的数据)
//
// Returns:
// - error: if an error occurs (如果发生错误)
func AppendToFile(path, data string) error {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for appending at %s: %w", path, err)
	}

	// Ensure that the file is properly closed and handle any errors during close.
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
// CopyFile 复制一个文件从 src 到 dst。如果目标文件存在，它将被覆盖。
// 它使用现代 I/O 方法以确保最佳性能和错误处理。
//
// Parameters:
// - src: the source file to copy (源文件)
// - dst: the destination file (目标文件)
//
// Returns:
// - error: if an error occurs (如果发生错误)
func CopyFile(src, dst string) error {
	// Open the source file for reading
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	// Ensure the source file is closed properly and handle any errors during the close operation
	defer func() {
		if cerr := sourceFile.Close(); cerr != nil {
			fmt.Printf("warning: failed to close source file %s properly: %v\n", src, cerr)
		}
	}()

	// Create the destination file for writing
	destinationFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	// Ensure the destination file is closed properly
	defer func() {
		if cerr := destinationFile.Close(); cerr != nil {
			fmt.Printf("warning: failed to close destination file %s properly: %v\n", dst, cerr)
		}
	}()

	// Copy data from the source file to the destination file
	if _, err = io.Copy(destinationFile, sourceFile); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	// Set the file mode of the destination file to match the source file
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
// ReadDirFiles 读取目录中所有文件的内容，并将它们作为映射返回，
// 映射的键为文件名，值为文件内容。它支持通过扩展名过滤文件。
//
// Parameters:
// - dirPath: the path of the directory to read (要读取的目录路径)
// - ext: the file extension to filter by (用于过滤的文件扩展名)
//
// Returns:
// - map[string]string: a map of file names to file contents (文件名到文件内容的映射)
// - error: if an error occurs (如果发生错误)
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
// Exists 检查给定路径上的文件或目录是否存在。
// 它使用 os.Stat，这是检查文件是否存在的高效且广泛使用的方法。
//
// Parameters:
// - path: the path to check (要检查的路径)
//
// Returns:
// - bool: true if the file or directory exists, false otherwise (如果文件或目录存在则返回 true，否则返回 false)
func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// ListFiles returns a list of all files (not directories) in the given directory.
// If the recursive flag is true, it will search through subdirectories as well.
//
// ListFiles 返回给定目录中所有文件（不包括目录）的列表。
// 如果 recursive 标志为 true，它还会搜索子目录。
//
// Parameters:
// - dir: the directory to list files from (要列出文件的目录)
// - recursive: whether to search through subdirectories (是否搜索子目录)
//
// Returns:
// - []string: a list of file paths (文件路径列表)
// - error: if an error occurs (如果发生错误)
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
