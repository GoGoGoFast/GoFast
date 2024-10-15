// Package hexutil 提供了一些实用工具函数，用于将字符串编码为十六进制表示，并将十六进制字符串解码为原始字符串。
// The util package provides utility functions for encoding strings into their hexadecimal representation and decoding hexadecimal strings back into their original string form.
package hexutil

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"strconv"
)

// EncodeHexStr 将字符串编码为指定编码的十六进制表示。如果未指定编码，则默认为 UTF-8。
// EncodeHexStr encodes a string into its hexadecimal representation using the specified encoding. If no encoding is specified, it defaults to UTF-8.
//
// 参数 (Parameters):
// - str: 需要编码的字符串 (The string to be encoded).
// - enc: 可选参数，指定编码方式 (Optional parameter specifying the encoding).
//
// 返回值 (Returns):
// - string: 编码后的十六进制字符串 (The encoded hexadecimal string).
// - error: 如果编码过程中出现错误，返回错误信息 (An error if any occurs during encoding).
func EncodeHexStr(str string, enc ...encoding.Encoding) (string, error) {
	var encoder *encoding.Encoder
	if len(enc) > 0 {
		encoder = enc[0].NewEncoder()
	} else {
		encoder = unicode.UTF8.NewEncoder()
	}

	encodedBytes, _, err := transform.Bytes(encoder, []byte(str))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(encodedBytes), nil
}

// DecodeHexStr 将十六进制字符串解码为其原始字符串表示，使用指定的编码。如果未指定编码，则默认为 UTF-8。
// 如果输入不是有效的十六进制字符串，则返回错误。
// DecodeHexStr decodes a hexadecimal string into its original string representation using the specified encoding. If no encoding is specified, it defaults to UTF-8.
// Returns an error if the input is not a valid hexadecimal string.
//
// 参数 (Parameters):
// - hexStr: 需要解码的十六进制字符串 (The hexadecimal string to be decoded).
// - enc: 可选参数，指定解码方式 (Optional parameter specifying the encoding).
//
// 返回值 (Returns):
// - string: 解码后的原始字符串 (The decoded original string).
// - error: 如果解码过程中出现错误，返回错误信息 (An error if any occurs during decoding).
func DecodeHexStr(hexStr string, enc ...encoding.Encoding) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}

	var decoder *encoding.Decoder
	if len(enc) > 0 {
		decoder = enc[0].NewDecoder()
	} else {
		decoder = unicode.UTF8.NewDecoder()
	}

	decodedBytes, _, err := transform.Bytes(decoder, bytes)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

// EncodeBinaryStr 将字符串编码为二进制表示。
// EncodeBinaryStr encodes a string into its binary representation.
func EncodeBinaryStr(str string) string {
	var binaryStr string
	for _, c := range []byte(str) {
		binaryStr += fmt.Sprintf("%08b", c)
	}
	return binaryStr
}

// DecodeBinaryStr 将二进制字符串解码为原始字符串表示。
// DecodeBinaryStr decodes a binary string into its original string representation.
func DecodeBinaryStr(binaryStr string) (string, error) {
	var bytes []byte
	for i := 0; i < len(binaryStr); i += 8 {
		byteVal, err := strconv.ParseUint(binaryStr[i:i+8], 2, 8)
		if err != nil {
			return "", err
		}
		bytes = append(bytes, byte(byteVal))
	}
	return string(bytes), nil
}

// EncodeOctalStr 将字符串编码为八进制表示。
// EncodeOctalStr encodes a string into its octal representation.
func EncodeOctalStr(str string) string {
	var octalStr string
	for _, c := range []byte(str) {
		octalStr += fmt.Sprintf("%03o", c)
	}
	return octalStr
}

// DecodeOctalStr 将八进制字符串解码为原始字符串表示。
// DecodeOctalStr decodes an octal string into its original string representation.
func DecodeOctalStr(octalStr string) (string, error) {
	var bytes []byte
	for i := 0; i < len(octalStr); i += 3 {
		byteVal, err := strconv.ParseUint(octalStr[i:i+3], 8, 8)
		if err != nil {
			return "", err
		}
		bytes = append(bytes, byte(byteVal))
	}
	return string(bytes), nil
}
