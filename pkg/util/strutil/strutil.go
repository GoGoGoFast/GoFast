// Package strutil provides utility functions for string manipulation.
// This package includes various utility functions inspired by both the Go standard library
// and additional ideas from libraries like HuTool's StrUtil. It is designed to make common
package strutil

// string operations easy and efficient, with optimizations for performance.x

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	mathrand "math/rand"
	"regexp"
	"strings"
	"time"
	"unicode"
)

// Concat concatenates multiple strings into a single string.
// It takes a variadic parameter of strings and efficiently appends them using a strings.Builder.
//
// Example usage:
// result := Concat("Hello", " ", "World")
// result will be "Hello World".
func Concat(strs ...string) string {
	var builder strings.Builder
	for _, str := range strs {
		builder.WriteString(str)
	}
	return builder.String()
}

// ConcatWithSeparator concatenates multiple strings with a specified separator.
// It uses strings.Join for simplicity and efficiency.
// If the input slice is empty, an empty string is returned.
//
// Example usage:
// result := ConcatWithSeparator([]string{"Hello", "World"}, "-")
// result will be "Hello-World".
func ConcatWithSeparator(strs []string, separator string) string {
	if len(strs) == 0 {
		return ""
	}
	return strings.Join(strs, separator)
}

// Substring extracts a substring from a given string.
// The substring is defined by the start and end indices. The range is inclusive of the start index
// and exclusive of the end index. If the indices are out of bounds or invalid (e.g., start > end), an empty string is returned.
//
// Example usage:
// substring := Substring("Hello World", 6, 11)
// substring will be "World".
func Substring(s string, start, end int) string {
	if start < 0 {
		start = len(s) + start // Support for negative indexing
	}
	if end > len(s) || start > end || start < 0 {
		return ""
	}
	return s[start:end]
}

// CountOccurrences counts the number of occurrences of a specific substring in a given string.
// It repeatedly searches for the substring and adjusts the input string to start after each match.
//
// Example usage:
// count := CountOccurrences("Hello World Hello", "Hello")
// count will be 2.
func CountOccurrences(s, substr string) int {
	return strings.Count(s, substr)
}

// EncodeBase64 encodes a string to Base64 format.
// It takes a string as input and returns the Base64 encoded representation of that string.
//
// Example usage:
// encoded := EncodeBase64("Hello World")
// encoded will be "SGVsbG8gV29ybGQ=".
func EncodeBase64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// DecodeBase64 decodes a Base64 encoded string.
// It returns the decoded string and an error if the input string is not valid Base64.
//
// Example usage:
// decoded, err := DecodeBase64("SGVsbG8gV29ybGQ=")
// decoded will be "Hello World".
func DecodeBase64(s string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(s)
	return string(decoded), err
}

// URLSafeEncodeBase64 encodes a string in URL-safe Base64 format.
// This ensures that the encoded string can be safely used in URLs, as it replaces characters
// like '+' and '/' with URL-safe alternatives.
//
// Example usage:
// urlSafeEncoded := URLSafeEncodeBase64("Hello World")
// urlSafeEncoded will be "SGVsbG8gV29ybGQ=" (URL-safe version).
func URLSafeEncodeBase64(s string) string {
	return base64.URLEncoding.EncodeToString([]byte(s))
}

// URLSafeDecodeBase64 decodes a URL-safe Base64 encoded string.
// It behaves similarly to DecodeBase64 but for URL-safe Base64 encoded strings.
//
// Example usage:
// decoded, err := URLSafeDecodeBase64("SGVsbG8gV29ybGQ=")
// decoded will be "Hello World".
func URLSafeDecodeBase64(s string) (string, error) {
	decoded, err := base64.URLEncoding.DecodeString(s)
	return string(decoded), err
}

// ToUpper converts a string to uppercase.
// This function uses the standard library's strings.ToUpper.
//
// Example usage:
// result := ToUpper("hello")
// result will be "HELLO".
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower converts a string to lowercase.
// This function uses the standard library's strings.ToLower.
//
// Example usage:
// result := ToLower("HELLO")
// result will be "hello".
func ToLower(s string) string {
	return strings.ToLower(s)
}

// ToCamelCase efficiently converts a snake_case or kebab-case string to camelCase.
// It skips unnecessary rune operations for ASCII strings.
// except the first one, and joins them together.
//
// Example usage:
// result := ToCamelCase("hello_world")
// result will be "helloWorld".
func ToCamelCase(s string) string {
	words := strings.Split(s, "_")
	for i := 1; i < len(words); i++ {
		words[i] = Capitalize(words[i])
	}
	return strings.Join(words, "")
}

