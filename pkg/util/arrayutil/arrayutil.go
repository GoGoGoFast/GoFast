package arrayutil

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// IsEmpty 判断数组或切片是否为空
func IsEmpty(arr interface{}) (bool, error) {
	if !IsArray(arr) {
		return false, errors.New("provided data is not a slice or array")
	}
	v := reflect.ValueOf(arr)
	return v.Len() == 0, nil
}

// IsNotEmpty 判断数组或切片是否非空
func IsNotEmpty(arr interface{}) (bool, error) {
	empty, err := IsEmpty(arr)
	if err != nil {
		return false, err
	}
	return !empty, nil
}

// NewArray 新建泛型数组
func NewArray(elementType reflect.Type, size int) interface{} {
	return reflect.MakeSlice(reflect.SliceOf(elementType), size, size).Interface()
}

// Resize 调整切片大小
func Resize(arr interface{}, newSize int) (interface{}, error) {
	if !IsArray(arr) {
		return nil, errors.New("provided data is not a slice or array")
	}
	v := reflect.ValueOf(arr)
	newArr := reflect.MakeSlice(v.Type(), newSize, newSize)
	reflect.Copy(newArr, v)
	return newArr.Interface(), nil
}

// AddAll 合并多个切片
func AddAll(arrays ...interface{}) ([]interface{}, error) {
	var result []interface{}
	for _, arr := range arrays {
		if !IsArray(arr) {
			return nil, errors.New("all inputs must be slices or arrays")
		}
		v := reflect.ValueOf(arr)
		for i := 0; i < v.Len(); i++ {
			result = append(result, v.Index(i).Interface())
		}
	}
	return result, nil
}

// Clone 克隆切片
func Clone(arr interface{}) (interface{}, error) {
	if !IsArray(arr) {
		return nil, errors.New("provided data is not a slice or array")
	}
	v := reflect.ValueOf(arr)
	newArr := reflect.MakeSlice(v.Type(), v.Len(), v.Len())
	reflect.Copy(newArr, v)
	return newArr.Interface(), nil
}

// Range 生成步进有序数组
func Range(start, end, step int) []int {
	if step == 0 {
		return []int{}
	}
	size := (end-start)/step + 1
	if size <= 0 {
		return []int{}
	}
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = start + i*step
	}
	return result
}

// Split 拆分字节切片
func Split(data []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for len(data) > 0 {
		if len(data) < chunkSize {
			chunkSize = len(data)
		}
		chunks = append(chunks, data[:chunkSize])
		data = data[chunkSize:]
	}
	return chunks
}

// Filter 过滤切片元素
func Filter(arr interface{}, filterFunc func(interface{}) bool) ([]interface{}, error) {
	if !IsArray(arr) {
		return nil, errors.New("provided data is not a slice or array")
	}
	v := reflect.ValueOf(arr)
	var result []interface{}
	for i := 0; i < v.Len(); i++ {
		if filterFunc(v.Index(i).Interface()) {
			result = append(result, v.Index(i).Interface())
		}
	}
	return result, nil
}

// Edit 编辑切片中的元素
func Edit(arr interface{}, editFunc func(interface{}) interface{}) (interface{}, error) {
	if !IsArray(arr) {
		return nil, errors.New("provided data is not a slice or array")
	}
	v := reflect.ValueOf(arr)
	for i := 0; i < v.Len(); i++ {
		v.Index(i).Set(reflect.ValueOf(editFunc(v.Index(i).Interface())))
	}
	return arr, nil
}

// Zip 将两个切片结合成键值对
func Zip(keys, values interface{}) (map[interface{}]interface{}, error) {
	if !IsArray(keys) {
		return nil, errors.New("keys must be a slice or array")
	}
	if !IsArray(values) {
		return nil, errors.New("values must be a slice or array")
	}
	vKeys := reflect.ValueOf(keys)
	vValues := reflect.ValueOf(values)
	if vKeys.Len() != vValues.Len() {
		return nil, errors.New("keys and values must have the same length")
	}
	result := make(map[interface{}]interface{}, vKeys.Len())
	for i := 0; i < vKeys.Len(); i++ {
		result[vKeys.Index(i).Interface()] = vValues.Index(i).Interface()
	}
	return result, nil
}

// Contains 判断切片中是否包含元素
func Contains(arr interface{}, elem interface{}) (bool, error) {
	if !IsArray(arr) {
		return false, errors.New("provided data is not a slice or array")
	}
	v := reflect.ValueOf(arr)
	for i := 0; i < v.Len(); i++ {
		if reflect.DeepEqual(v.Index(i).Interface(), elem) {
			return true, nil
		}
	}
	return false, nil
}

// Wrap 将原始类型切片包装成泛型切片
func Wrap(arr interface{}) ([]interface{}, error) {
	if !IsArray(arr) {
		return nil, errors.New("provided data is not a slice or array")
	}
	v := reflect.ValueOf(arr)
	result := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = v.Index(i).Interface()
	}
	return result, nil
}

// Unwrap 将泛型切片拆包为原始类型切片
func Unwrap(arr []interface{}, elemType reflect.Type) (interface{}, error) {
	if elemType == nil {
		return nil, errors.New("element type cannot be nil")
	}
	v := reflect.MakeSlice(reflect.SliceOf(elemType), len(arr), len(arr))
	for i := 0; i < len(arr); i++ {
		v.Index(i).Set(reflect.ValueOf(arr[i]))
	}
	return v.Interface(), nil
}

// IsArray 判断是否为数组或切片
func IsArray(arr interface{}) bool {
	v := reflect.ValueOf(arr)
	return v.Kind() == reflect.Array || v.Kind() == reflect.Slice
}

// ToString 将数组或切片转为字符串
func ToString(arr interface{}, sep string) (string, error) {
	if !IsArray(arr) {
		return "", errors.New("provided data is not a slice or array")
	}
	v := reflect.ValueOf(arr)
	var result []string
	for i := 0; i < v.Len(); i++ {
		result = append(result, fmt.Sprintf("%v", v.Index(i).Interface())) // 使用 %v 处理不同类型
	}
	return strings.Join(result, sep), nil
}

// Print 打印数组或切片的内容
func Print(arr interface{}) error {
	if !IsArray(arr) {
		return errors.New("provided data is not a slice or array")
	}
	v := reflect.ValueOf(arr)
	for i := 0; i < v.Len(); i++ {
		fmt.Println(v.Index(i).Interface())
	}
	return nil
}
