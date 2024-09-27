package objectutil

import (
	"encoding/json"
	"errors"
	"reflect"
)

// BaseObject 定义基础对象接口
// BaseObject defines a base object interface.
type BaseObject interface {
	GetID() interface{} // 获取对象的ID
}

// DeepClone 执行深度复制
// DeepClone performs a deep copy of the provided object.
// 参数:
//
//	src: 需要被复制的对象，必须是非空指针。
//
// 返回值:
//
//	返回复制后的对象和可能的错误信息。
func DeepClone(src interface{}) (interface{}, error) {
	if src == nil {
		return nil, nil
	}
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() != reflect.Ptr || srcVal.IsNil() {
		return nil, errors.New("source must be a non-nil pointer")
	}

	switch srcVal.Elem().Kind() {
	case reflect.Struct:
		dstVal := reflect.New(srcVal.Elem().Type())
		dstVal.Elem().Set(srcVal.Elem())
		return dstVal.Interface(), nil
	case reflect.Slice:
		len := srcVal.Elem().Len()
		dstSlice := reflect.MakeSlice(srcVal.Elem().Type(), len, len)
		for i := 0; i < len; i++ {
			clone, err := DeepClone(srcVal.Elem().Index(i).Interface())
			if err != nil {
				return nil, err
			}
			dstSlice.Index(i).Set(reflect.ValueOf(clone))
		}
		return dstSlice.Interface(), nil
	case reflect.Map:
		dstMap := reflect.MakeMap(srcVal.Elem().Type())
		iter := srcVal.Elem().MapRange()
		for iter.Next() {
			keyClone, err := DeepClone(iter.Key().Interface())
			if err != nil {
				return nil, err
			}
			valueClone, err := DeepClone(iter.Value().Interface())
			if err != nil {
				return nil, err
			}
			dstMap.SetMapIndex(reflect.ValueOf(keyClone), reflect.ValueOf(valueClone))
		}
		return dstMap.Interface(), nil
	default:
		return srcVal.Elem().Interface(), nil
	}
}

// IsEqual 检查两个对象是否相等
// IsEqual checks if two objects are equal.
// 参数:
//
//	a: 第一个对象。
//	b: 第二个对象。
//
// 返回值:
//
//	如果相等返回 true，否则返回 false。
func IsEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// Merge 合并两个映射
// Merge merges two maps of string keys and interface values.
// 参数:
//
//	a: 第一个映射。
//	b: 第二个映射。
//
// 返回值:
//
//	返回合并后的映射。
func Merge(a, b map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] = v
	}
	return result
}

// GetField 获取结构体字段的值
// GetField retrieves the value of a field from a struct.
// 参数:
//
//	obj: 目标对象，必须是非空指针。
//	field: 字段名称。
//
// 返回值:
//
//	返回字段的值和可能的错误信息。
func GetField(obj interface{}, field string) (interface{}, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return nil, errors.New("object must be a non-nil pointer")
	}

	fieldVal := val.Elem().FieldByName(field)
	if !fieldVal.IsValid() {
		return nil, errors.New("field not found")
	}
	return fieldVal.Interface(), nil
}

// SetField 设置结构体字段的值
// SetField sets the value of a field in a struct.
// 参数:
//
//	obj: 目标对象，必须是非空指针。
//	field: 字段名称。
//	value: 要设置的值。
//
// 返回值:
//
//	可能的错误信息。
func SetField(obj interface{}, field string, value interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.New("object must be a non-nil pointer")
	}

	fieldVal := val.Elem().FieldByName(field)
	if !fieldVal.IsValid() || !fieldVal.CanSet() {
		return errors.New("field not found or cannot be set")
	}
	fieldVal.Set(reflect.ValueOf(value))
	return nil
}

// TypeOf 返回对象的类型
// TypeOf returns the type of an object.
// 参数:
//
//	obj: 目标对象。
//
// 返回值:
//
//	对象的类型字符串。
func TypeOf(obj interface{}) string {
	return reflect.TypeOf(obj).String()
}