// ToSnakeCase converts a camelCase string to snake_case.
// It identifies uppercase characters in the input string and inserts an underscore before each one.
// The resulting string is entirely in lowercase.
//
// Example usage:
// result := ToSnakeCase("helloWorld")
// result will be "hello_world".
func ToSnakeCase(s string) string {
	var builder strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) && i != 0 {
			builder.WriteByte('_')
		}
		builder.WriteRune(unicode.ToLower(r))
	}
	return builder.String()
}

// Capitalize optimizes the capitalization for ASCII strings.
// If the string contains only ASCII characters, it avoids converting to []rune.
//
// Example usage:
// result := Capitalize("hello world")
// result will be "Hello world".
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	// Handle ASCII characters efficiently
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-'a'+'A') + s[1:]
	}
	// Fallback to rune conversion for non-ASCII characters
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// HasBlank checks if a string contains any whitespace characters.
// It iterates over each rune in the string and returns true if any whitespace is found.
//
// Example usage:
// hasBlank := HasBlank("Hello World")
// hasBlank will be true.
func HasBlank(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}

// HasNoBlank checks if a string contains no whitespace characters.
// It iterates over each rune in the string and returns true if no whitespace is found.
//
// Example usage:
// hasNoBlank := HasNoBlank("HelloWorld")
// hasNoBlank will be true.
func HasNoBlank(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// HasEmpty checks if any of the provided strings are empty or nil.
// It returns true if any string in the variadic input is empty.
//
// Example usage:
// hasEmpty := HasEmpty("Hello", "", "World")
// hasEmpty will be true.
func HasEmpty(strs ...string) bool {
	for _, str := range strs {
		if len(str) == 0 {
			return true
		}
	}
	return false
}

// HasNoEmpty checks if none of the provided strings are empty or nil.
// It returns true if no string in the variadic input is empty.
//
// Example usage:
// hasNoEmpty := HasNoEmpty("Hello", "World")
// hasNoEmpty will be true.
func HasNoEmpty(strs ...string) bool {
	for _, str := range strs {
		if len(str) == 0 {
			return false
		}
	}
	return true
}

// IsBlank checks if a string is entirely composed of whitespace characters or is empty.
// It returns true if the string contains only whitespace or is empty.
//
// Example usage:
// isBlank := IsBlank("   ")
// isBlank will be true.
func IsBlank(s string) bool {
	for _, r := range s {
		if !unicode.IsSpace(r) {
			return false
		}
	}
	return true
}

// NotBlank checks if a string is not entirely composed of whitespace characters.
// It returns true if the string contains any non-whitespace characters.
//
// Example usage:
// notBlank := NotBlank("Hello")
// notBlank will be true.
func NotBlank(s string) bool {
	return !IsBlank(s)
}

// IsEmpty checks if a string is empty (length 0).
// It returns true if the string is empty.
//
// Example usage:
// isEmpty := IsEmpty("")
// isEmpty will be true.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// NotEmpty checks if a string is not empty (length > 0).
// It returns true if the string is not empty.
//
// Example usage:
// notEmpty := NotEmpty("Hello")
// notEmpty will be true.
func NotEmpty(s string) bool {
	return !IsEmpty(s)
}

// Sub extracts a substring from the input string, supporting both ASCII and Unicode characters efficiently.
// The function handles negative start indices, which count from the end of the string, similar to Python.
// If the start index is negative, it is adjusted to refer to the correct position in the string.
// If the start or length parameters are out of range, an empty string is returned.
//
// The function optimizes for ASCII strings by directly manipulating bytes when possible,
// and falls back to handling Unicode characters using runes (Unicode code points) when needed.
//
// Parameters:
// - s: The input string from which to extract the substring.
// - start: The starting index of the substring. A negative value counts from the end of the string.
// - length: The length of the substring to extract.
//
// Returns:
//   - A substring from the input string starting at 'start' and of 'length' characters.
//     If the indices are invalid, it returns an empty string.
//
// Example usage:
//
//	Sub("abcdefgh", 2, 3)    // Returns "cde"
//	Sub("abcdefgh", -3, 2)   // Returns "fg"
func Sub(s string, start, length int) string {
	// Check if the string is composed entirely of ASCII characters (single byte per character)
	isASCII := true
	for i := 0; i < len(s); i++ {
		if s[i] >= 0x80 { // Characters with a byte value >= 0x80 are non-ASCII
			isASCII = false
			break
		}
	}

	// If the string is ASCII, process it using byte indexing for efficiency
	if isASCII {
		if start < 0 {
			start = len(s) + start // Support for negative index
		}
		end := start + length
		// Return empty string if indices are out of range
		if start < 0 || end > len(s) || start > len(s) {
			return ""
		}
		return s[start:end]
	}

	// For non-ASCII strings, use rune (Unicode code point) indexing
	runes := []rune(s) // Convert string to a slice of runes
	if start < 0 {
		start = len(runes) + start // Support for negative index
	}
	end := start + length
	// Return empty string if indices are out of range
	if start < 0 || end > len(runes) || start > len(runes) {
		return ""
	}
	return string(runes[start:end]) // Convert the slice of runes back to a string
}

// Str converts a byte slice to a string.
// This is a simple wrapper over string conversion.
//
// Example usage:
// str := Str([]byte{72, 101, 108, 108, 111})
// str will be "Hello".
func Str(b []byte) string {
	return string(b)
}

// Bytes converts a string to a byte slice.
// This is a simple wrapper over []byte conversion.
//
// Example usage:
// bytes := Bytes("Hello")
// bytes will be []byte{72, 101, 108, 108, 111}.
func Bytes(s string) []byte {
	return []byte(s)
}

// Reverse reverses the order of characters in a string.
// It converts the string to a slice of runes to handle multi-byte characters correctly.
//
// Example usage:
// reversed := Reverse("Hello")
// reversed will be "olleH".
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// RemovePrefix removes a specified prefix from a string, if it exists.
// If the prefix is not present, the original string is returned.
//
// Example usage:
// result := RemovePrefix("HelloWorld", "Hello")
// result will be "World".
func RemovePrefix(s, prefix string) string {
	return strings.TrimPrefix(s, prefix)
}

// RemoveSuffix removes a specified suffix from a string, if it exists.
// If the suffix is not present, the original string is returned.
//
// Example usage:
// result := RemoveSuffix("HelloWorld", "World")
// result will be "Hello".
func RemoveSuffix(s, suffix string) string {
	return strings.TrimSuffix(s, suffix)
}

// Format replaces placeholders in a template string with provided arguments.
// Placeholders are denoted by "{}". The function replaces each occurrence of "{}"
// with the corresponding argument in the provided order.
//
// Example usage:
// template := "{} loves {}"
// formatted := Format(template, "Alice", "Bob")
// formatted will be "Alice loves Bob".
func Format(template string, args ...interface{}) string {
	result := template
	for _, arg := range args {
		result = strings.Replace(result, "{}", arg.(string), 1)
	}
	return result
}

// PadLeft pads the input string on the left with the specified padChar until it reaches the target length.
// If the string is already longer than or equal to the target length, it returns the original string.
//
// Example usage:
// result := PadLeft("Go", 5, 'x')
// result will be "xxGo".
func PadLeft(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(string(padChar), length-len(s))
	return padding + s
}

// PadRight pads the input string on the right with the specified padChar until it reaches the target length.
// If the string is already longer than or equal to the target length, it returns the original string.
//
// Example usage:
// result := PadRight("Go", 5, 'x')
// result will be "Goxx".
func PadRight(s string, length int, padChar rune) string {
	if len(s) >= length {
		return s
	}
	padding := strings.Repeat(string(padChar), length-len(s))
	return s + padding
}

// JoinNonEmpty joins the non-empty strings in the slice with a specified separator.
// It skips over any empty strings.
//
// Example usage:
// result := JoinNonEmpty([]string{"Hello", "", "World"}, " ")
// result will be "Hello World".
func JoinNonEmpty(strs []string, separator string) string {
	var nonEmpty []string
	for _, str := range strs {
		if str != "" {
			nonEmpty = append(nonEmpty, str)
		}
	}
	return strings.Join(nonEmpty, separator)
}

// RandomString generates a random string of the specified length using characters from a given set.
// By default, it prioritizes performance using non-secure random generation (math/rand).
// If secure is true, it uses crypto/rand for cryptographically secure random numbers.
//
// Parameters:
//   - length: The length of the random string to generate.
//   - charset: The set of characters to choose from for generating the random string.
//   - secure: A boolean indicating whether to use cryptographically secure random number generation.
//     Defaults to false (non-secure, performance-first).
//
// Returns:
// - A random string of the specified length.
// - An error if secure random generation fails (only applicable if secure is true).
func RandomString(length int, charset string, secure ...bool) (string, error) {
	isSecure := false
	if len(secure) > 0 {
		isSecure = secure[0] // Use the value of secure if provided
	}

	var result strings.Builder
	result.Grow(length)

	if isSecure {
		// Secure random generation using crypto/rand
		for i := 0; i < length; i++ {
			num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
			if err != nil {
				return "", err // Return error if secure random generation fails
			}
			result.WriteByte(charset[num.Int64()])
		}
	} else {
		// Non-secure random generation using a new math/rand instance
		r := mathrand.New(mathrand.NewSource(time.Now().UnixNano())) // Create a new random generator
		for i := 0; i < length; i++ {
			result.WriteByte(charset[r.Intn(len(charset))])
		}
	}

	return result.String(), nil
}

// IsNumeric checks if the input string consists only of numeric characters (0-9).
//
// Example usage:
// result := IsNumeric("12345")
// result will be true.
func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// WordWrap wraps a string into lines of a specified width, breaking at spaces when possible.
// If a word exceeds the width, it will be placed on a new line.
//
// Example usage:
// result := WordWrap("This is a very long string", 10)
// result will be:
// "This is a
// very long
// string"
func WordWrap(s string, width int) string {
	var builder strings.Builder
	words := strings.Fields(s)
	lineLength := 0

	for i, word := range words {
		if lineLength+len(word) > width {
			builder.WriteString("\n")
			lineLength = 0
		} else if i > 0 {
			builder.WriteString(" ")
			lineLength++
		}
		builder.WriteString(word)
		lineLength += len(word)
	}
	return builder.String()
}

// ContainsAny checks if the input string contains any of the specified substrings.
// It returns true if any of the substrings are found.
//
// Example usage:
// result := ContainsAny("hello world", []string{"earth", "world"})
// result will be true.
func ContainsAny(s string, substrings []string) bool {
	for _, substr := range substrings {
		if strings.Contains(s, substr) {
			return true
		}
	}
	return false
}

// Pluralize adds an "s" to the input string if the count is not 1. Useful for basic pluralization.
//
// Example usage:
// result := Pluralize("apple", 2)
// result will be "apples".
func Pluralize(word string, count int) string {
	if count == 1 {
		return word
	}
	return word + "s"
}

// StripControlChars removes non-printable characters (control characters) from the input string.
//
// Example usage:
// result := StripControlChars("Hello\x00World")
// result will be "HelloWorld".
func StripControlChars(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if !unicode.IsControl(r) {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// RemoveDuplicates removes duplicate characters from the input string, leaving only the first occurrence of each.
// It returns the deduplicated string.
//
// Example usage:
// result := RemoveDuplicates("aabbcc")
// result will be "abc".
func RemoveDuplicates(s string) string {
	seen := make(map[rune]struct{})
	var builder strings.Builder

	for _, char := range s {
		if _, exists := seen[char]; !exists {
			builder.WriteRune(char)
			seen[char] = struct{}{}
		}
	}

	return builder.String()
}

// Slugify converts a string to a URL-friendly slug by removing non-alphanumeric characters and replacing spaces with hyphens.
//
// Example usage:
// result := Slugify("Hello World!")
// result will be "hello-world".
func Slugify(s string) string {
	// Convert to lowercase and replace spaces with hyphens.
	slug := strings.ToLower(s)
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove non-alphanumeric characters except hyphens.
	re := regexp.MustCompile("[^a-z0-9-]")
	slug = re.ReplaceAllString(slug, "")

	// Trim any leading or trailing hyphens.
	return strings.Trim(slug, "-")
}

// SplitAt splits a string into two parts at the specified index.
//
// Example usage:
// left, right := SplitAt("hello world", 5)
// left will be "hello", right will be " world".
func SplitAt(s string, index int) (string, string) {
	if index < 0 || index > len(s) {
		return s, ""
	}
	return s[:index], s[index:]
}

// RemoveNonAlphaNumeric removes all non-alphanumeric characters from the input string.
//
// Example usage:
// result := RemoveNonAlphaNumeric("a-b_c!d@e#f$g%h")
// result will be "abcdefgh".
func RemoveNonAlphaNumeric(s string) string {
	re := regexp.MustCompile("[^a-zA-Z0-9]")
	return re.ReplaceAllString(s, "")
}

// LevenshteinDistance calculates the Levenshtein distance between two strings,
// which represents the minimum number of single-character edits required to transform one string into the other.
//
// Example usage:
// distance := LevenshteinDistance("kitten", "sitting")
// result will be 3.
func LevenshteinDistance(a, b string) int {
	lenA, lenB := len(a), len(b)
	if lenA == 0 {
		return lenB
	}
	if lenB == 0 {
		return lenA
	}

	matrix := make([][]int, lenA+1)
	for i := range matrix {
		matrix[i] = make([]int, lenB+1)
	}

	for i := 0; i <= lenA; i++ {
		matrix[i][0] = i
	}
	for j := 0; j <= lenB; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= lenA; i++ {
		for j := 1; j <= lenB; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			matrix[i][j] = min(matrix[i-1][j]+1, matrix[i][j-1]+1, matrix[i-1][j-1]+cost)
		}
	}
	return matrix[lenA][lenB]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}

// RegexReplace replaces all occurrences of a pattern in a string with a replacement string.
// It uses the regexp package for pattern matching and replacement.
//
// Example usage:
// result := RegexReplace("Hello 123 World", "\\d+", "#")
// result will be "Hello # World".
func RegexReplace(s, pattern, replacement string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return re.ReplaceAllString(s, replacement), nil
}
