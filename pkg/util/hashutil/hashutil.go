package hashutil

import (
	"hash/crc32"
	"hash/fnv"
	"log"
)

// HashUtil 包提供了各种 hash 算法的实现。
// HashUtil package provides implementations of various hash algorithms.

// AdditiveHash 加法 hash 算法。
// AdditiveHash computes the additive hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func AdditiveHash(input string) uint32 {
	var hash uint32
	for i := 0; i < len(input); i++ {
		hash += uint32(input[i])
	}
	return hash
}

// RotatingHash 旋转 hash 算法。
// RotatingHash computes the rotating hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func RotatingHash(input string) uint32 {
	var hash uint32
	for i := 0; i < len(input); i++ {
		hash = (hash << 4) ^ (hash >> 28) ^ uint32(input[i])
	}
	return hash
}

// OneByOneHash 一次一个 hash 算法。
// OneByOneHash computes the one-by-one hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func OneByOneHash(input string) uint32 {
	var hash uint32
	for i := 0; i < len(input); i++ {
		hash += uint32(input[i])
		hash += (hash << 10)
		hash ^= (hash >> 6)
	}
	hash += (hash << 3)
	hash ^= (hash >> 11)
	hash += (hash << 15)
	return hash
}

// Bernstein bernstein Bernstein's hash 算法。
// Bernstein computes the Bernstein's hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func Bernstein(input string) uint32 {
	var hash uint32 = 5381
	for i := 0; i < len(input); i++ {
		hash = 33*hash + uint32(input[i])
	}
	return hash
}

// Universal universal Universal Hashing 算法。
// Universal computes the Universal Hashing hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func Universal(input string, prime uint32) uint32 {
	var hash uint32
	for i := 0; i < len(input); i++ {
		hash = (hash*prime + uint32(input[i])) % 4294967291
	}
	return hash
}

// Zobrist Hashing 算法。
// Zobrist computes the Zobrist Hashing hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func Zobrist(input string, table []uint32) uint32 {
	var hash uint32
	for i := 0; i < len(input); i++ {
		hash ^= table[input[i]]
	}
	return hash
}

// FnvHash 改进的32位 FNV 算法。
// FnvHash computes the 32-bit FNV-1 hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func FnvHash(input string) uint32 {
	h := fnv.New32a()
	_, err := h.Write([]byte(input))
	if err != nil {
		log.Fatalf("Failed to write input to hasher: %v", err)
	}
	return h.Sum32()
}

// IntHash Thomas Wang 的算法，整数 hash。
// IntHash computes the hash of an integer using Thomas Wang's algorithm.
// 参数 (param): key uint32 - 需要被哈希的整数。
// 参数 (param): key uint32 - the integer to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func IntHash(key uint32) uint32 {
	key = ^key + (key << 15)
	key = key ^ (key >> 12)
	key = key + (key << 2)
	key = key ^ (key >> 4)
	key = key * 2057
	key = key ^ (key >> 16)
	return key
}

// RsHash RS 算法 hash。
// RsHash computes the RS hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func RsHash(input string) uint32 {
	var b, a uint32 = 378551, 63689
	var hash uint32
	for i := 0; i < len(input); i++ {
		hash = hash*a + uint32(input[i])
		a *= b
	}
	return hash
}

// JsHash JS 算法。
// jsHash computes the JS hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func JsHash(input string) uint32 {
	var hash uint32
	for i := 0; i < len(input); i++ {
		hash ^= (hash << 5) + uint32(input[i]) + (hash >> 2)
	}
	return hash
}

// PjwHash PJW 算法。
// PjwHash computes the PJW hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func PjwHash(input string) uint32 {
	var threeQuarters, oneEighth, highBits uint32 = 24, 4, 0xF0000000
	var hash uint32
	for i := 0; i < len(input); i++ {
		hash = (hash << oneEighth) + uint32(input[i])
		if test := hash & highBits; test != 0 {
			hash = (hash ^ (test >> threeQuarters)) & (^highBits)
		}
	}
	return hash
}

// ElfHash ELF 算法。
// ElfHash computes the ELF hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func ElfHash(input string) uint32 {
	var hash, x uint32
	for i := 0; i < len(input); i++ {
		hash = (hash << 4) + uint32(input[i])
		if x = hash & 0xF0000000; x != 0 {
			hash ^= x >> 24
			hash &= ^x
		}
	}
	return hash
}

// BkdrHash BKDR 算法。
// BkdrHash computes the BKDR hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func BkdrHash(input string) uint32 {
	var seed, hash uint32 = 131, 0
	for i := 0; i < len(input); i++ {
		hash = hash*seed + uint32(input[i])
	}
	return hash
}

// SdbmHash SDBM 算法。
// SdbmHash computes the SDBM hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func SdbmHash(input string) uint32 {
	var hash uint32
	for i := 0; i < len(input); i++ {
		hash = uint32(input[i]) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}

// DjbHash DJB 算法。
// DjbHash computes the DJB hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func DjbHash(input string) uint32 {
	var hash uint32 = 5381
	for i := 0; i < len(input); i++ {
		hash = ((hash << 5) + hash) + uint32(input[i])
	}
	return hash
}

// DekHash DEK 算法。
// DekHash computes the DEK hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func DekHash(input string) uint32 {
	var hash uint32 = uint32(len(input))
	for i := 0; i < len(input); i++ {
		hash = ((hash << 5) ^ (hash >> 27)) ^ uint32(input[i])
	}
	return hash
}

// ApHash AP 算法。
// ApHash computes the AP hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func ApHash(input string) uint32 {
	var hash uint32
	for i := 0; i < len(input); i++ {
		if (i & 1) == 0 {
			hash ^= (hash << 7) ^ uint32(input[i]) ^ (hash >> 3)
		} else {
			hash ^= ^((hash << 11) ^ uint32(input[i]) ^ (hash >> 5))
		}
	}
	return hash
}

// TianlHash TianL Hash 算法。
// TianlHash computes the TianL hash of a given string.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint32 - 哈希值。
// 返回值 (return): uint32 - the hash value.
func TianlHash(input string) uint32 {
	var hash uint32
	for i := 0; i < len(input); i++ {
		hash += uint32(input[i])
		hash += (hash << 10)
		hash ^= (hash >> 6)
	}
	hash += (hash << 3)
	hash ^= (hash >> 11)
	hash += (hash << 15)
	return hash
}

// JavaDefaultHash JAVA 自带的 hash 算法。
// JavaDefaultHash computes the default hash algorithm used in Java.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): int32 - 哈希值。
// 返回值 (return): int32 - the hash value.
func JavaDefaultHash(input string) int32 {
	var hash int32
	for i := 0; i < len(input); i++ {
		hash = 31*hash + int32(input[i])
	}
	return hash
}

// MixHash 混合 hash 算法，输出 64 位的值。
// MixHash computes a mixed hash of a given string, outputting a 64-bit value.
// 参数 (param): input string - 需要被哈希的字符串。
// 参数 (param): input string - the string to be hashed.
// 返回值 (return): uint64 - 哈希值。
// 返回值 (return): uint64 - the hash value.
func MixHash(input string) uint64 {
	h1 := fnv.New32a()
	_, err := h1.Write([]byte(input))
	if err != nil {
		log.Fatalf("Failed to write input to FNV hasher: %v", err)
	}
	h2 := crc32.ChecksumIEEE([]byte(input))
	return uint64(h1.Sum32())<<32 | uint64(h2)
}
