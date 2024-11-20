// Package resource provides utility functions for handling resources from various input sources.
// This package simplifies reading, streaming, and buffering resources.
package resource

import (
	"bufio"
	"bytes"
	"io"
	"net/url"
)

// ReaderResource reads resources from an io.Reader
type ReaderResource struct {
	name string
	in   io.Reader
}

// NewReaderResource creates a new ReaderResource instance
//
// Parameters:
// - name: the name of the resource
// - in: the io.Reader input source
//
// Returns:
// - *ReaderResource: a new ReaderResource instance
func NewReaderResource(name string, in io.Reader) *ReaderResource {
	return &ReaderResource{name: name, in: in}
}

// GetName returns the name of the resource
//
// Returns:
// - string: the name of the resource
func (r *ReaderResource) GetName() string {
	return r.name
}

// GetUrl returns the URL of the resource (nil in this implementation)
//
// Returns:
// - *url.URL: the URL of the resource
func (r *ReaderResource) GetUrl() *url.URL {
	return nil
}

// GetStream returns an input stream of the resource
//
// Returns:
// - io.Reader: an input stream of the resource
// - error: if an error occurs
func (r *ReaderResource) GetStream() (io.Reader, error) {
	return r.in, nil
}

// GetReader returns a BufferedReader of the resource
//
// Parameters:
// - charset: the charset to use (not used in this implementation)
//
// Returns:
// - *bufio.Reader: a BufferedReader of the resource
// - error: if an error occurs
func (r *ReaderResource) GetReader(charset string) (*bufio.Reader, error) {
	return bufio.NewReader(r.in), nil
}

// ReadStr reads the content of the resource as a string
//
// Parameters:
// - charset: the charset to use (not used in this implementation)
//
// Returns:
// - string: the content of the resource as a string
// - error: if an error occurs
func (r *ReaderResource) ReadStr(charset string) (string, error) {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.in)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
