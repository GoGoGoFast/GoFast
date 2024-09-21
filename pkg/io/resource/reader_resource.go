// Package resource provides utility functions for handling resources from various input sources.
// This package simplifies reading, streaming, and buffering resources.
//
// Package resource 提供了处理各种输入源资源的实用函数。
// 该包简化了读取、流式处理和缓冲资源。
package resource

import (
	"bufio"
	"bytes"
	"io"
	"net/url"
)

// ReaderResource reads resources from an io.Reader
//
// ReaderResource 从 io.Reader 读取资源
type ReaderResource struct {
	name string
	in   io.Reader
}

// NewReaderResource creates a new ReaderResource instance
//
// # NewReaderResource 创建一个新的 ReaderResource 实例
//
// Parameters:
// - name: the name of the resource (资源的名称)
// - in: the io.Reader input source (io.Reader 输入源)
//
// Returns:
// - *ReaderResource: a new ReaderResource instance (一个新的 ReaderResource 实例)
func NewReaderResource(name string, in io.Reader) *ReaderResource {
	return &ReaderResource{name: name, in: in}
}

// GetName returns the name of the resource
//
// # GetName 返回资源的名称
//
// Returns:
// - string: the name of the resource (资源的名称)
func (r *ReaderResource) GetName() string {
	return r.name
}

// GetUrl returns the URL of the resource (nil in this implementation)
//
// GetUrl 返回资源的 URL（在此实现中为 nil）
//
// Returns:
// - *url.URL: the URL of the resource (资源的 URL)
func (r *ReaderResource) GetUrl() *url.URL {
	return nil
}

// GetStream returns an input stream of the resource
//
// # GetStream 返回资源的输入流
//
// Returns:
// - io.Reader: an input stream of the resource (资源的输入流)
// - error: if an error occurs (如果发生错误)
func (r *ReaderResource) GetStream() (io.Reader, error) {
	return r.in, nil
}

// GetReader returns a BufferedReader of the resource
//
// # GetReader 返回资源的 BufferedReader
//
// Parameters:
// - charset: the charset to use (not used in this implementation) (要使用的字符集（在此实现中未使用）)
//
// Returns:
// - *bufio.Reader: a BufferedReader of the resource (资源的 BufferedReader)
// - error: if an error occurs (如果发生错误)
func (r *ReaderResource) GetReader(charset string) (*bufio.Reader, error) {
	return bufio.NewReader(r.in), nil
}

// ReadStr reads the content of the resource as a string
//
// # ReadStr 将资源的内容读取为字符串
//
// Parameters:
// - charset: the charset to use (not used in this implementation) (要使用的字符集（在此实现中未使用）)
//
// Returns:
// - string: the content of the resource as a string (资源的内容作为字符串)
// - error: if an error occurs (如果发生错误)
func (r *ReaderResource) ReadStr(charset string) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.in)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
