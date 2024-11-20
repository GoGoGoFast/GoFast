package charsetutil_test

import (
	"GoAllInOne/pkg/util/charsetutil"
	"golang.org/x/text/encoding/unicode"
	"testing"
)

// 测试 Charset 方法的正确性
func TestCharset(t *testing.T) {
	cu := &charsetutil.CharsetUtil{}

	// 测试 ISO-8859-1 编码
	isoEncoding, err := cu.Charset(charsetutil.ISO_8859_1)
	if err != nil {
		t.Errorf("获取 ISO-8859-1 编码失败: %v", err)
	}
	if isoEncoding == nil {
		t.Errorf("ISO-8859-1 编码为空")
	}

	// 测试 UTF-8 编码
	utf8Encoding, err := cu.Charset(charsetutil.UTF_8)
	if err != nil {
		t.Errorf("获取 UTF-8 编码失败: %v", err)
	}
	if utf8Encoding == nil {
		t.Errorf("UTF-8 编码为空")
	}

	// 测试 GBK 编码
	gbkEncoding, err := cu.Charset(charsetutil.GBK)
	if err != nil {
		t.Errorf("获取 GBK 编码失败: %v", err)
	}
	if gbkEncoding == nil {
		t.Errorf("GBK 编码为空")
	}

	// 测试不支持的编码
	_, err = cu.Charset("Unsupported")
	if err == nil {
		t.Errorf("对不支持的编码未返回错误")
	}
}

// 测试 Convert 方法
func TestConvert(t *testing.T) {
	cu := &charsetutil.CharsetUtil{}

	// 测试从 UTF-8 转到 ISO-8859-1
	input := "Hello, World!"
	result, err := cu.Convert(input, charsetutil.UTF_8, charsetutil.ISO_8859_1)
	if err != nil {
		t.Errorf("转换到 ISO-8859-1 失败: %v", err)
	}
	if result != input {
		t.Errorf("转换到 ISO-8859-1 后字符串不匹配，期望: %s, 实际: %s", input, result)
	}

	// 测试从 ISO-8859-1 转回 UTF-8
	result, err = cu.Convert(result, charsetutil.ISO_8859_1, charsetutil.UTF_8)
	if err != nil {
		t.Errorf("转换回 UTF-8 失败: %v", err)
	}
	if result != input {
		t.Errorf("从 ISO-8859-1 转回 UTF-8 后字符串不匹配，期望: %s, 实际: %s", input, result)
	}

	// 测试从 UTF-8 转到 GBK
	result, err = cu.Convert(input, charsetutil.UTF_8, charsetutil.GBK)
	if err != nil {
		t.Errorf("转换到 GBK 失败: %v", err)
	}

	// 测试从 GBK 转回 UTF-8
	result, err = cu.Convert(result, charsetutil.GBK, charsetutil.UTF_8)
	if err != nil {
		t.Errorf("从 GBK 转回 UTF-8 失败: %v", err)
	}
	if result != input {
		t.Errorf("从 GBK 转回 UTF-8 后字符串不匹配，期望: %s, 实际: %s", input, result)
	}
}

// 测试 DefaultCharset 和 DefaultCharsetName
func TestDefaultCharset(t *testing.T) {
	cu := &charsetutil.CharsetUtil{}

	// 测试 DefaultCharset
	if cu.DefaultCharset() != unicode.UTF8 {
		t.Errorf("DefaultCharset 与预期不符")
	}

	// 测试 DefaultCharsetName
	if cu.DefaultCharsetName() != charsetutil.UTF_8 {
		t.Errorf("DefaultCharsetName 与预期不符，期望: %s, 实际: %s", charsetutil.UTF_8, cu.DefaultCharsetName())
	}
}
