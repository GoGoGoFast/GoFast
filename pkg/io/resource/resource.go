package resource

import (
	"bufio"
	"io"
	"net/url"
)

// Resource A common resource interface is defined to obtain resource flow, read text and so on
type Resource interface {
	GetName() string
	GetUrl() *url.URL
	GetStream() (io.Reader, error)
	GetReader(charset string) (*bufio.Reader, error)
	ReadStr(charset string) (string, error)
}
