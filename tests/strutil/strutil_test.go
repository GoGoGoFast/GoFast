package strutil_test

import (
	"testing"

	"VeloCore/pkg/util/strutil"
)

func TestConcat(t *testing.T) {
	result := strutil.Concat("Hello", " ", "World")
	expected := "Hello World"
	if result != expected {
		t.Errorf("Concat() = %v; want %v", result, expected)
	}
}

func TestConcatWithSeparator(t *testing.T) {
	result := strutil.ConcatWithSeparator([]string{"Hello", "World"}, "-")
	expected := "Hello-World"
	if result != expected {
		t.Errorf("ConcatWithSeparator() = %v; want %v", result, expected)
	}
}

func TestSubstring(t *testing.T) {
	result := strutil.Substring("Hello World", 6, 11)
	expected := "World"
	if result != expected {
		t.Errorf("Substring() = %v; want %v", result, expected)
	}
}

func TestCountOccurrences(t *testing.T) {
	result := strutil.CountOccurrences("Hello World Hello", "Hello")
	expected := 2
	if result != expected {
		t.Errorf("CountOccurrences() = %v; want %v", result, expected)
	}
}

func TestEncodeBase64(t *testing.T) {
	result := strutil.EncodeBase64("Hello World")
	expected := "SGVsbG8gV29ybGQ="
	if result != expected {
		t.Errorf("EncodeBase64() = %v; want %v", result, expected)
	}
}

func TestDecodeBase64(t *testing.T) {
	result, err := strutil.DecodeBase64("SGVsbG8gV29ybGQ=")
	expected := "Hello World"
	if err != nil || result != expected {
		t.Errorf("DecodeBase64() = %v, %v; want %v, <nil>", result, err, expected)
	}
}

func TestURLSafeEncodeBase64(t *testing.T) {
	result := strutil.URLSafeEncodeBase64("Hello World")
	expected := "SGVsbG8gV29ybGQ="
	if result != expected {
		t.Errorf("URLSafeEncodeBase64() = %v; want %v", result, expected)
	}
}

func TestURLSafeDecodeBase64(t *testing.T) {
	result, err := strutil.URLSafeDecodeBase64("SGVsbG8gV29ybGQ=")
	expected := "Hello World"
	if err != nil || result != expected {
		t.Errorf("URLSafeDecodeBase64() = %v, %v; want %v, <nil>", result, err, expected)
	}
}

func TestToUpper(t *testing.T) {
	result := strutil.ToUpper("hello")
	expected := "HELLO"
	if result != expected {
		t.Errorf("ToUpper() = %v; want %v", result, expected)
	}
}

func TestToLower(t *testing.T) {
	result := strutil.ToLower("HELLO")
	expected := "hello"
	if result != expected {
		t.Errorf("ToLower() = %v; want %v", result, expected)
	}
}

func TestToCamelCase(t *testing.T) {
	result := strutil.ToCamelCase("hello_world")
	expected := "helloWorld"
	if result != expected {
		t.Errorf("ToCamelCase() = %v; want %v", result, expected)
	}
}

func TestToSnakeCase(t *testing.T) {
	result := strutil.ToSnakeCase("helloWorld")
	expected := "hello_world"
	if result != expected {
		t.Errorf("ToSnakeCase() = %v; want %v", result, expected)
	}
}

func TestCapitalize(t *testing.T) {
	result := strutil.Capitalize("hello world")
	expected := "Hello world"
	if result != expected {
		t.Errorf("Capitalize() = %v; want %v", result, expected)
	}
}

func TestHasBlank(t *testing.T) {
	result := strutil.HasBlank("Hello World")
	expected := true
	if result != expected {
		t.Errorf("HasBlank() = %v; want %v", result, expected)
	}
}

func TestHasNoBlank(t *testing.T) {
	result := strutil.HasNoBlank("HelloWorld")
	expected := true
	if result != expected {
		t.Errorf("HasNoBlank() = %v; want %v", result, expected)
	}
}

func TestHasEmpty(t *testing.T) {
	result := strutil.HasEmpty("Hello", "", "World")
	expected := true
	if result != expected {
		t.Errorf("HasEmpty() = %v; want %v", result, expected)
	}
}

func TestHasNoEmpty(t *testing.T) {
	result := strutil.HasNoEmpty("Hello", "World")
	expected := true
	if result != expected {
		t.Errorf("HasNoEmpty() = %v; want %v", result, expected)
	}
}

