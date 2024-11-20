package ziputil_test

import (
	"os"
	"path/filepath"
	"testing"

	"GoAllInOne/pkg/util/ziputil"
	"github.com/stretchr/testify/assert"
)

func TestZipUnzip(t *testing.T) {
	// 创建临时目录
	sourceDir, err := os.MkdirTemp("", "ziptest_src")
	assert.NoError(t, err)
	defer os.RemoveAll(sourceDir)

	// 创建临时文件
	tmpFile, err := os.CreateTemp(sourceDir, "file*.txt")
	assert.NoError(t, err)
	_, err = tmpFile.WriteString("This is a test file.")
	assert.NoError(t, err)
	defer tmpFile.Close()

	// 创建临时目录
	subDir, err := os.MkdirTemp(sourceDir, "subdir")
	assert.NoError(t, err)

	// 再次创建临时文件
	tmpFile2, err := os.CreateTemp(subDir, "file*.txt")
	assert.NoError(t, err)
	_, err = tmpFile2.WriteString("This is another test file.")
	assert.NoError(t, err)
	defer func(tmpFile2 *os.File) {
		err := tmpFile2.Close()
		if err != nil {

		}
	}(tmpFile2)

	// 创建目标zip文件
	targetZip := filepath.Join(os.TempDir(), "test.zip")
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(targetZip)

	// 压缩目录
	err = ziputil.Zip(sourceDir, targetZip, true)
	assert.NoError(t, err)

	// 创建解压目标目录
	destDir, err := os.MkdirTemp("", "ziptest_dest")
	assert.NoError(t, err)
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
	}(destDir)

	// 解压文件
	err = ziputil.Unzip(targetZip, destDir)
	assert.NoError(t, err)

	// 验证是否解压成功
	extractedFile1 := filepath.Join(destDir, filepath.Base(sourceDir), filepath.Base(tmpFile.Name()))
	extractedFile2 := filepath.Join(destDir, filepath.Base(sourceDir), filepath.Base(subDir), filepath.Base(tmpFile2.Name()))

	assert.FileExists(t, extractedFile1)
	assert.FileExists(t, extractedFile2)

	content1, err := os.ReadFile(extractedFile1)
	assert.NoError(t, err)
	assert.Equal(t, "This is a test file.", string(content1))

	content2, err := os.ReadFile(extractedFile2)
	assert.NoError(t, err)
	assert.Equal(t, "This is another test file.", string(content2))
}

func TestZipFile(t *testing.T) {
	// 创建临时文件
	tmpFile, err := os.CreateTemp("", "file*.txt")
	assert.NoError(t, err)
	_, err = tmpFile.WriteString("This is a test file.")
	assert.NoError(t, err)
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// 创建目标zip文件
	targetZip := filepath.Join(os.TempDir(), "test_file.zip")
	defer os.Remove(targetZip)

	// 压缩文件
	err = ziputil.Zip(tmpFile.Name(), targetZip, false)
	assert.NoError(t, err)

	// 创建解压目标目录
	destDir, err := os.MkdirTemp("", "ziptest_dest")
	assert.NoError(t, err)
	defer os.RemoveAll(destDir)

	// 解压文件
	err = ziputil.Unzip(targetZip, destDir)
	assert.NoError(t, err)

	// 验证是否解压成功
	extractedFile := filepath.Join(destDir, filepath.Base(tmpFile.Name()))
	assert.FileExists(t, extractedFile)

	content, err := os.ReadFile(extractedFile)
	assert.NoError(t, err)
	assert.Equal(t, "This is a test file.", string(content))
}