// IsNil 检查对象是否为 nil
// IsNil checks if an interface is nil.
// 参数:
//
//	obj: 目标对象。
//
// 返回值:
//
//	如果为 nil 返回 true，否则返回 false。
func IsNil(obj interface{}) bool {
	if obj == nil {
		return true
	}
	val := reflect.ValueOf(obj)
	return val.Kind() == reflect.Ptr && val.IsNil()
}

// Serialize 将对象转换为 JSON
// Serialize converts an object to JSON.
// 参数:
//
//	obj: 需要序列化的对象。
//
// 返回值:
//
//	JSON 字节数组和可能的错误信息。
func Serialize(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

// Deserialize 将 JSON 数据转换为对象
// Deserialize converts JSON data to an object.
// 参数:
//
//	data: JSON 字节数组。
//	obj: 目标对象，必须是非空指针。
//
// 返回值:
//
//	可能的错误信息。
func Deserialize(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}

// Compare 比较两个结构体的字段
// Compare compares two structs field by field.
// 参数:
//
//	a: 第一个结构体，必须是指针。
//	b: 第二个结构体，必须是指针。
//
// 返回值:
//
//	如果相等返回 true，否则返回 false，可能的错误信息。
func Compare(a, b interface{}) (bool, error) {
	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b)

	if valA.Kind() != reflect.Ptr || valB.Kind() != reflect.Ptr {
		return false, errors.New("both arguments must be pointers to structs")
	}

	if valA.Elem().Type() != valB.Elem().Type() {
		return false, errors.New("struct types must match")
	}

	for i := 0; i < valA.Elem().NumField(); i++ {
		fieldA := valA.Elem().Field(i)
		fieldB := valB.Elem().Field(i)

		if !reflect.DeepEqual(fieldA.Interface(), fieldB.Interface()) {
			return false, nil
		}
	}
	return true, nil
}

// DeepCompare 执行深度比较
// DeepCompare performs a deep comparison of two objects.
// 参数:
//
//	a: 第一个对象。
//	b: 第二个对象。
//
// 返回值:
//
//	如果相等返回 true，否则返回 false。
func DeepCompare(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// Validate 验证对象是否实现了 BaseObject 接口
// Validate checks if the object satisfies the BaseObject interface.
// 参数:
//
//	obj: 目标对象。
//
// 返回值:
//
//	可能的错误信息。
func Validate(obj interface{}) error {
	if bo, ok := obj.(BaseObject); ok {
		if bo.GetID() == nil {
			return errors.New("object ID cannot be nil")
		}
		return nil
	}
	return errors.New("object does not implement BaseObject")
}

// SetFields 设置结构体多个字段的值
// SetFields sets the values of multiple fields in a struct.
// 参数:
//
//	obj: 目标对象，必须是非空指针。
//	fields: 包含字段名和对应值的映射。
//
// 返回值:
//
//	可能的错误信息。
func SetFields(obj interface{}, fields map[string]interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.New("object must be a non-nil pointer")
	}

	structVal := val.Elem()
	structType := structVal.Type()

	for fieldName, value := range fields {
		field, found := structType.FieldByName(fieldName)
		if !found {
			return errors.New("field not found: " + fieldName)
		}

		fieldVal := structVal.FieldByName(fieldName)
		if !fieldVal.IsValid() || !fieldVal.CanSet() {
			return errors.New("field not found or cannot be set: " + fieldName)
		}

		_ = field.Type

		fieldVal.Set(reflect.ValueOf(value))
	}

	return nil
}

