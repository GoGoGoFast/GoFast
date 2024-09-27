package idutil

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

// RandomUUID 生成带 "-" 的 UUID (版本4)
func RandomUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// UUID version 4 (random)
	b[6] = (b[6] & 0x0f) | 0x40 // 设置版本为 4
	b[8] = (b[8] & 0x3f) | 0x80 // 设置 variant 为 RFC4122

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// SimpleUUID 生成不带 "-" 的 UUID
func SimpleUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	// UUID version 4 (random)
	b[6] = (b[6] & 0x0f) | 0x40 // 设置版本为 4
	b[8] = (b[8] & 0x3f) | 0x80 // 设置 variant 为 RFC4122

	return hex.EncodeToString(b)
}

// ObjectId 生成 MongoDB 风格的 ObjectId
func ObjectId() string {
	b := make([]byte, 12)
	_, err := rand.Read(b[4:])
	if err != nil {
		panic(err)
	}

	// 时间戳部分，前4个字节是当前秒数
	timestamp := uint32(time.Now().Unix())
	b[0] = byte(timestamp >> 24)
	b[1] = byte(timestamp >> 16)
	b[2] = byte(timestamp >> 8)
	b[3] = byte(timestamp)

	return hex.EncodeToString(b)
}

// Snowflake 结构体
type Snowflake struct {
	sync.Mutex
	timestamp int64
	sequence  int64
	nodeID    int64
}

// NewSnowflake 创建一个新的 Snowflake 生成器
func NewSnowflake(nodeID int64) *Snowflake {
	return &Snowflake{
		timestamp: 0,
		sequence:  0,
		nodeID:    nodeID,
	}
}

// NextId 生成下一个 Snowflake ID
func (s *Snowflake) NextId() int64 {
	s.Lock()
	defer s.Unlock()

	const twepoch = int64(1288834974657) // Twitter的初始纪元时间
	const nodeIDBits = uint(10)
	const sequenceBits = uint(12)

	nodeIDShift := sequenceBits
	timestampShift := sequenceBits + nodeIDBits
	sequenceMask := int64(-1 ^ (-1 << sequenceBits))

	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = now

	return ((now - twepoch) << timestampShift) |
		(s.nodeID << nodeIDShift) |
		s.sequence
}

// NextIdStr 生成下一个 Snowflake ID 字符串
func (s *Snowflake) NextIdStr() string {
	return fmt.Sprintf("%d", s.NextId())
}
