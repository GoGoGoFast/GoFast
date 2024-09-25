package strutil

import (

	strutil2 "VeloCore/pkg/util/strutil"
	"testing"
)

func TestConcat(t *testing.T) {
	result := strutil2.Concat("Go", "Lang", "!")
	expected := "GoLang!"
	if result != expected {
		t.Errorf("Concat() = %v; want %v", result, expected)
	}
}

func TestSubstring(t *testing.T) {
	result := strutil2.Substring("GoLang", 0, 2)
	expected := "Go"
	if result != expected {
		t.Errorf("Substring() = %v; want %v", result, expected)
	}
}

func TestToCamelCase(t *testing.T) {
	result := strutil2.ToCamelCase("go_lang_util")
	expected := "goLangUtil"
	if result != expected {
		t.Errorf("ToCamelCase() = %v; want %v", result, expected)
	}
}

func TestToSnakeCase(t *testing.T) {
	result := strutil2.ToSnakeCase("GoLangUtil")
	expected := "go_lang_util"
	if result != expected {
		t.Errorf("ToSnakeCase() = %v; want %v", result, expected)
	}
}

func TestBase64(t *testing.T) {
	original := "hello"
	encoded := strutil2.EncodeBase64(original)
	decoded, err := strutil2.DecodeBase64(encoded)
	if err != nil || decoded != original {
		t.Errorf("Base64 encoding/decoding failed: %v", err)
	}
}

func TestReverse(t *testing.T) {
	result := strutil2.Reverse("abc")
	expected := "cba"
	if result != expected {
		t.Errorf("Reverse() = %v; want %v", result, expected)
	}
}

// func TestMatchRegex(t *testing.T) {
// 	matched, _ := strutil.MatchRegex("golang123", `golang\d+`)
// 	if !matched {
// 		t.Errorf("MatchRegex() failed; expected true, got false")
// 	}
// }

// func TestTrim(t *testing.T) {
// 	result := strutil.Trim("  hello  ")
// 	expected := "hello"
// 	if result != expected {
// 		t.Errorf("Trim() = %v; want %v", result, expected)
// 	}
// }

func TestHasBlank(t *testing.T) {
	result := strutil2.HasBlank("Go Lang")
	if !result {
		t.Errorf("HasBlank() = false; want true")
	}
}

func TestHasEmpty(t *testing.T) {
	result := strutil2.HasEmpty("Go", "", "Lang")
	if !result {
		t.Errorf("HasEmpty() = false; want true")
	}
}

func TestRemovePrefix(t *testing.T) {
	result := strutil2.RemovePrefix("GoLang", "Go")
	expected := "Lang"
	if result != expected {
		t.Errorf("RemovePrefix() = %v; want %v", result, expected)
	}
}

func TestRemoveSuffix(t *testing.T) {
	result := strutil2.RemoveSuffix("GoLang", "Lang")
	expected := "Go"
	if result != expected {
		t.Errorf("RemoveSuffix() = %v; want %v", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := strutil2.Sub("GoLang", 0, 2)
	expected := "Go"
	if result != expected {
		t.Errorf("Sub() = %v; want %v", result, expected)
	}

	result = strutil2.Sub("GoLang", -4, 4)
	expected = "Lang"
	if result != expected {
		t.Errorf("Sub() with negative index = %v; want %v", result, expected)
	}
}

func TestStrBytes(t *testing.T) {
	str := "hello"
	bytes := strutil2.Bytes(str)
	newStr := strutil2.Str(bytes)
	if str != newStr {
		t.Errorf("Str(Bytes()) = %v; want %v", newStr, str)
	}
}

func TestFormat(t *testing.T) {
	result := strutil2.Format("Hello, %s!", "World")
	expected := "Hello, World!"
	if result != expected {
		t.Errorf("Format() = %v; want %v", result, expected)
	}
}
