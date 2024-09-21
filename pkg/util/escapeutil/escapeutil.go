package escapeutil

import (
	"net/url"
)

// EscapeUtil 包提供字符串的转码和解码功能
// EscapeUtil package provides encoding and decoding functions for strings

// Escape 编码指定的字符串。
// Escape encodes the given string.
// 参数 (param): input string - 需要被编码的字符串。
// 参数 (param): input string - the string to be encoded.
// 返回值 (return): string - 编码后的字符串。
// 返回值 (return): string - the encoded string.
func Escape(input string) string {
	// 使用net/url包进行编码，保留ASCII字母、数字和特定标点符号
	// Use net/url package for encoding, keeping ASCII letters, numbers, and specific punctuation
	return url.QueryEscape(input)
}

// Unescape 解码指定的字符串。
// Unescape decodes the given string.
// 参数 (param): input string - 需要被解码的字符串。
// 参数 (param): input string - the string to be decoded.
// 返回值 (return): string - 解码后的字符串。
// 返回值 (return): string - the decoded string.
func Unescape(input string) (string, error) {
	// 使用net/url包进行解码
	// Use net/url package for decoding
	decoded, err := url.QueryUnescape(input)
	if err != nil {
		return "", err
	}
	return decoded, nil
}

// SafeUnescape 安全地解码指定的字符串，如果输入字符串不是被Escape编码的，返回原文。
// SafeUnescape safely decodes the given string, returning the original text if it was not encoded by Escape.
// 参数 (param): input string - 需要被解码的字符串。
// 参数 (param): input string - the string to be decoded.
// 返回值 (return): string - 解码后的字符串或者原文。
// 返回值 (return): string - the decoded string or the original text if not encoded.
func SafeUnescape(input string) string {
	decoded, err := Unescape(input)
	if err != nil {
		// 如果解码失败，返回输入字符串
		// If decoding fails, return the input string
		return input
	}
	return decoded
}

// EscapeUtil 包的测试代码
// Test code for EscapeUtil package
func main() {
	original := "Hello World! 你好，世界！"
	encoded := Escape(original)
	decoded, _ := Unescape(encoded)
	safeDecoded := SafeUnescape(original)

	println("Original:", original)
	println("Encoded:", encoded)
	println("Decoded:", decoded)
	println("Safe Decoded:", safeDecoded)
}
