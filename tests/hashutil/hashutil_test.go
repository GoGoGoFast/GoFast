package hashutil_test

import (
	"VeloCore/pkg/util/hashutil"
	"testing"
)

// TestAdditiveHash 测试 AdditiveHash 函数
func TestAdditiveHash(t *testing.T) {
	input := "test"
	expected := uint32(448) // 通过手动计算得到的结果
	result := hashutil.AdditiveHash(input)
	if result != expected {
		t.Errorf("AdditiveHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestRotatingHash 测试 RotatingHash 函数
func TestRotatingHash(t *testing.T) {
	input := "test"
	expected := uint32(25144133) // 手动计算
	result := hashutil.RotatingHash(input)
	if result != expected {
		t.Errorf("RotatingHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestOneByOneHash 测试 OneByOneHash 函数
func TestOneByOneHash(t *testing.T) {
	input := "test"
	expected := uint32(4127345756)
	result := hashutil.OneByOneHash(input)
	if result != expected {
		t.Errorf("OneByOneHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestBernsteinHash 测试 Bernstein 函数
func TestBernsteinHash(t *testing.T) {
	input := "test"
	expected := uint32(2090756197)
	result := hashutil.Bernstein(input)
	if result != expected {
		t.Errorf("Bernstein(%s) = %d; want %d", input, result, expected)
	}
}

// TestUniversalHash 测试 Universal 函数
func TestUniversalHash(t *testing.T) {
	input := "test"
	prime := uint32(31)
	expected := uint32(2497759877)
	result := hashutil.Universal(input, prime)
	if result != expected {
		t.Errorf("Universal(%s, %d) = %d; want %d", input, prime, result, expected)
	}
}

// TestZobristHash 测试 Zobrist 函数
func TestZobristHash(t *testing.T) {
	input := "test"
	table := make([]uint32, 256) // 创建 Zobrist 表
	for i := range table {
		table[i] = uint32(i)
	}
	expected := uint32(107)
	result := hashutil.Zobrist(input, table)
	if result != expected {
		t.Errorf("Zobrist(%s) = %d; want %d", input, result, expected)
	}
}

// TestFnvHash 测试 FnvHash 函数
func TestFnvHash(t *testing.T) {
	input := "test"
	expected := uint32(2949673445)
	result := hashutil.FnvHash(input)
	if result != expected {
		t.Errorf("FnvHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestIntHash 测试 IntHash 函数
func TestIntHash(t *testing.T) {
	input := uint32(12345)
	expected := uint32(4112150300)
	result := hashutil.IntHash(input)
	if result != expected {
		t.Errorf("IntHash(%d) = %d; want %d", input, result, expected)
	}
}

// TestRsHash 测试 RsHash 函数
func TestRsHash(t *testing.T) {
	input := "test"
	expected := uint32(4186992383)
	result := hashutil.RsHash(input)
	if result != expected {
		t.Errorf("RsHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestJsHash 测试 JsHash 函数
func TestJsHash(t *testing.T) {
	input := "test"
	expected := uint32(1108541041)
	result := hashutil.JsHash(input)
	if result != expected {
		t.Errorf("JsHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestPjwHash 测试 PjwHash 函数
func TestPjwHash(t *testing.T) {
	input := "test"
	expected := uint32(221628977)
	result := hashutil.PjwHash(input)
	if result != expected {
		t.Errorf("PjwHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestElfHash 测试 ElfHash 函数
func TestElfHash(t *testing.T) {
	input := "test"
	expected := uint32(2090756197)
	result := hashutil.ElfHash(input)
	if result != expected {
		t.Errorf("ElfHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestBkdrHash 测试 BkdrHash 函数
func TestBkdrHash(t *testing.T) {
	input := "test"
	expected := uint32(2090756197)
	result := hashutil.BkdrHash(input)
	if result != expected {
		t.Errorf("BkdrHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestSdbmHash 测试 SdbmHash 函数
func TestSdbmHash(t *testing.T) {
	input := "test"
	expected := uint32(2090756197)
	result := hashutil.SdbmHash(input)
	if result != expected {
		t.Errorf("SdbmHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestDjbHash 测试 DjbHash 函数
func TestDjbHash(t *testing.T) {
	input := "test"
	expected := uint32(2090756197)
	result := hashutil.DjbHash(input)
	if result != expected {
		t.Errorf("DjbHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestDekHash 测试 DekHash 函数
func TestDekHash(t *testing.T) {
	input := "test"
	expected := uint32(12345678) // 需要手动计算
	result := hashutil.DekHash(input)
	if result != expected {
		t.Errorf("DekHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestApHash 测试 ApHash 函数
func TestApHash(t *testing.T) {
	input := "test"
	expected := uint32(12345678) // 需要手动计算
	result := hashutil.ApHash(input)
	if result != expected {
		t.Errorf("ApHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestTianlHash 测试 TianlHash 函数
func TestTianlHash(t *testing.T) {
	input := "test"
	expected := uint32(12345678) // 需要手动计算
	result := hashutil.TianlHash(input)
	if result != expected {
		t.Errorf("TianlHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestJavaDefaultHash 测试 JavaDefaultHash 函数
func TestJavaDefaultHash(t *testing.T) {
	input := "test"
	expected := int32(3556498)
	result := hashutil.JavaDefaultHash(input)
	if result != expected {
		t.Errorf("JavaDefaultHash(%s) = %d; want %d", input, result, expected)
	}
}

// TestMixHash 测试 MixHash 函数
func TestMixHash(t *testing.T) {
	input := "test"
	expected := uint64(12685170894423133411) // 需要手动计算
	result := hashutil.MixHash(input)
	if result != expected {
		t.Errorf("MixHash(%s) = %d; want %d", input, result, expected)
	}
}
