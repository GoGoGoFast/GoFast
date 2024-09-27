package creditcodeutil

import (
	"math/rand"
	"strings"
	"time"
)

// CreditCodeUtil 社会信用代码工具类
type CreditCodeUtil struct{}

// isValidChar 校验字符是否符合要求
func isValidChar(char rune) bool {
	return (char >= '0' && char <= '9') || (char >= 'A' && char <= 'Z')
}

// ValidateCreditCode 校验社会信用代码
func ValidateCreditCode(code string) bool {
	if len(code) != 18 {
		return false
	}

	// 校验每个字符是否有效
	for _, char := range code {
		if !isValidChar(char) {
			return false
		}
	}

	// 计算校验位
	// 这里可以根据需要实现具体的校验逻辑
	// 当前实现仅返回true，需替换为实际的校验逻辑
	return true
}

// randomChar 生成随机字符
func randomChar() rune {
	if rand.Intn(2) == 0 {
		return rune(rand.Intn(10) + '0') // 生成数字
	}
	return rune(rand.Intn(26) + 'A') // 生成大写字母
}

// RandomCreditCode 生成随机社会信用代码
func RandomCreditCode() string {
	var sb strings.Builder
	sb.WriteRune(randomChar()) // 第一部分：登记管理部门代码
	sb.WriteRune(randomChar()) // 第二部分：机构类别代码
	for i := 0; i < 6; i++ {
		sb.WriteRune(rune(rand.Intn(10) + '0')) // 第三部分：登记管理机关行政区划码
	}
	for i := 0; i < 9; i++ {
		sb.WriteRune(randomChar()) // 第四部分：主体标识码
	}
	sb.WriteRune(randomChar()) // 第五部分：校验码

	return sb.String()
}

func init() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子
}
