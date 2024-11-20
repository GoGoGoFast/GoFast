// Package resource provides utility functions for handling resources from various input sources.
// This package simplifies reading, streaming, and buffering resources.
package resource

import (
	"bufio"
	"io"
	"net/url"
	"strings"
)

// StringResource reads resources from a string
type StringResource struct {
	name string
	data string
}

// NewStringResource creates a new StringResource instance
//
// Parameters:
// - name: the name of the resource
// - data: the string data of the resource
//
// Returns:
// - *StringResource: a new StringResource instance
func NewStringResource(name string, data string) *StringResource {
	return &StringResource{name: name, data: data}
}

// GetName returns the name of the resource
//
// Returns:
// - string: the name of the resource
func (r *StringResource) GetName() string {
	return r.name
}

// GetUrl returns the URL of the resource (nil in this implementation)
//
// Returns:
// - *url.URL: the URL of the resource
func (r *StringResource) GetUrl() *url.URL {
	return nil
}

// GetStream returns an input stream of the resource
//
// Returns:
// - io.Reader: an input stream of the resource
// - error: if an error occurs
func (r *StringResource) GetStream() (io.Reader, error) {
	return strings.NewReader(r.data), nil
}

// GetReader returns a BufferedReader of the resource
//
// Parameters:
// - charset: the charset to use (not used in this implementation)
//
// Returns:
// - *bufio.Reader: a BufferedReader of the resource
// - error: if an error occurs
func (r *StringResource) GetReader(charset string) (*bufio.Reader, error) {
	return bufio.NewReader(strings.NewReader(r.data)), nil
}

// ReadStr reads the content of the resource as a string
//
// Parameters:
// - charset: the charset to use (not used in this implementation)
//
// Returns:
// - string: the content of the resource as a string
// - error: if an error occurs
func (r *StringResource) ReadStr(charset string) (string, error) {
	return r.data, nil
}
