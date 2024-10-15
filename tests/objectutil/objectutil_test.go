package objectutil_test

import (
	"reflect"
	"testing"

	"VeloCore/pkg/util/objectutil"
)

// 定义测试结构体
type TestStruct struct {
	ID         int
	Name       string
	Tags       []string
	Attrs      map[string]interface{}
	ExtraField interface{}
}

// TestDeepClone 测试 DeepClone 函数
func TestDeepClone(t *testing.T) {
	original := &TestStruct{
		ID:    1,
		Name:  "Test",
		Tags:  []string{"tag1", "tag2"},
		Attrs: map[string]interface{}{"key1": "value1", "key2": 2},
	}
	clone, err := objectutil.DeepClone(original)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !objectutil.DeepCompare(original, clone) {
		t.Errorf("expected %v, got %v", original, clone)
	}
}

// TestIsEqual 测试 IsEqual 函数
func TestIsEqual(t *testing.T) {
	a := &TestStruct{ID: 1, Name: "A"}
	b := &TestStruct{ID: 1, Name: "A"}
	if !objectutil.IsEqual(a, b) {
		t.Errorf("expected objects to be equal")
	}
}

// TestMerge 测试 Merge 函数
func TestMerge(t *testing.T) {
	mapA := map[string]interface{}{"key1": "value1", "key2": "value2"}
	mapB := map[string]interface{}{"key2": "new_value2", "key3": "value3"}
	result := objectutil.Merge(mapA, mapB)
	expected := map[string]interface{}{
		"key1": "value1",
		"key2": "new_value2",
		"key3": "value3",
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

// TestGetField 测试 GetField 函数
func TestGetField(t *testing.T) {
	obj := &TestStruct{ID: 10}
	val, err := objectutil.GetField(obj, "ID")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if val != 10 {
		t.Errorf("expected %v, got %v", 10, val)
	}
}

// TestSetField 测试 SetField 函数
func TestSetField(t *testing.T) {
	obj := &TestStruct{}
	err := objectutil.SetField(obj, "ID", 42)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if obj.ID != 42 {
		t.Errorf("expected %v, got %v", 42, obj.ID)
	}
}

// TestSerializeDeserialize 测试序列化和反序列化
func TestSerializeDeserialize(t *testing.T) {
	obj := &TestStruct{ID: 1, Name: "Test"}
	data, err := objectutil.Serialize(obj)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var newObj TestStruct
	err = objectutil.Deserialize(data, &newObj)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !objectutil.IsEqual(obj, &newObj) {
		t.Errorf("expected %v, got %v", obj, newObj)
	}
}

// TestCompare 测试 Compare 函数
func TestCompare(t *testing.T) {
	obj1 := &TestStruct{ID: 1, Name: "Object 1"}
	obj2 := &TestStruct{ID: 1, Name: "Object 1"}
	equal, err := objectutil.Compare(obj1, obj2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !equal {
		t.Errorf("expected objects to be equal")
	}
}

// TestConvert 测试 Convert 函数
func TestConvert(t *testing.T) {
	type AnotherStruct struct {
		ID   int
		Name string
	}
	obj := &TestStruct{ID: 1, Name: "Test", ExtraField: "Ignored"} // 源结构体有额外字段
	converted, err := objectutil.Convert(obj, reflect.TypeOf(AnotherStruct{}))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	convertedObj := converted.(*AnotherStruct)
	if convertedObj.ID != 1 || convertedObj.Name != "Test" {
		t.Errorf("expected ID: 1 and Name: Test, got ID: %v and Name: %v", convertedObj.ID, convertedObj.Name)
	}
}

// TestSetFields 测试 SetFields 函数
func TestSetFields(t *testing.T) {
	obj := &TestStruct{}
	fields := map[string]interface{}{
		"ID":   100,
		"Name": "Updated Name",
	}
	err := objectutil.SetFields(obj, fields)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if obj.ID != 100 || obj.Name != "Updated Name" {
		t.Errorf("expected ID: 100, Name: Updated Name, got ID: %v, Name: %v", obj.ID, obj.Name)
	}
}

// TestToMap 测试 ToMap 函数
func TestToMap(t *testing.T) {
	obj := &TestStruct{
		ID:    1,
		Name:  "Test",
		Tags:  []string{"tag1", "tag2"},
		Attrs: map[string]interface{}{"key1": "value1"},
	}
	m, err := objectutil.ToMap(obj)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if m["ID"] != 1 || m["Name"] != "Test" {
		t.Errorf("expected ID: 1, Name: Test, got ID: %v, Name: %v", m["ID"], m["Name"])
	}
}

// TestFromMap 测试 FromMap 函数
func TestFromMap(t *testing.T) {
	data := map[string]interface{}{
		"ID":   2,
		"Name": "From Map",
	}
	obj := &TestStruct{}
	err := objectutil.FromMap(obj, data)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if obj.ID != 2 || obj.Name != "From Map" {
		t.Errorf("expected ID: 2, Name: From Map, got ID: %v, Name: %v", obj.ID, obj.Name)
	}
}
