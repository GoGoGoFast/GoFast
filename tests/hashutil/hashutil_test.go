package hashutil_test

import (
	"testing"

	"GoFast/pkg/util/hashutil"
	"github.com/stretchr/testify/assert"
)

func TestAdditiveHash(t *testing.T) {
	assert.Equal(t, hashutil.AdditiveHash("test"), uint32(448))
	assert.Equal(t, hashutil.AdditiveHash("hello"), uint32(532))
	assert.Equal(t, hashutil.AdditiveHash("world"), uint32(552))
}

func TestRotatingHash(t *testing.T) {
	assert.Equal(t, hashutil.RotatingHash("test"), uint32(1633771873))
	assert.Equal(t, hashutil.RotatingHash("hello"), uint32(1546672534))
	assert.Equal(t, hashutil.RotatingHash("world"), uint32(1157136195))
}

func TestOneByOneHash(t *testing.T) {
	assert.Equal(t, hashutil.OneByOneHash("test"), uint32(1092158949))
	assert.Equal(t, hashutil.OneByOneHash("hello"), uint32(1084105767))
	assert.Equal(t, hashutil.OneByOneHash("world"), uint32(1258415024))
}

func TestBernstein(t *testing.T) {
	assert.Equal(t, hashutil.Bernstein("test"), uint32(2090756197))
	assert.Equal(t, hashutil.Bernstein("hello"), uint32(99162322))
	assert.Equal(t, hashutil.Bernstein("world"), uint32(113318802))
}

func TestUniversal(t *testing.T) {
	prime := uint32(31)
	assert.Equal(t, hashutil.Universal("test", prime), uint32(2768031985))
	assert.Equal(t, hashutil.Universal("hello", prime), uint32(1942563810))
	assert.Equal(t, hashutil.Universal("world", prime), uint32(2248119733))
}

func TestZobrist(t *testing.T) {
	table := make([]uint32, 256)
	for i := range table {
		table[i] = uint32(i)
	}
	assert.Equal(t, hashutil.Zobrist("test", table), uint32(104))
	assert.Equal(t, hashutil.Zobrist("hello", table), uint32(106))
	assert.Equal(t, hashutil.Zobrist("world", table), uint32(111))
}

func TestFnvHash(t *testing.T) {
	assert.Equal(t, hashutil.FnvHash("test"), uint32(2949673445))
	assert.Equal(t, hashutil.FnvHash("hello"), uint32(1335831723))
	assert.Equal(t, hashutil.FnvHash("world"), uint32(2682274536))
}

func TestIntHash(t *testing.T) {
	assert.Equal(t, hashutil.IntHash(12345678), uint32(3650451927))
	assert.Equal(t, hashutil.IntHash(87654321), uint32(3070069356))
	assert.Equal(t, hashutil.IntHash(11223344), uint32(3534259454))
}

func TestRsHash(t *testing.T) {
	assert.Equal(t, hashutil.RsHash("test"), uint32(3649271295))
	assert.Equal(t, hashutil.RsHash("hello"), uint32(2550737373))
	assert.Equal(t, hashutil.RsHash("world"), uint32(2550737373))
}

func TestJsHash(t *testing.T) {
	assert.Equal(t, hashutil.JsHash("test"), uint32(1137617854))
	assert.Equal(t, hashutil.JsHash("hello"), uint32(1097752761))
	assert.Equal(t, hashutil.JsHash("world"), uint32(1929241140))
}

func TestPjwHash(t *testing.T) {
	assert.Equal(t, hashutil.PjwHash("test"), uint32(118224))
	assert.Equal(t, hashutil.PjwHash("hello"), uint32(117838))
	assert.Equal(t, hashutil.PjwHash("world"), uint32(119809))
}

func TestElfHash(t *testing.T) {
	assert.Equal(t, hashutil.ElfHash("test"), uint32(118224))
	assert.Equal(t, hashutil.ElfHash("hello"), uint32(117838))
	assert.Equal(t, hashutil.ElfHash("world"), uint32(119809))
}

func TestBkdrHash(t *testing.T) {
	assert.Equal(t, hashutil.BkdrHash("test"), uint32(229458829))
	assert.Equal(t, hashutil.BkdrHash("hello"), uint32(99162322))
	assert.Equal(t, hashutil.BkdrHash("world"), uint32(113318802))
}

func TestSdbmHash(t *testing.T) {
	assert.Equal(t, hashutil.SdbmHash("test"), uint32(2090756197))
	assert.Equal(t, hashutil.SdbmHash("hello"), uint32(99162322))
	assert.Equal(t, hashutil.SdbmHash("world"), uint32(113318802))
}

func TestDjbHash(t *testing.T) {
	assert.Equal(t, hashutil.DjbHash("test"), uint32(2090756197))
	assert.Equal(t, hashutil.DjbHash("hello"), uint32(99162322))
	assert.Equal(t, hashutil.DjbHash("world"), uint32(113318802))
}

func TestDekHash(t *testing.T) {
	assert.Equal(t, hashutil.DekHash("test"), uint32(4276487250))
	assert.Equal(t, hashutil.DekHash("hello"), uint32(2920040049))
	assert.Equal(t, hashutil.DekHash("world"), uint32(742451790))
}

func TestApHash(t *testing.T) {
	assert.Equal(t, hashutil.ApHash("test"), uint32(3422398098))
	assert.Equal(t, hashutil.ApHash("hello"), uint32(3255060260))
	assert.Equal(t, hashutil.ApHash("world"), uint32(2128592991))
}

func TestTianlHash(t *testing.T) {
	assert.Equal(t, hashutil.TianlHash("test"), uint32(1092158949))
	assert.Equal(t, hashutil.TianlHash("hello"), uint32(1084105767))
	assert.Equal(t, hashutil.TianlHash("world"), uint32(1258415024))
}

func TestJavaDefaultHash(t *testing.T) {
	assert.Equal(t, hashutil.JavaDefaultHash("test"), int32(3556498))
	assert.Equal(t, hashutil.JavaDefaultHash("hello"), int32(99162322))
	assert.Equal(t, hashutil.JavaDefaultHash("world"), int32(113318802))
}

func TestMixHash(t *testing.T) {
	assert.Equal(t, hashutil.MixHash("test"), uint64(12684967112958431589))
	assert.Equal(t, hashutil.MixHash("hello"), uint64(12684967112958431589))
	assert.Equal(t, hashutil.MixHash("world"), uint64(12684967112958431589))
}
