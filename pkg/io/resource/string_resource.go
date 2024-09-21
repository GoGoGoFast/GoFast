// Package resource provides utility functions for handling resources from various input sources.
// This package simplifies reading, streaming, and buffering resources.
//
// Package resource 提供了处理各种输入源资源的实用函数。
// 该包简化了读取、流式处理和缓冲资源。
package resource

import (
	"bufio"
	"io"
	"net/url"
	"strings"
)

// StringResource reads resources from a string
//
// StringResource 从字符串读取资源
type StringResource struct {
	name string
	data string
}

// NewStringResource creates a new StringResource instance
//
// # NewStringResource 创建一个新的 StringResource 实例
//
// Parameters:
// - name: the name of the resource (资源的名称)
// - data: the string data of the resource (资源的字符串数据)
//
// Returns:
// - *StringResource: a new StringResource instance (一个新的 StringResource 实例)
func NewStringResource(name string, data string) *StringResource {
	return &StringResource{name: name, data: data}
}

// GetName returns the name of the resource
//
// # GetName 返回资源的名称
//
// Returns:
// - string: the name of the resource (资源的名称)
func (r *StringResource) GetName() string {
	return r.name
}

// GetUrl returns the URL of the resource (nil in this implementation)
//
// GetUrl 返回资源的 URL（在此实现中为 nil）
//
// Returns:
// - *url.URL: the URL of the resource (资源的 URL)
func (r *StringResource) GetUrl() *url.URL {
	return nil
}

// GetStream returns an input stream of the resource
//
// # GetStream 返回资源的输入流
//
// Returns:
// - io.Reader: an input stream of the resource (资源的输入流)
// - error: if an error occurs (如果发生错误)
func (r *StringResource) GetStream() (io.Reader, error) {
	return strings.NewReader(r.data), nil
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
func (r *StringResource) GetReader(charset string) (*bufio.Reader, error) {
	return bufio.NewReader(strings.NewReader(r.data)), nil
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
func (r *StringResource) ReadStr(charset string) (string, error) {
	return r.data, nil
}
