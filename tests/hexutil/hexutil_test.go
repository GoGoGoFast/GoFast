package hexutil_test

import (
	"GoFast/pkg/util/hexutil"
	"testing"
)

func TestEncodeHexStr(t *testing.T) {
	input := "hello"
	expected := "68656c6c6f"
	result, err := hexutil.EncodeHexStr(input)
	if err != nil {
		t.Errorf("EncodeHexStr(%s) returned error: %v", input, err)
	}
	if result != expected {
		t.Errorf("EncodeHexStr(%s) = %s; want %s", input, result, expected)
	}
}

func TestDecodeHexStr(t *testing.T) {
	input := "68656c6c6f"
	expected := "hello"
	result, err := hexutil.DecodeHexStr(input)
	if err != nil {
		t.Errorf("DecodeHexStr(%s) returned error: %v", input, err)
	}
	if result != expected {
		t.Errorf("DecodeHexStr(%s) = %s; want %s", input, result, expected)
	}
}

func TestEncodeBinaryStr(t *testing.T) {
	input := "hello"
	expected := "0110100001100101011011000110110001101111"
	result := hexutil.EncodeBinaryStr(input)
	if result != expected {
		t.Errorf("EncodeBinaryStr(%s) = %s; want %s", input, result, expected)
	}
}

func TestDecodeBinaryStr(t *testing.T) {
	input := "0110100001100101011011000110110001101111"
	expected := "hello"
	result, err := hexutil.DecodeBinaryStr(input)
	if err != nil {
		t.Errorf("DecodeBinaryStr(%s) returned error: %v", input, err)
	}
	if result != expected {
		t.Errorf("DecodeBinaryStr(%s) = %s; want %s", input, result, expected)
	}
}

func TestEncodeOctalStr(t *testing.T) {
	input := "hello"
	expected := "150145154154157"
	result := hexutil.EncodeOctalStr(input)
	if result != expected {
		t.Errorf("EncodeOctalStr(%s) = %s; want %s", input, result, expected)
	}
}

func TestDecodeOctalStr(t *testing.T) {
	input := "150145154154157"
	expected := "hello"
	result, err := hexutil.DecodeOctalStr(input)
	if err != nil {
		t.Errorf("DecodeOctalStr(%s) returned error: %v", input, err)
	}
	if result != expected {
		t.Errorf("DecodeOctalStr(%s) = %s; want %s", input, result, expected)
	}
}
