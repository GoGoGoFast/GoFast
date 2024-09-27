package randomutil

import (
	"crypto/rand"
	"math/big"
)

// RandomInt 生成一个 [min, max) 范围内的安全随机整数
func RandomInt(min, max int64) int64 {
	n, _ := rand.Int(rand.Reader, big.NewInt(max-min))
	return n.Int64() + min
}

// RandomBytes 生成指定长度的安全随机字节数组
func RandomBytes(length int) []byte {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

// RandomEle 从切片中随机返回一个元素
func RandomEle[T any](list []T) T {
	idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(list))))
	return list[idx.Int64()]
}

// RandomEleSet 从切片中随机返回不重复的元素集，返回一个集合
func RandomEleSet[T comparable](list []T, num int) []T {
	if num > len(list) {
		num = len(list)
	}
	set := make(map[T]struct{})
	result := make([]T, 0, num)
	for len(result) < num {
		ele := RandomEle(list)
		if _, found := set[ele]; !found {
			set[ele] = struct{}{}
			result = append(result, ele)
		}
	}
	return result
}

// RandomString 生成指定长度的随机字符串（包含数字和字母）
func RandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[idx.Int64()]
	}
	return string(result)
}

// RandomNumbers 生成指定长度的随机数字字符串
func RandomNumbers(length int) string {
	const digits = "0123456789"
	result := make([]byte, length)
	for i := range result {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		result[i] = digits[idx.Int64()]
	}
	return string(result)
}

// Weighted 权重随机生成器结构
type Weighted struct {
	Item   interface{}
	Weight int64
}

// WeightRandom 权重随机生成器，根据权重随机获取对象
func WeightRandom(items []Weighted) interface{} {
	totalWeight := int64(0)
	for _, item := range items {
		totalWeight += item.Weight
	}

	randomWeight, _ := rand.Int(rand.Reader, big.NewInt(totalWeight))
	for _, item := range items {
		randomWeight = randomWeight.Sub(randomWeight, big.NewInt(item.Weight))
		if randomWeight.Sign() < 0 {
			return item.Item
		}
	}
	return nil
}
