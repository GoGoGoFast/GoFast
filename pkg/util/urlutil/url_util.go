package urlutil

import (
	"net/url"
	"path/filepath"
)

// URLUtil 提供了URL操作的实用工具。
// URLUtil provides utilities for URL operations.
type URLUtil struct{}

// NewURL 通过字符串创建URL对象。
// NewURL creates a URL object from a string.
// 参数 (param): urlStr string - URL字符串。
// 参数 (param): urlStr string - the URL string.
// 返回值 (return): *url.URL - 对应的URL对象。
// 返回值 (return): *url.URL - the corresponding URL object.
// 返回值 (return): error - 如果解析失败，返回错误信息。
// 返回值 (return): error - returns an error if parsing fails.
func (u *URLUtil) NewURL(urlStr string) (*url.URL, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	return parsedURL, nil
}

// GetURL 通过资源名获取Classpath下的资源URL。
// GetURL retrieves the URL of a resource in the ClassPath by its name.
// 参数 (param): resource string - 资源名。
// 参数 (param): resource string - the resource name.
// 返回值 (return): *url.URL - 对应的URL对象。
// 返回值 (return): *url.URL - the corresponding URL object.
// 返回值 (return): error - 如果解析失败，返回错误信息。
// 返回值 (return): error - returns an error if parsing fails.
func (u *URLUtil) GetURL(resource string) (*url.URL, error) {
	// 假设资源在 ClassPath 下
	return u.NewURL("classpath://" + resource)
}

// Normalize 标准化URL，去除多余的斜杠和反斜杠。
// Normalize standardizes the URL by removing redundant slashes and backslashes.
// 参数 (param): urlStr string - 需要标准化的URL字符串。
// 参数 (param): urlStr string - the URL string to be normalized.
// 返回值 (return): string - 标准化后的URL字符串。
// 返回值 (return): string - the normalized URL string.
// 返回值 (return): error - 如果解析失败，返回错误信息。
// 返回值 (return): error - returns an error if parsing fails.
func (u *URLUtil) Normalize(urlStr string) (string, error) {
	parsedURL, err := u.NewURL(urlStr)
	if err != nil {
		return "", err
	}
	normalizedPath := filepath.Clean(parsedURL.Path)
	parsedURL.Path = normalizedPath
	return parsedURL.String(), nil
}

// Encode 进行URL编码。
// Encode encodes the given string using URL encoding.
// 参数 (param): content string - 需要编码的内容。
// 参数 (param): content string - the content to be encoded.
// 返回值 (return): string - 编码后的字符串。
// 返回值 (return): string - the encoded string.
func (u *URLUtil) Encode(content string) string {
	return url.QueryEscape(content)
}

// Decode 解码URL编码的字符串。
// Decode decodes the given URL-encoded string.
// 参数 (param): content string - 需要解码的内容。
// 参数 (param): content string - the content to be decoded.
// 返回值 (return): string - 解码后的字符串。
// 返回值 (return): string - the decoded string.
// 返回值 (return): error - 如果解码失败，返回错误信息。
// 返回值 (return): error - returns an error if decoding fails.
func (u *URLUtil) Decode(content string) (string, error) {
	return url.QueryUnescape(content)
}

// GetPath 获取URL的路径部分。
// GetPath retrieves the path part of a URL.
// 参数 (param): urlStr string - 需要解析的URL字符串。
// 参数 (param): urlStr string - the URL string to be parsed.
// 返回值 (return): string - URL的路径部分。
// 返回值 (return): string - the path part of the URL.
// 返回值 (return): error - 如果解析失败，返回错误信息。
// 返回值 (return): error - returns an error if parsing fails.
func (u *URLUtil) GetPath(urlStr string) (string, error) {
	parsedURL, err := u.NewURL(urlStr)
	if err != nil {
		return "", err
	}
	return parsedURL.Path, nil
}

// ToURI 将URL字符串转换为URI对象。
// ToURI converts a URL string to a URI object.
// 参数 (param): urlStr string - 需要转换的URL字符串。
// 参数 (param): urlStr string - the URL string to be converted.
// 返回值 (return): *url.URL - 对应的URI对象。
// 返回值 (return): *url.URL - the corresponding URI object.
// 返回值 (return): error - 如果解析失败，返回错误信息。
// 返回值 (return): error - returns an error if parsing fails.
func (u *URLUtil) ToURI(urlStr string) (*url.URL, error) {
	return u.NewURL(urlStr)
}
