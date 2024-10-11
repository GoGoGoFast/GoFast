package idcardutil

import (
	"errors"
	"regexp"
	"time"
)

var provinceCodes = map[string]string{
	"11": "北京",
	"12": "天津",
	"13": "河北",
	"14": "山西",
	"15": "内蒙古",
	"21": "辽宁",
	"22": "吉林",
	"23": "黑龙江",
	"31": "上海",
	"32": "江苏",
	"33": "浙江",
	"34": "安徽",
	"35": "福建",
	"36": "江西",
	"37": "山东",
	"41": "河南",
	"42": "湖北",
	"43": "湖南",
	"44": "广东",
	"45": "广西",
	"46": "海南",
	"50": "重庆",
	"51": "四川",
	"52": "贵州",
	"53": "云南",
	"54": "西藏",
	"61": "陕西",
	"62": "甘肃",
	"63": "青海",
	"64": "宁夏",
	"65": "新疆",
	"71": "台湾",
	"81": "香港",
	"82": "澳门",
}

var idCard15Pattern = regexp.MustCompile(`^\d{15}$`)
var idCard18Pattern = regexp.MustCompile(`^\d{17}(\d|X)$`)

// IsValidCard 验证身份证是否合法
func IsValidCard(idCard string) bool {
	if len(idCard) == 15 {
		return idCard15Pattern.MatchString(idCard)
	} else if len(idCard) == 18 {
		return idCard18Pattern.MatchString(idCard) && validate18IDCard(idCard)
	}
	return false
}

// 验证18位身份证
func validate18IDCard(idCard string) bool {
	// 校验位计算
	power := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checksum := 0
	for i := 0; i < 17; i++ {
		num := int(idCard[i] - '0')
		checksum += num * power[i]
	}
	checkDigits := "10X98765432"
	checkDigit := checkDigits[checksum%11]

	return checkDigit == idCard[17]
}

// 计算校验位
func calculateChecksum(idCard string) byte {
	power := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	checksum := 0
	for i := 0; i < 17; i++ {
		num := int(idCard[i] - '0')
		checksum += num * power[i]
	}
	checkDigits := "10X98765432"
	return checkDigits[checksum%11]
}

// GetBirthByIdCard 获取生日
func GetBirthByIdCard(idCard string) (string, error) {
	if !IsValidCard(idCard) {
		return "", errors.New("无效的身份证号码")
	}

	if len(idCard) == 15 {
		return "19" + idCard[6:12], nil
	}
	return idCard[6:14], nil
}

// GetAgeByIdCard 获取年龄
func GetAgeByIdCard(idCard string, currentDate time.Time) (int, error) {
	birthStr, err := GetBirthByIdCard(idCard)
	if err != nil {
		return 0, err
	}

	birthDate, err := time.Parse("20060102", birthStr)
	if err != nil {
		return 0, err
	}

	age := currentDate.Year() - birthDate.Year()
	if currentDate.YearDay() < birthDate.YearDay() {
		age--
	}
	return age, nil
}

// GetGenderByIdCard 获取性别
func GetGenderByIdCard(idCard string) (string, error) {
	if !IsValidCard(idCard) {
		return "", errors.New("无效的身份证号码")
	}

	if len(idCard) == 15 {
		return "未知", nil
	}

	if idCard[16]%2 == '1' {
		return "男", nil
	}
	return "女", nil
}

// GetProvinceByIdCard 获取省份
func GetProvinceByIdCard(idCard string) (string, error) {
	if !IsValidCard(idCard) {
		return "", errors.New("无效的身份证号码")
	}

	provinceCode := idCard[:2]
	if province, exists := provinceCodes[provinceCode]; exists {
		return province, nil
	}
	return "未知", nil
}
