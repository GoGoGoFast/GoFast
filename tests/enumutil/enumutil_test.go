package enumutil_test

import (
	"reflect"
	"testing"

	"GoAllInOne/pkg/util/enumutil"
)

// 定义一个测试用的枚举类型
type Color struct {
	Red   ColorEnum
	Green ColorEnum
	Blue  ColorEnum
}

type ColorEnum struct {
	Name  string
	Value int
}

// 模拟一个枚举对象
var Colors = Color{
	Red:   ColorEnum{Name: "Red", Value: 1},
	Green: ColorEnum{Name: "Green", Value: 2},
	Blue:  ColorEnum{Name: "Blue", Value: 3},
}

// 测试GetNames函数
func TestGetNames(t *testing.T) {
	expected := []string{"Red", "Green", "Blue"}
	names := enumutil.GetNames(Colors)

	if !reflect.DeepEqual(names, expected) {
		t.Errorf("GetNames() = %v; want %v", names, expected)
	}
}

// 测试GetFieldValues函数
func TestGetFieldValues(t *testing.T) {
	expected := []interface{}{1, 2, 3}
	values := enumutil.GetFieldValues(Colors, "Value")

	if !reflect.DeepEqual(values, expected) {
		t.Errorf("GetFieldValues() = %v; want %v", values, expected)
	}
}

// 测试GetBy函数
func TestGetBy(t *testing.T) {
	expected := Colors.Green
	result := enumutil.GetBy(Colors, func(name string) bool {
		return name == "Green"
	})

	if result != expected {
		t.Errorf("GetBy() = %v; want %v", result, expected)
	}
}

// 测试GetEnumMap函数
func TestGetEnumMap(t *testing.T) {
	expected := map[string]interface{}{
		"Red":   Colors.Red,
		"Green": Colors.Green,
		"Blue":  Colors.Blue,
	}
	enumMap := enumutil.GetEnumMap(Colors)

	if !reflect.DeepEqual(enumMap, expected) {
		t.Errorf("GetEnumMap() = %v; want %v", enumMap, expected)
	}
}

// 测试GetNameFieldMap函数
func TestGetNameFieldMap(t *testing.T) {
	expected := map[string]interface{}{
		"Red":   1,
		"Green": 2,
		"Blue":  3,
	}
	fieldMap := enumutil.GetNameFieldMap(Colors, "Value")

	if !reflect.DeepEqual(fieldMap, expected) {
		t.Errorf("GetNameFieldMap() = %v; want %v", fieldMap, expected)
	}
}
