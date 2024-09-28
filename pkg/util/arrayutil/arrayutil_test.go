package arrayutil_test

import (
    "reflect"
    "testing"

   "VeloCore/pkg/util/arrayutil"
)

func TestIsEmpty(t *testing.T) {
    // 测试切片为空的情况
    arr := []int{}
    if!arrayutil.IsEmpty(arr) {
       t.Errorf("Expected empty slice to be recognized as empty")
    }

    // 测试数组为空的情况
    arr2 := [0]int{}
    if!arrayutil.IsEmpty(arr2) {
       t.Errorf("Expected empty array to be recognized as empty")
    }

    // 测试非空切片的情况
    arr3 := []int{1, 2, 3}
    if arrayutil.IsEmpty(arr3) {
       t.Errorf("Expected non-empty slice to not be recognized as empty")
    }
}

func TestIsNotEmpty(t *testing.T) {
    // 测试非空切片的情况
    arr := []int{1, 2, 3}
    if!arrayutil.IsNotEmpty(arr) {
       t.Errorf("Expected non-empty slice to be recognized as non-empty")
    }

    // 测试空切片的情况
    arr2 := []int{}
    if arrayutil.IsNotEmpty(arr2) {
       t.Errorf("Expected empty slice to not be recognized as non-empty")
    }
}

func TestNewArray(t *testing.T) {
    // 测试创建整数类型的切片
    newArr := arrayutil.NewArray(reflect.TypeOf(int(0)), 5)
    arr, ok := newArr.([]int)
    if!ok || len(arr)!= 5 {
       t.Errorf("Failed to create new int array")
    }
}

func TestResize(t *testing.T) {
    arr := []int{1, 2, 3}
    newArr, err := arrayutil.Resize(arr, 5)
    if err!= nil {
       t.Errorf("Unexpected error: %v", err)
    }
    resizedArr, ok := newArr.([]int)
    if!ok || len(resizedArr)!= 5 {
       t.Errorf("Failed to resize slice")
    }
}

func TestAddAll(t *testing.T) {
    arr1 := []int{1, 2, 3}
    arr2 := []int{4, 5, 6}
    combined, err := arrayutil.AddAll(arr1, arr2)
    if err!= nil {
       t.Errorf("Unexpected error: %v", err)
    }
    expected := []interface{}{1, 2, 3, 4, 5, 6}
    if!reflect.DeepEqual(combined, expected) {
       t.Errorf("Combined slices do not match expected result")
    }
}

func TestClone(t *testing.T) {
    arr := []int{1, 2, 3}
    clonedArr, err := arrayutil.Clone(arr)
    if err!= nil {
       t.Errorf("Unexpected error: %v", err)
    }
    cloned, ok := clonedArr.([]int)
    if!ok ||!reflect.DeepEqual(cloned, arr) {
       t.Errorf("Cloned slice is not correct")
    }
}

func TestRange(t *testing.T) {
    result := arrayutil.Range(1, 10, 2)
    expected := []int{1, 3, 5, 7, 9}
    if!reflect.DeepEqual(result, expected) {
       t.Errorf("Range function result does not match expected")
    }
}

func TestSplit(t *testing.T) {
    data := []byte("123456789")
    chunks := arrayutil.Split(data, 3)
    expected := [][]byte{[]byte("123"), []byte("456"), []byte("789")}
    if!reflect.DeepEqual(chunks, expected) {
       t.Errorf("Split function result does not match expected")
    }
}

func TestFilter(t *testing.T) {
    arr := []int{1, 2, 3, 4, 5}
    filtered := arrayutil.Filter(arr, func(i interface{}) bool {
       return i.(int)%2 == 0
    })
    expected := []interface{}{2, 4}
    if!reflect.DeepEqual(filtered, expected) {
       t.Errorf("Filter function result does not match expected")
    }
}

func TestEdit(t *testing.T) {
    arr := []int{1, 2, 3}
    edited := arrayutil.Edit(arr, func(i interface{}) interface{} {
       return i.(int) * 2
    })
    expected := []int{2, 4, 6}
    if!reflect.DeepEqual(edited, expected) {
       t.Errorf("Edit function result does not match expected")
    }
}

func TestZip(t *testing.T) {
    keys := []string{"a", "b", "c"}
    values := []int{1, 2, 3}
    zipped, err := arrayutil.Zip(keys, values)
    if err!= nil {
       t.Errorf("Unexpected error: %v", err)
    }
    expected := map[interface{}]interface{}{"a": 1, "b": 2, "c": 3}
    if!reflect.DeepEqual(zipped, expected) {
       t.Errorf("Zip function result does not match expected")
    }
}

func TestContains(t *testing.T) {
    arr := []int{1, 2, 3}
    if!arrayutil.Contains(arr, 2) {
       t.Errorf("Expected slice to contain element")
    }
    if arrayutil.Contains(arr, 4) {
       t.Errorf("Expected slice not to contain element")
    }
}

func TestWrap(t *testing.T) {
    arr := []int{1, 2, 3}
    wrapped := arrayutil.Wrap(arr)
    expected := []interface{}{1, 2, 3}
    if!reflect.DeepEqual(wrapped, expected) {
       t.Errorf("Wrap function result does not match expected")
    }
}

func TestUnwrap(t *testing.T) {
    wrapped := []interface{}{1, 2, 3}
    unwrapped := arrayutil.Unwrap(wrapped, reflect.TypeOf(int(0)))
    arr, ok := unwrapped.([]int)
    if!ok ||!reflect.DeepEqual(arr, []int{1, 2, 3}) {
       t.Errorf("Unwrap function result does not match expected")
    }
}

func TestIsArray(t *testing.T) {
    arr := []int{}
    if!arrayutil.IsArray(arr) {
       t.Errorf("Expected slice to be recognized as array/slice")
    }
    arr2 := [0]int{}
    if!arrayutil.IsArray(arr2) {
       t.Errorf("Expected array to be recognized as array/slice")
    }
}

func TestToString(t *testing.T) {
    arr := []int{1, 2, 3}
    result := arrayutil.ToString(arr, ",")
    expected := "1,2,3"
    if result!= expected {
       t.Errorf("ToString function result does not match expected")
    }
}