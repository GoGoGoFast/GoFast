package reutil

import (
	"fmt"
	"regexp"
	"strings"
)

// Get 获取匹配的字符串
func Get(regex, content string, groupIndex int) (string, error) {
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return "", fmt.Errorf("正则编译失败: %v", err)
	}

	matches := pattern.FindStringSubmatch(content)
	if len(matches) > groupIndex {
		return matches[groupIndex], nil
	}
	return "", nil
}

// ExtractMulti 抽取多个分组并拼接
func ExtractMulti(regex, content string, template string) (string, error) {
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return "", fmt.Errorf("正则编译失败: %v", err)
	}

	matches := pattern.FindStringSubmatch(content)
	if len(matches) == 0 {
		return "", nil
	}

	result := template
	for i, match := range matches[1:] { // 第一个元素是完整匹配
		result = strings.ReplaceAll(result, fmt.Sprintf("$%d", i+1), match)
	}
	return result, nil
}

// DelFirst 删除第一个匹配到的内容
func DelFirst(regex, content string) (string, error) {
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return "", fmt.Errorf("正则编译失败: %v", err)
	}

	return pattern.ReplaceAllString(content, ""), nil
}

// FindAll 查找所有匹配文本
func FindAll(regex, content string) ([]string, error) {
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf("正则编译失败: %v", err)
	}

	matches := pattern.FindAllString(content, -1)
	return matches, nil
}

// GetFirstNumber 找到第一个匹配的数字
func GetFirstNumber(content string) (int, error) {
	pattern := regexp.MustCompile(`\d+`)
	matches := pattern.FindString(content)
	if matches == "" {
		return 0, fmt.Errorf("未找到匹配的数字")
	}
	var number int
	_, err := fmt.Sscanf(matches, "%d", &number)
	return number, err
}

// IsMatch 判断字符串是否匹配给定正则
func IsMatch(regex, content string) (bool, error) {
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return false, fmt.Errorf("正则编译失败: %v", err)
	}
	return pattern.MatchString(content), nil
}

// ReplaceAll 替换所有匹配文本
func ReplaceAll(content, regex, replacementTemplate string) (string, error) {
	pattern, err := regexp.Compile(regex)
	if err != nil {
		return "", fmt.Errorf("正则编译失败: %v", err)
	}

	result := pattern.ReplaceAllStringFunc(content, func(match string) string {
		// 使用分组替换
		matches := pattern.FindStringSubmatch(match)
		for i := 1; i < len(matches); i++ {
			replacementTemplate = strings.ReplaceAll(replacementTemplate, fmt.Sprintf("$%d", i), matches[i])
		}
		return replacementTemplate
	})
	return result, nil
}

// Escape 转义给定字符串
func Escape(input string) string {
	return regexp.QuoteMeta(input)
}
