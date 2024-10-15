package idcardUtil_test_test

import (
	"VeloCore/pkg/util/idutil"
	"testing"
)

func TestRandomUUID(t *testing.T) {
	uuid := idutil.RandomUUID()
	if len(uuid) != 36 {
		t.Errorf("RandomUUID() length = %d; want 36", len(uuid))
	}
	if uuid[8] != '-' || uuid[13] != '-' || uuid[18] != '-' || uuid[23] != '-' {
		t.Errorf("RandomUUID() = %s; want correctly formatted UUID with dashes", uuid)
	}
}

func TestSimpleUUID(t *testing.T) {
	uuid := idutil.SimpleUUID()
	if len(uuid) != 32 {
		t.Errorf("SimpleUUID() length = %d; want 32", len(uuid))
	}
}

func TestObjectId(t *testing.T) {
	objectId := idutil.ObjectId()
	if len(objectId) != 24 {
		t.Errorf("ObjectId() length = %d; want 24", len(objectId))
	}
}

func TestSnowflake(t *testing.T) {
	nodeID := int64(1)
	sf := idutil.NewSnowflake(nodeID)

	id1 := sf.NextId()
	id2 := sf.NextId()

	if id1 == id2 {
		t.Errorf("Snowflake.NextId() generated duplicate IDs: %d and %d", id1, id2)
	}

	idStr1 := sf.NextIdStr()
	idStr2 := sf.NextIdStr()

	if idStr1 == idStr2 {
		t.Errorf("Snowflake.NextIdStr() generated duplicate IDs: %s and %s", idStr1, idStr2)
	}

	// 检查 Snowflake ID 是否单调递增
	if id1 >= id2 {
		t.Errorf("Snowflake.NextId() = %d should be less than %d", id1, id2)
	}
}

func TestSnowflakeUniquePerNode(t *testing.T) {
	// 测试不同节点生成的 Snowflake ID 是否唯一
	node1 := idutil.NewSnowflake(1)
	node2 := idutil.NewSnowflake(2)

	id1 := node1.NextId()
	id2 := node2.NextId()

	if id1 == id2 {
		t.Errorf("Snowflake generated duplicate IDs across different nodes: %d and %d", id1, id2)
	}
}

func BenchmarkRandomUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idutil.RandomUUID()
	}
}

func BenchmarkSimpleUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idutil.SimpleUUID()
	}
}

func BenchmarkObjectId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		idutil.ObjectId()
	}
}

func BenchmarkSnowflake(b *testing.B) {
	sf := idutil.NewSnowflake(1)
	for i := 0; i < b.N; i++ {
		sf.NextId()
	}
}
