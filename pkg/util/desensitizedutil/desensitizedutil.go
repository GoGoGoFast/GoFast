package desensitizedutil

import (
	"strings"
)

// DesensitizedUtil 脱敏工具类
type DesensitizedUtil struct{}

// Hide 根据指定的起始和结束位置隐藏字符串中的部分字符
func Hide(str string, start, end int) string {
	if start < 0 || end < 0 || start+end >= len(str) {
		return str
	}
	return str[:start] + strings.Repeat("*", len(str)-start-end) + str[len(str)-end:]
}

// IdCardNum 脱敏身份证号码
func IdCardNum(idCard string, start, end int) string {
	return Hide(idCard, start, end)
}

// MobilePhone 脱敏手机号
func MobilePhone(mobile string) string {
	return Hide(mobile, 3, 4)
}

// Password 脱敏密码
func Password(password string) string {
	return strings.Repeat("*", len(password))
}

// Email 脱敏电子邮件
func Email(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return email // 如果格式不正确，则返回原始邮箱
	}
	local := parts[0]
	domain := parts[1]
	if len(local) <= 2 {
		return "***@" + domain
	}
	return local[:1] + strings.Repeat("*", len(local)-2) + local[len(local)-1:] + "@" + domain
}

// UserId 脱敏用户ID
func UserId(userId string) string {
	return Hide(userId, 1, 1)
}

// Address 脱敏地址
func Address(address string) string {
	parts := strings.Split(address, " ")
	if len(parts) > 1 {
		return Hide(parts[0], 2, 2) + " " + strings.Join(parts[1:], " ")
	}
	return Hide(address, 2, 2)
}

// PlateNumber 脱敏车牌号
func PlateNumber(plate string) string {
	if len(plate) < 6 {
		return plate // 如果车牌号不完整，返回原始车牌
	}
	return plate[:1] + "****" + plate[5:]
}

// BankCard 脱敏银行卡号
func BankCard(cardNumber string) string {
	if len(cardNumber) < 8 {
		return cardNumber // 如果银行卡号不完整，返回原始银行卡号
	}
	return Hide(cardNumber, 4, 4)
}
