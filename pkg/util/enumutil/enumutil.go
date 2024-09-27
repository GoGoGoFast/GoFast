package enumutil

import (
	"reflect"
)

// GetNames 获取枚举类中所有枚举对象的名称列表
func GetNames(enumType interface{}) []string {
	enumValues := reflect.ValueOf(enumType)
	names := make([]string, enumValues.NumField())
	for i := 0; i < enumValues.NumField(); i++ {
		names[i] = enumValues.Type().Field(i).Name
	}
	return names
}

// GetFieldValues 获取枚举类中各枚举对象下指定字段的值
func GetFieldValues(enumType interface{}, fieldName string) []interface{} {
	enumValues := reflect.ValueOf(enumType)
	values := make([]interface{}, enumValues.NumField())
	for i := 0; i < enumValues.NumField(); i++ {
		field := enumValues.Field(i).FieldByName(fieldName)
		values[i] = field.Interface()
	}
	return values
}

// GetBy 根据传入条件获得对应枚举
func GetBy(enumType interface{}, condition func(string) bool) interface{} {
	enumValues := reflect.ValueOf(enumType)
	for i := 0; i < enumValues.NumField(); i++ {
		name := enumValues.Type().Field(i).Name
		if condition(name) {
			return enumValues.Field(i).Interface()
		}
	}
	return nil
}

// GetEnumMap 获取枚举名称与枚举对象的映射
func GetEnumMap(enumType interface{}) map[string]interface{} {
	enumValues := reflect.ValueOf(enumType)
	enumMap := make(map[string]interface{})
	for i := 0; i < enumValues.NumField(); i++ {
		name := enumValues.Type().Field(i).Name
		enumMap[name] = enumValues.Field(i).Interface()
	}
	return enumMap
}

// GetNameFieldMap 获取枚举名与指定字段值的映射
func GetNameFieldMap(enumType interface{}, fieldName string) map[string]interface{} {
	enumValues := reflect.ValueOf(enumType)
	fieldMap := make(map[string]interface{})
	for i := 0; i < enumValues.NumField(); i++ {
		name := enumValues.Type().Field(i).Name
		field := enumValues.Field(i).FieldByName(fieldName)
		fieldMap[name] = field.Interface()
	}
	return fieldMap
}
