package idcardutil_test

import (
	"testing"
	"time"

	"GoAllInOne/pkg/util/idcardutil"
)

func TestIsValidCard(t *testing.T) {
	tests := []struct {
		idCard   string
		expected bool
	}{
		{"11010119900307211X", true},  // 北京市 东城区
		{"110101199003072110", false}, // 校验位错误
		{"31010119900307211X", true},  // 上海市 黄浦区
		{"32031119770706001X", true},  // 江苏省 常州市
		{"32031119770706001Z", false}, // 校验位错误
		{"320311770706001", false},    // 长度不足
		{"320311770706001234", false}, // 长度多余
	}

	for _, test := range tests {
		if result := idcardutil.IsValidCard(test.idCard); result != test.expected {
			t.Errorf("IsValidCard(%v) = %v; expected %v", test.idCard, result, test.expected)
		}
	}
}

func TestGetBirthByIdCard(t *testing.T) {
	tests := []struct {
		idCard   string
		expected string
		hasError bool
	}{
		{"11010119900307211X", "19900307", false},
		{"31010119900307211X", "19900307", false},
		{"32031119770706001X", "19770706", false},
		{"11010119900307211", "", true},
	}

	for _, test := range tests {
		birth, err := idcardutil.GetBirthByIdCard(test.idCard)
		if (err != nil) != test.hasError || birth != test.expected {
			t.Errorf("GetBirthByIdCard(%v) = %v, %v; expected %v, %v", test.idCard, birth, err, test.expected, test.hasError)
		}
	}
}

func TestGetAgeByIdCard(t *testing.T) {
	currentDate := time.Date(2024, 10, 15, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		idCard   string
		expected int
		hasError bool
	}{
		{"11010119900307211X", 34, false},
		{"31010119900307211X", 34, false},
		{"32031119770706001X", 47, false},
		{"11010119900307211", 0, true},
	}

	for _, test := range tests {
		age, err := idcardutil.GetAgeByIdCard(test.idCard, currentDate)
		if (err != nil) != test.hasError || age != test.expected {
			t.Errorf("GetAgeByIdCard(%v, %v) = %v, %v; expected %v, %v", test.idCard, currentDate, age, err, test.expected, test.hasError)
		}
	}
}

func TestGetGenderByIdCard(t *testing.T) {
	tests := []struct {
		idCard   string
		expected string
		hasError bool
	}{
		{"11010119900307211X", "男", false},
		{"31010119900307211X", "男", false},
		{"32031119770706001X", "男", false},
		{"110101199003072120", "女", false},
		{"11010119900307211", "", true},
	}

	for _, test := range tests {
		gender, err := idcardutil.GetGenderByIdCard(test.idCard)
		if (err != nil) != test.hasError || gender != test.expected {
			t.Errorf("GetGenderByIdCard(%v) = %v, %v; expected %v, %v", test.idCard, gender, err, test.expected, test.hasError)
		}
	}
}

func TestGetRegionByIdCard(t *testing.T) {

	tests := []struct {
		idCard   string
		expected string
		hasError bool
	}{
		{"11010119900307211X", "北京市 市辖区 东城区", false},
		{"31010119900307211X", "上海市 市辖区 黄浦区", false},
		{"32031119770706001X", "江苏省 常州市 新区", false},
		{"110101199003072110", "北京市 市辖区", false},
		{"120000199003072110", "未知", false},
		{"11010119900307211", "", true},
	}

	for _, test := range tests {
		region, err := idcardutil.GetRegionByIdCard(test.idCard)
		if (err != nil) != test.hasError || region != test.expected {
			t.Errorf("GetRegionByIdCard(%v) = %v, %v; expected %v, %v", test.idCard, region, err, test.expected, test.hasError)
		}
	}
}
