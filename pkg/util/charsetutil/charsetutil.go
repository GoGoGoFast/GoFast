package charsetutil

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
)

// 常量定义
const (
	ISO_8859_1 = "ISO-8859-1"
	UTF_8      = "UTF-8"
	GBK        = "GBK"
)

// CharsetUtil 字符编码工具类
type CharsetUtil struct{}

// Charset 方法将编码字符串转为对应的 Charset 对象
func (c *CharsetUtil) Charset(name string) (encoding.Encoding, error) {
	switch strings.ToUpper(name) {
	case ISO_8859_1:
		return charmap.ISO8859_1, nil
	case UTF_8:
		return unicode.UTF8, nil
	case GBK:
		return charmap.Windows1252, nil // 使用 Windows1252 作为临时替代
	default:
		return nil, fmt.Errorf("不支持的编码: %s", name)
	}
}

// Convert 方法在两种编码之间转换
func (c *CharsetUtil) Convert(input string, from string, to string) (string, error) {
	var fromEncoding encoding.Encoding
	var toEncoding encoding.Encoding

	if from != UTF_8 {
		var err error
		fromEncoding, err = c.Charset(from)
		if err != nil {
			return "", err
		}
	}

	if to != UTF_8 {
		var err error
		toEncoding, err = c.Charset(to)
		if err != nil {
			return "", err
		}
	}

	// 将输入字符串转为[]byte
	src := []byte(input)

	var transformed []byte
	var err error

	// 如果需要从其他编码转换到 UTF-8
	if fromEncoding != nil {
		reader := transform.NewReader(bytes.NewReader(src), fromEncoding.NewDecoder())
		transformed, err = ioutil.ReadAll(reader)
		if err != nil {
			return "", err
		}
		src = transformed
	}

	// 如果需要将 UTF-8 转换到其他编码
	if toEncoding != nil {
		reader := transform.NewReader(bytes.NewReader(src), toEncoding.NewEncoder())
		transformed, err = ioutil.ReadAll(reader)
		if err != nil {
			return "", err
		}
		return string(transformed), nil
	}

	return string(src), nil
}

// DefaultCharset 返回系统默认编码
func (c *CharsetUtil) DefaultCharset() encoding.Encoding {
	return unicode.UTF8
}

// DefaultCharsetName 返回系统默认编码字符串
func (c *CharsetUtil) DefaultCharsetName() string {
	return UTF_8
}
