package escapeutil_test

import (
	"testing"

	"GoAllInOne/pkg/util/escapeutil"
)

// 测试Escape函数
func TestEscape(t *testing.T) {
	original := "Hello World! 你好，世界！"
	expected := "Hello+World%21+%E4%BD%A0%E5%A5%BD%EF%BC%8C%E4%B8%96%E7%95%8C%EF%BC%81"
	encoded := escapeutil.Escape(original)

	if encoded != expected {
		t.Errorf("Escape() = %v; want %v", encoded, expected)
	}
}

// 测试Unescape函数
func TestUnescape(t *testing.T) {
	encoded := "Hello+World%21+%E4%BD%A0%E5%A5%BD%EF%BC%8C%E4%B8%96%E7%95%8C%EF%BC%81"
	expected := "Hello World! 你好，世界！"
	decoded, err := escapeutil.Unescape(encoded)

	if err != nil {
		t.Errorf("Unescape() error = %v", err)
	}

	if decoded != expected {
		t.Errorf("Unescape() = %v; want %v", decoded, expected)
	}
}

// 测试SafeUnescape函数
func TestSafeUnescape(t *testing.T) {
	original := "Hello World! 你好，世界！"
	encoded := "Hello+World%21+%E4%BD%A0%E5%A5%BD%EF%BC%8C%E4%B8%96%E7%95%8C%EF%BC%81"

	// Case 1: 正常解码
	expectedDecoded := original
	safeDecoded := escapeutil.SafeUnescape(encoded)
	if safeDecoded != expectedDecoded {
		t.Errorf("SafeUnescape() = %v; want %v", safeDecoded, expectedDecoded)
	}

	// Case 2: 输入未被编码的字符串，应该返回原文
	expectedOriginal := original
	safeDecodedOriginal := escapeutil.SafeUnescape(original)
	if safeDecodedOriginal != expectedOriginal {
		t.Errorf("SafeUnescape() = %v; want %v", safeDecodedOriginal, expectedOriginal)
	}
}
