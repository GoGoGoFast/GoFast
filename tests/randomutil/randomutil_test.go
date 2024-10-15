package randomutil_test

import (
	"VeloCore/pkg/util/randomutil"
	"testing"
	"unicode"
)

// 测试 RandomInt
func TestRandomInt(t *testing.T) {
	min, max := int64(1), int64(100)
	for i := 0; i < 100; i++ {
		num := randomutil.RandomInt(min, max)
		if num < min || num >= max {
			t.Errorf("RandomInt() generated out of range number: %d", num)
		}
	}
}

// 测试 RandomBytes
func TestRandomBytes(t *testing.T) {
	length := 16
	bytes := randomutil.RandomBytes(length)
	if len(bytes) != length {
		t.Errorf("RandomBytes() generated byte slice of incorrect length: got %d, want %d", len(bytes), length)
	}
}

// 测试 RandomEle
func TestRandomEle(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	seen := make(map[int]bool)
	for i := 0; i < 100; i++ {
		elem := randomutil.RandomEle(list)
		seen[elem] = true
		if elem < 1 || elem > 5 {
			t.Errorf("RandomEle() returned out of range element: %d", elem)
		}
	}
	if len(seen) < len(list) {
		t.Errorf("RandomEle() did not return all possible elements over 100 iterations")
	}
}

// 测试 RandomEleSet
func TestRandomEleSet(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	set := randomutil.RandomEleSet(list, 3)
	if len(set) != 3 {
		t.Errorf("RandomEleSet() returned incorrect number of elements: got %d, want 3", len(set))
	}

	unique := make(map[int]bool)
	for _, ele := range set {
		unique[ele] = true
		if ele < 1 || ele > 5 {
			t.Errorf("RandomEleSet() returned out of range element: %d", ele)
		}
	}
	if len(unique) != len(set) {
		t.Errorf("RandomEleSet() returned duplicate elements")
	}
}

// 测试 RandomString
func TestRandomString(t *testing.T) {
	length := 10
	str := randomutil.RandomString(length)
	if len(str) != length {
		t.Errorf("RandomString() generated string of incorrect length: got %d, want %d", len(str), length)
	}
	for _, ch := range str {
		if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			t.Errorf("RandomString() generated invalid character: %c", ch)
		}
	}
}

// 测试 RandomNumbers
func TestRandomNumbers(t *testing.T) {
	length := 8
	numStr := randomutil.RandomNumbers(length)
	if len(numStr) != length {
		t.Errorf("RandomNumbers() generated string of incorrect length: got %d, want %d", len(numStr), length)
	}
	for _, ch := range numStr {
		if !unicode.IsDigit(ch) {
			t.Errorf("RandomNumbers() generated invalid character: %c", ch)
		}
	}
}

// 测试 WeightRandom
// 测试 WeightRandom
func TestWeightRandom(t *testing.T) {
	items := []randomutil.Weighted{
		{Item: "A", Weight: 1},
		{Item: "B", Weight: 2},
		{Item: "C", Weight: 3},
	}

	results := map[string]int{"A": 0, "B": 0, "C": 0}
	totalTests := 100000 // 增加测试次数
	for i := 0; i < totalTests; i++ {
		item := randomutil.WeightRandom(items)
		results[item.(string)]++
	}

	// 计算每个元素实际的比例
	expectedRatios := map[string]float64{
		"A": 1.0 / 6.0,
		"B": 2.0 / 6.0,
		"C": 3.0 / 6.0,
	}

	// 允许的误差百分比
	tolerance := 0.01

	for item, expectedRatio := range expectedRatios {
		actualRatio := float64(results[item]) / float64(totalTests)
		if (actualRatio < expectedRatio-tolerance) || (actualRatio > expectedRatio+tolerance) {
			t.Errorf("WeightRandom() distribution incorrect for item %s: got %f, want around %f", item, actualRatio, expectedRatio)
		}
	}
}
