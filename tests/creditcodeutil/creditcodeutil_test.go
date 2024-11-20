package creditcodeutil_test

import (
	"GoFast/pkg/util/creditcodeutil"
	"testing"
)

// TestValidateCreditCode 测试 ValidateCreditCode 方法
func TestValidateCreditCode(t *testing.T) {
	validCode := "91310110MA1GL3QW03"       // 18位社会信用代码
	invalidCodeLength := "123456789"        // 长度不足
	invalidCodeChar := "91310110MA1GL3QW0@" // 含有非法字符

	tests := []struct {
		code       string
		shouldPass bool
	}{
		{validCode, true},
		{invalidCodeLength, false},
		{invalidCodeChar, false},
	}

	for _, test := range tests {
		result := creditcodeutil.ValidateCreditCode(test.code)
		if result != test.shouldPass {
			t.Errorf("ValidateCreditCode(%s) = %v, 预期 %v", test.code, result, test.shouldPass)
		}
	}
}

// TestRandomCreditCode 测试 RandomCreditCode 方法生成的随机代码
func TestRandomCreditCode(t *testing.T) {
	code := creditcodeutil.RandomCreditCode()

	// 校验生成的社会信用代码是否符合基本要求
	if len(code) != 18 {
		t.Errorf("RandomCreditCode() 返回的代码长度为 %d，预期为 18", len(code))
	}

	// 测试生成的代码是否可以通过 ValidateCreditCode 校验
	if !creditcodeutil.ValidateCreditCode(code) {
		t.Errorf("RandomCreditCode() 返回的代码未能通过校验: %s", code)
	}
}
