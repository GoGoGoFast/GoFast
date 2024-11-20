package reutil_test

import (
	"GoFast/pkg/util/reutil"
	"testing"
)

func TestGet(t *testing.T) {
	content := "The price is $15"
	regex := `\$(\d+)`
	result, err := reutil.Get(regex, content, 1)
	if err != nil {
		t.Errorf("Get() error = %v", err)
	}
	if result != "15" {
		t.Errorf("Get() = %v, want %v", result, "15")
	}
}

func TestExtractMulti(t *testing.T) {
	content := "My name is John Doe"
	regex := `My name is (\w+) (\w+)`
	template := "First: $1, Last: $2"
	result, err := reutil.ExtractMulti(regex, content, template)
	if err != nil {
		t.Errorf("ExtractMulti() error = %v", err)
	}
	expected := "First: John, Last: Doe"
	if result != expected {
		t.Errorf("ExtractMulti() = %v, want %v", result, expected)
	}
}

func TestDelFirst(t *testing.T) {
	content := "The quick brown fox jumps over the lazy dog"
	regex := `\bfox\b`
	result, err := reutil.DelFirst(regex, content)
	if err != nil {
		t.Errorf("DelFirst() error = %v", err)
	}
	expected := "The quick brown  jumps over the lazy dog"
	if result != expected {
		t.Errorf("DelFirst() = %v, want %v", result, expected)
	}
}

func TestFindAll(t *testing.T) {
	content := "123 abc 456 def 789"
	regex := `\d+`
	result, err := reutil.FindAll(regex, content)
	if err != nil {
		t.Errorf("FindAll() error = %v", err)
	}
	expected := []string{"123", "456", "789"}
	if len(result) != len(expected) {
		t.Errorf("FindAll() = %v, want %v", result, expected)
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("FindAll() result[%d] = %v, want %v", i, result[i], expected[i])
		}
	}
}

func TestGetFirstNumber(t *testing.T) {
	content := "abc 123 def"
	result, err := reutil.GetFirstNumber(content)
	if err != nil {
		t.Errorf("GetFirstNumber() error = %v", err)
	}
	expected := 123
	if result != expected {
		t.Errorf("GetFirstNumber() = %v, want %v", result, expected)
	}
}

func TestIsMatch(t *testing.T) {
	content := "hello world"
	regex := `hello`
	match, err := reutil.IsMatch(regex, content)
	if err != nil {
		t.Errorf("IsMatch() error = %v", err)
	}
	if !match {
		t.Errorf("IsMatch() = %v, want true", match)
	}
}

func TestReplaceAll(t *testing.T) {
	content := "The numbers are 123 and 456."
	regex := `(\d+)`
	replacement := "[$1]"
	result, err := reutil.ReplaceAll(content, regex, replacement)
	if err != nil {
		t.Errorf("ReplaceAll() error = %v", err)
	}
	expected := "The numbers are [123] and [456]."
	if result != expected {
		t.Errorf("ReplaceAll() = %v, want %v", result, expected)
	}
}

func TestEscape(t *testing.T) {
	input := "1+1=2"
	result := reutil.Escape(input)
	expected := `1\+1=2`
	if result != expected {
		t.Errorf("Escape() = %v, want %v", result, expected)
	}
}
