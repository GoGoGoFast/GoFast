// Package resource provides utility functions for handling resources from byte arrays.
// This package simplifies reading, streaming, and buffering byte array resources.
package resource

import (
	"bufio"
	"bytes"
	"io"
	"net/url"
)

// BytesResource reads resources from a byte array
type BytesResource struct {
	name string
	data []byte
}

// NewBytesResource creates a new BytesResource instance
//
// Parameters:
// - name: the name of the resource
// - data: the byte array data of the resource
//
// Returns:
// - *BytesResource: a new BytesResource instance
func NewBytesResource(name string, data []byte) *BytesResource {
	return &BytesResource{name: name, data: data}
}

// GetName returns the name of the resource
//
// Returns:
// - string: the name of the resource
func (r *BytesResource) GetName() string {
	return r.name
}

// GetUrl returns the URL of the resource (nil in this implementation)
//
// Returns:
// - *url.URL: the URL of the resource
func (r *BytesResource) GetUrl() *url.URL {
	return nil
}

// GetStream returns an input stream of the resource
//
// Returns:
// - io.Reader: an input stream of the resource
// - error: if an error occurs
func (r *BytesResource) GetStream() (io.Reader, error) {
	return bytes.NewReader(r.data), nil
}

// GetReader returns a BufferedReader of the resource
//
// Parameters:
// - charset: the charset to use (not used in this implementation)
//
// Returns:
// - *bufio.Reader: a BufferedReader of the resource
// - error: if an error occurs
func (r *BytesResource) GetReader(charset string) (*bufio.Reader, error) {
	return bufio.NewReader(bytes.NewReader(r.data)), nil
}

// ReadStr reads the content of the resource as a string
//
// Parameters:
// - charset: the charset to use (not used in this implementation)
//
// Returns:
// - string: the content of the resource as a string
// - error: if an error occurs
func (r *BytesResource) ReadStr(charset string) (string, error) {
	return string(r.data), nil
}
