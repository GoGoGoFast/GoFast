package desensitizedutil

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// DesensitizedUtil 脱敏工具类
type DesensitizedUtil struct{}

// Hide 根据指定的起始和结束位置隐藏字符串中的部分字符，处理 Unicode 字符
func Hide(str string, start, end int) string {
	runeStr := []rune(str) // 将字符串转为 rune 切片，处理 Unicode 字符
	strLen := len(runeStr)

	if start < 0 || end < 0 || start+end >= strLen {
		return str
	}

	// 生成隐藏的部分
	hiddenPart := strings.Repeat("*", strLen-start-end)

	// 返回脱敏结果
	return string(runeStr[:start]) + hiddenPart + string(runeStr[strLen-end:])
}

// IdCardNum 脱敏身份证号码，保留首位和最后两位
func IdCardNum(idCard string) string {
	return Hide(idCard, 1, 2)
}

// MobilePhone 脱敏手机号，保留前三位和最后四位
func MobilePhone(mobile string) string {
	return Hide(mobile, 3, 4)
}

// Password 脱敏密码，隐藏全部字符
func Password(password string) string {
	return strings.Repeat("*", len(password))
}

// Email 脱敏电子邮件，保留首字母和@之后的域名
func Email(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email
	}
	local := parts[0]
	domain := parts[1]
	if len(local) <= 1 {
		return "*" + "@" + domain
	}
	return local[:1] + strings.Repeat("*", len(local)-1) + "@" + domain
}

// UserId 脱敏用户ID，保留首位和最后一位
func UserId(userId string) string {
	return Hide(userId, 1, 1)
}

// Address 脱敏地址
// 中文地址保留前2个字符，英文地址保留前4个字符，处理短地址和标准地址
func Address(address string) string {
	isChinese, _ := regexp.MatchString("[\u4e00-\u9fa5]", address)
	if isChinese {
		return desensitizeChineseAddress(address)
	}
	return desensitizeEnglishAddress(address)
}

// desensitizeChineseAddress 脱敏中文地址，适应短地址的处理
func desensitizeChineseAddress(address string) string {
	runeAddress := []rune(address)
	addressLen := len(runeAddress)

	if addressLen <= 3 {
		// 地址太短，保留前2个字符，并且在后面补上两个星号
		return string(runeAddress[:2]) + "**"
	}

	// 地址较长，保留前2个字符，并隐藏后续部分
	return string(runeAddress[:2]) + strings.Repeat("*", addressLen-2)
}

// desensitizeEnglishAddress 脱敏英文地址，适应短地址的处理
func desensitizeEnglishAddress(address string) string {
	runeAddress := []rune(address)

	if len(runeAddress) <= 4 {
		// 如果地址较短，保留前4个字符，隐藏后续部分
		return Hide(address, 4, 0)
	}

	// 否则隐藏街道号和街道名，保留城市、州/省和国家
	parts := strings.Split(address, ",")
	if len(parts) == 1 {
		return Hide(address, 4, 0) // 保留前4个字符
	}
	return "****, " + strings.Join(parts[1:], ",")
}

// PlateNumber 脱敏车牌号，处理多字节字符
func PlateNumber(plate string) string {
	runePlate := []rune(plate)
	if len(runePlate) < 6 {
		return plate // 如果车牌号不完整，返回原始车牌
	}
	return string(runePlate[:1]) + "****" + string(runePlate[len(runePlate)-2:])
}

// BankCard 脱敏银行卡号，保留前四位和最后四位
func BankCard(cardNumber string) string {
	if utf8.RuneCountInString(cardNumber) < 8 {
		return cardNumber
	}
	return Hide(cardNumber, 4, 4)
}