// FilterFields 过滤结构体的字段
// FilterFields filters the fields of a struct based on a given condition function.
// 参数:
//
//	obj: 目标对象，必须是非空指针。
//	condition: 条件函数，接受字段名和字段值，返回是否满足条件。
//
// 返回值:
//
//	返回满足条件的字段名和对应值的映射。
func FilterFields(obj interface{}, condition func(fieldName string, fieldValue interface{}) bool) map[string]interface{} {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return nil
	}

	structVal := val.Elem()
	structType := structVal.Type()

	result := make(map[string]interface{})
	for i := 0; i < structVal.NumField(); i++ {
		fieldName := structType.Field(i).Name
		fieldValue := structVal.Field(i).Interface()
		if condition(fieldName, fieldValue) {
			result[fieldName] = fieldValue
		}
	}

	return result
}

// Convert 转换对象类型
// Convert converts an object of one type to another.
// 参数:
//
//	src: 源对象。
//	dstType: 目标对象类型。
//
// 返回值:
//
//	返回转换后的对象和可能的错误信息。
func Convert(src interface{}, dstType reflect.Type) (interface{}, error) {
	srcVal := reflect.ValueOf(src)
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	dstVal := reflect.New(dstType).Elem()
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dstVal.Field(i)
		if dstField.CanSet() && srcField.IsValid() {
			dstField.Set(srcField)
		}
	}

	return dstVal.Addr().Interface(), nil
}

// Clone 直接克隆对象
// Clone creates a direct clone of the provided object.
// 参数:
//
//	src: 需要克隆的对象，必须是非空指针。
//
// 返回值:
//
//	返回克隆后的对象和可能的错误信息。
func Clone(src interface{}) (interface{}, error) {
	return DeepClone(src)
}

// MergeAndOverride 合并两个映射，后者覆盖前者
// MergeAndOverride merges two maps, with the second map overriding the first in case of key conflicts.
// 参数:
//
//	a: 第一个映射。
//	b: 第二个映射。
//
// 返回值:
//
//	返回合并后的映射。
func MergeAndOverride(a, b map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range a {
		result[k] = v
	}
	for k, v := range b {
		result[k] = v
	}
	return result
}

// GetFieldNames 获取结构体字段名称
// GetFieldNames retrieves the field names of a struct.
// 参数:
//
//	obj: 目标对象，必须是非空指针。
//
// 返回值:
//
//	返回字段名称的切片和可能的错误信息。
func GetFieldNames(obj interface{}) ([]string, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return nil, errors.New("object must be a non-nil pointer")
	}

	var fieldNames []string
	structVal := val.Elem()
	structType := structVal.Type()

	for i := 0; i < structVal.NumField(); i++ {
		fieldNames = append(fieldNames, structType.Field(i).Name)
	}

	return fieldNames, nil
}

// ToMap 将结构体转换为映射
// ToMap converts a struct to a map.
// 参数:
//
//	obj: 目标对象，必须是非空指针。
//
// 返回值:
//
//	返回映射和可能的错误信息。
func ToMap(obj interface{}) (map[string]interface{}, error) {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return nil, errors.New("object must be a non-nil pointer")
	}

	result := make(map[string]interface{})
	structVal := val.Elem()
	structType := structVal.Type()

	for i := 0; i < structVal.NumField(); i++ {
		fieldName := structType.Field(i).Name
		fieldValue := structVal.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result, nil
}

// FromMap 从映射中填充结构体
// FromMap populates a struct from a map.
// 参数:
//
//	obj: 目标对象，必须是非空指针。
//	data: 包含字段名和对应值的映射。
//
// 返回值:
//
//	可能的错误信息。
func FromMap(obj interface{}, data map[string]interface{}) error {
	val := reflect.ValueOf(obj)
	if val.Kind() != reflect.Ptr || val.IsNil() {
		return errors.New("object must be a non-nil pointer")
	}

	structVal := val.Elem()
	structType := structVal.Type()

	for fieldName, value := range data {
		fieldVal := structVal.FieldByName(fieldName)
		if !fieldVal.IsValid() || !fieldVal.CanSet() {
			return errors.New("field not found or cannot be set: " + fieldName)
		}
		fieldVal.Set(reflect.ValueOf(value))
	}
	_ = structType.Field

	return nil
}
