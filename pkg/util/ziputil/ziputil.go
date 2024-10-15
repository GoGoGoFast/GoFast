package ziputil

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Zip 压缩目录或文件
func Zip(source, target string, includeBaseDir bool) error {
	// 创建目标zip文件
	zipFile, err := os.Create(target)
	if err != nil {
		return fmt.Errorf("无法创建目标压缩文件: %v", err)
	}
	// 确保在函数结束时关闭文件
	defer safeClose(zipFile)

	// 创建zip writer
	archive := zip.NewWriter(zipFile)
	defer safeCloseWriter(archive)

	// 获取待压缩文件的文件信息
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("无法获取源文件信息: %v", err)
	}

	// 开始压缩
	if sourceInfo.IsDir() {
		// 如果是目录，压缩整个目录
		baseDir := ""
		if includeBaseDir {
			baseDir = filepath.Base(source)
		}
		err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			// 设置压缩路径
			if baseDir != "" {
				header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			} else {
				header.Name = strings.TrimPrefix(path, source)
			}

			// 如果是目录，加上"/"
			if info.IsDir() {
				header.Name += "/"
			} else {
				header.Method = zip.Deflate
			}

			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}

			// 如果是文件，则将其内容写入zip文件
			if !info.IsDir() {
				err = writeFileContentToZip(path, writer) // 传递 writer 而不是 archive
				if err != nil {
					return err
				}
			}

			return nil
		})
	} else {
		// 如果是文件，直接压缩文件
		// 获取文件信息并创建对应的 zip 文件头
		header, err := zip.FileInfoHeader(sourceInfo)
		if err != nil {
			return fmt.Errorf("无法创建文件头: %v", err)
		}
		header.Method = zip.Deflate
		header.Name = filepath.Base(source)

		// 创建 zip 文件中的一个条目，并获取 io.Writer
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return fmt.Errorf("无法创建 zip 文件条目: %v", err)
		}

		// 将文件内容写入 zip 条目
		err = writeFileContentToZip(source, writer)
		if err != nil {
			return err
		}
	}

	return nil
}

// 将文件内容写入zip中
func writeFileContentToZip(filePath string, writer io.Writer) error {
	// 打开待压缩的文件
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("无法打开文件: %v", err)
	}
	defer safeClose(file)

	// 将文件内容复制到zip中
	_, err = io.Copy(writer, file)
	if err != nil {
		return fmt.Errorf("写入zip失败: %v", err)
	}

	return nil
}

// 安全关闭文件，避免重复关闭导致的错误
func safeClose(f *os.File) {
	if f != nil {
		err := f.Close()
		if err != nil && !errors.Is(err, os.ErrClosed) {
			fmt.Printf("关闭文件失败: %v\n", err)
		}
	}
}

// 安全关闭zip writer
func safeCloseWriter(w *zip.Writer) {
	if w != nil {
		err := w.Close()
		if err != nil {
			fmt.Printf("关闭zip writer失败: %v\n", err)
		}
	}
}

// 将文件内容写入zip中
func writeFileToZip(path string, writer io.Writer) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("无法打开文件: %v", err)
	}
	defer safeClose(file)

	_, err = io.Copy(writer, file)
	if err != nil {
		return fmt.Errorf("写入文件内容失败: %v", err)
	}
	return nil
}

// Unzip 解压文件
func Unzip(source, target string) error {
	reader, err := zip.OpenReader(source)
	if err != nil {
		return fmt.Errorf("无法打开zip文件: %v", err)
	}
	defer safeCloseReader(reader)

	for _, file := range reader.File {
		filePath := filepath.Join(target, file.Name)

		// 检查路径，防止zip slip攻击
		if !strings.HasPrefix(filePath, filepath.Clean(target)+string(os.PathSeparator)) {
			return fmt.Errorf("非法文件路径: %s", filePath)
		}

		// 如果是目录，则创建
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("创建目录失败: %v", err)
			}
		} else {
			// 如果是文件，解压文件
			err := extractFile(file, filePath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 安全关闭zip reader
func safeCloseReader(r *zip.ReadCloser) {
	if r != nil {
		err := r.Close()
		if err != nil && !errors.Is(err, os.ErrClosed) {
			fmt.Printf("关闭zip reader失败: %v\n", err)
		}
	}
}

// 解压具体的文件
func extractFile(file *zip.File, target string) error {
	// 打开压缩文件中的内容
	srcFile, err := file.Open()
	if err != nil {
		return fmt.Errorf("无法打开压缩文件: %v", err)
	}
	defer safeCloseCloser(srcFile)

	// 创建解压后的文件
	targetFile, err := os.Create(target)
	if err != nil {
		return fmt.Errorf("无法创建目标文件: %v", err)
	}
	defer safeClose(targetFile)

	_, err = io.Copy(targetFile, srcFile)
	if err != nil {
		return fmt.Errorf("解压文件失败: %v", err)
	}

	return nil
}

// 安全关闭io.Closer
func safeCloseCloser(c io.Closer) {
	if c != nil {
		err := c.Close()
		if err != nil && !errors.Is(err, os.ErrClosed) {
			fmt.Printf("关闭失败: %v\n", err)
		}
	}
}