func TestIsBlank(t *testing.T) {
	result := strutil.IsBlank("   ")
	expected := true
	if result != expected {
		t.Errorf("IsBlank() = %v; want %v", result, expected)
	}
}

func TestNotBlank(t *testing.T) {
	result := strutil.NotBlank("Hello")
	expected := true
	if result != expected {
		t.Errorf("NotBlank() = %v; want %v", result, expected)
	}
}

func TestIsEmpty(t *testing.T) {
	result := strutil.IsEmpty("")
	expected := true
	if result != expected {
		t.Errorf("IsEmpty() = %v; want %v", result, expected)
	}
}

func TestNotEmpty(t *testing.T) {
	result := strutil.NotEmpty("Hello")
	expected := true
	if result != expected {
		t.Errorf("NotEmpty() = %v; want %v", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := strutil.Sub("abcdefgh", 2, 3)
	expected := "cde"
	if result != expected {
		t.Errorf("Sub() = %v; want %v", result, expected)
	}

	result = strutil.Sub("abcdefgh", -3, 2)
	expected = "fg"
	if result != expected {
		t.Errorf("Sub() = %v; want %v", result, expected)
	}
}

func TestStr(t *testing.T) {
	result := strutil.Str([]byte{72, 101, 108, 108, 111})
	expected := "Hello"
	if result != expected {
		t.Errorf("Str() = %v; want %v", result, expected)
	}
}

func TestBytes(t *testing.T) {
	result := strutil.Bytes("Hello")
	expected := []byte{72, 101, 108, 108, 111}
	if string(result) != string(expected) {
		t.Errorf("Bytes() = %v; want %v", result, expected)
	}
}

func TestReverse(t *testing.T) {
	result := strutil.Reverse("Hello")
	expected := "olleH"
	if result != expected {
		t.Errorf("Reverse() = %v; want %v", result, expected)
	}
}

func TestRemovePrefix(t *testing.T) {
	result := strutil.RemovePrefix("HelloWorld", "Hello")
	expected := "World"
	if result != expected {
		t.Errorf("RemovePrefix() = %v; want %v", result, expected)
	}
}

func TestRemoveSuffix(t *testing.T) {
	result := strutil.RemoveSuffix("HelloWorld", "World")
	expected := "Hello"
	if result != expected {
		t.Errorf("RemoveSuffix() = %v; want %v", result, expected)
	}
}

func TestFormat(t *testing.T) {
	result := strutil.Format("{} loves {}", "Alice", "Bob")
	expected := "Alice loves Bob"
	if result != expected {
		t.Errorf("Format() = %v; want %v", result, expected)
	}
}

func TestPadLeft(t *testing.T) {
	result := strutil.PadLeft("Go", 5, 'x')
	expected := "xxxGo"
	if result != expected {
		t.Errorf("PadLeft() = %v; want %v", result, expected)
	}
}

func TestPadRight(t *testing.T) {
	result := strutil.PadRight("Go", 5, 'x')
	expected := "Goxxx"
	if result != expected {
		t.Errorf("PadRight() = %v; want %v", result, expected)
	}
}

func TestJoinNonEmpty(t *testing.T) {
	result := strutil.JoinNonEmpty([]string{"Hello", "", "World"}, " ")
	expected := "Hello World"
	if result != expected {
		t.Errorf("JoinNonEmpty() = %v; want %v", result, expected)
	}
}

func TestRandomString(t *testing.T) {
	result, err := strutil.RandomString(10, "abcdefghijklmnopqrstuvwxyz")
	if err != nil {
		t.Fatalf("RandomString() failed with error: %v", err)
	}
	if len(result) != 10 {
		t.Errorf("RandomString() length = %d; want %d", len(result), 10)
	}
}

func TestIsNumeric(t *testing.T) {
	result := strutil.IsNumeric("12345")
	expected := true
	if result != expected {
		t.Errorf("IsNumeric() = %v; want %v", result, expected)
	}
}

func TestWordWrap(t *testing.T) {
	result := strutil.WordWrap("This is a very long string that needs to be wrapped properly.", 10)
	expected := "This is a\nvery long\nstring that\nneeds to be\nwrapped\nproperly."
	if result != expected {
		t.Errorf("WordWrap() = %v; want %v", result, expected)
	}
}
