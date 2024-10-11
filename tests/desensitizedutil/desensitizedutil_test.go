package desensitizedutil_test

import (
	"VeloCore/pkg/util/desensitizedutil"
	"testing"
)

// TestHide 测试隐藏字符的逻辑
func TestHide(t *testing.T) {
	tests := []struct {
		input    string
		start    int
		end      int
		expected string
	}{
		{"short", 1, 1, "s***t"},
		{"longerstring", 3, 3, "lon******ing"},
		{"short", 0, 0, "*****"},
		{"toolong", 5, 5, "toolong"},
	}

	for _, tt := range tests {
		result := desensitizedutil.Hide(tt.input, tt.start, tt.end)
		if result != tt.expected {
			t.Errorf("Hide(%v, %v, %v) = %v, 预期 %v", tt.input, tt.start, tt.end, result, tt.expected)
		}
	}
}

// TestIdCardNum 测试身份证脱敏
func TestIdCardNum(t *testing.T) {
	result := desensitizedutil.IdCardNum("51343620000320711X")
	expected := "5***************1X"
	if result != expected {
		t.Errorf("IdCardNum() = %v, 预期 %v", result, expected)
	}
}

// TestMobilePhone 测试手机号脱敏
func TestMobilePhone(t *testing.T) {
	result := desensitizedutil.MobilePhone("18049531999")
	expected := "180****1999"
	if result != expected {
		t.Errorf("MobilePhone() = %v, 预期 %v", result, expected)
	}
}

// TestPassword 测试密码脱敏
func TestPassword(t *testing.T) {
	result := desensitizedutil.Password("1234567890")
	expected := "**********"
	if result != expected {
		t.Errorf("Password() = %v, 预期 %v", result, expected)
	}
}

// TestEmail 测试电子邮件脱敏
func TestEmail(t *testing.T) {
	result := desensitizedutil.Email("john.doe@example.com")
	expected := "j*******@example.com"
	if result != expected {
		t.Errorf("Email() = %v, 预期 %v", result, expected)
	}
}

// TestUserId 测试用户ID脱敏
func TestUserId(t *testing.T) {
	result := desensitizedutil.UserId("user12345")
	expected := "u*******5"
	if result != expected {
		t.Errorf("UserId() = %v, 预期 %v", result, expected)
	}
}

// TestAddress 测试地址脱敏
func TestAddress(t *testing.T) {
	// 测试中文短地址
	addressCnShort := "上海市"
	expectedCnShort := "上海**"
	resultCnShort := desensitizedutil.Address(addressCnShort)
	if resultCnShort != expectedCnShort {
		t.Errorf("Address() = %v, 预期 %v", resultCnShort, expectedCnShort)
	}

	// 测试中文正常长度地址
	addressCn := "北京市朝阳区"
	expectedCn := "北京****"
	resultCn := desensitizedutil.Address(addressCn)
	if resultCn != expectedCn {
		t.Errorf("Address() = %v, 预期 %v", resultCn, expectedCn)
	}

	// 测试英文短地址
	addressEnShort := "London"
	expectedEnShort := "Lond**"
	resultEnShort := desensitizedutil.Address(addressEnShort)
	if resultEnShort != expectedEnShort {
		t.Errorf("Address() = %v, 预期 %v", resultEnShort, expectedEnShort)
	}

	// 测试英文正常长度地址
	addressEn := "1234 Elm Street, Springfield, IL, USA"
	expectedEn := "****,  Springfield, IL, USA"
	resultEn := desensitizedutil.Address(addressEn)
	if resultEn != expectedEn {
		t.Errorf("Address() = %v, 预期 %v", resultEn, expectedEn)
	}
}

// TestPlateNumber 测试车牌号脱敏
func TestPlateNumber(t *testing.T) {
	result := desensitizedutil.PlateNumber("京A12345")
	expected := "京****45"
	if result != expected {
		t.Errorf("PlateNumber() = %v, 预期 %v", result, expected)
	}
}

// TestBankCard 测试银行卡号脱敏
func TestBankCard(t *testing.T) {
	result := desensitizedutil.BankCard("1234567890123456")
	expected := "1234********3456"
	if result != expected {
		t.Errorf("BankCard() = %v, 预期 %v", result, expected)
	}
}
