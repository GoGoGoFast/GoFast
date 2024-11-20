package arrayutil_test

import (
	"GoFast/pkg/util/arrayutil"
	"reflect"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	if empty, err := arrayutil.IsEmpty([]int{}); err != nil || !empty {
		t.Errorf("IsEmpty([]int{}) = %v, %v; want true, nil", empty, err)
	}

	if empty, err := arrayutil.IsEmpty([]int{1}); err != nil || empty {
		t.Errorf("IsEmpty([]int{1}) = %v, %v; want false, nil", empty, err)
	}

	if _, err := arrayutil.IsEmpty("not an array"); err == nil {
		t.Errorf("IsEmpty(not an array) returned no error; want an error")
	}
}

func TestIsNotEmpty(t *testing.T) {
	if notEmpty, err := arrayutil.IsNotEmpty([]int{1}); err != nil || !notEmpty {
		t.Errorf("IsNotEmpty([]int{1}) = %v, %v; want true, nil", notEmpty, err)
	}

	if notEmpty, err := arrayutil.IsNotEmpty([]int{}); err != nil || notEmpty {
		t.Errorf("IsNotEmpty([]int{}) = %v, %v; want false, nil", notEmpty, err)
	}

	if _, err := arrayutil.IsNotEmpty("not an array"); err == nil {
		t.Errorf("IsNotEmpty(not an array) returned no error; want an error")
	}
}

func TestNewArray(t *testing.T) {
	arr := arrayutil.NewArray(reflect.TypeOf(0), 5)
	if v := reflect.ValueOf(arr); v.Kind() != reflect.Slice || v.Len() != 5 {
		t.Errorf("NewArray(int, 5) = %v; want slice of length 5", arr)
	}
}

func TestResize(t *testing.T) {
	resized, err := arrayutil.Resize([]int{1, 2, 3}, 5)
	if err != nil {
		t.Errorf("Resize([]int{1, 2, 3}, 5) returned error: %v", err)
	}
	if v := reflect.ValueOf(resized); v.Kind() != reflect.Slice || v.Len() != 5 {
		t.Errorf("Resize([]int{1, 2, 3}, 5) = %v; want slice of length 5", resized)
	}

	if _, err := arrayutil.Resize("not an array", 5); err == nil {
		t.Errorf("Resize(not an array, 5) returned no error; want an error")
	}
}

func TestAddAll(t *testing.T) {
	result, err := arrayutil.AddAll([]int{1, 2}, []int{3, 4})
	if err != nil || len(result) != 4 {
		t.Errorf("AddAll([]int{1, 2}, []int{3, 4}) = %v, %v; want length 4, nil", result, err)
	}

	if _, err := arrayutil.AddAll([]int{1}, "not an array"); err == nil {
		t.Errorf("AddAll([]int{1}, not an array) returned no error; want an error")
	}
}

func TestClone(t *testing.T) {
	cloned, err := arrayutil.Clone([]int{1, 2, 3})
	if err != nil {
		t.Errorf("Clone([]int{1, 2, 3}) returned error: %v", err)
	}
	if !reflect.DeepEqual(cloned, []int{1, 2, 3}) {
		t.Errorf("Clone([]int{1, 2, 3}) = %v; want original slice", cloned)
	}

	if _, err := arrayutil.Clone("not an array"); err == nil {
		t.Errorf("Clone(not an array) returned no error; want an error")
	}
}

func TestFilter(t *testing.T) {
	filtered, err := arrayutil.Filter([]int{1, 2, 3}, func(v interface{}) bool {
		return v.(int)%2 == 0
	})
	if err != nil || len(filtered) != 1 || filtered[0] != 2 {
		t.Errorf("Filter([]int{1, 2, 3}) returned %v, %v; want [2], nil", filtered, err)
	}

	if _, err := arrayutil.Filter("not an array", nil); err == nil {
		t.Errorf("Filter(not an array, nil) returned no error; want an error")
	}
}

func TestEdit(t *testing.T) {
	edited, err := arrayutil.Edit([]int{1, 2, 3}, func(v interface{}) interface{} {
		return v.(int) * 2
	})
	if err != nil {
		t.Errorf("Edit([]int{1, 2, 3}, ...) returned error: %v", err)
	}
	if !reflect.DeepEqual(edited, []int{2, 4, 6}) {
		t.Errorf("Edit([]int{1, 2, 3}, ...) = %v; want [2, 4, 6]", edited)
	}

	if _, err := arrayutil.Edit("not an array", nil); err == nil {
		t.Errorf("Edit(not an array, nil) returned no error; want an error")
	}
}

func TestZip(t *testing.T) {
	result, err := arrayutil.Zip([]string{"a", "b"}, []int{1, 2})
	if err != nil {
		t.Errorf("Zip([]string{\"a\", \"b\"}, []int{1, 2}) returned error: %v", err)
	}
	expected := map[interface{}]interface{}{"a": 1, "b": 2}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Zip([]string{\"a\", \"b\"}, []int{1, 2}) = %v; want %v", result, expected)
	}

	if _, err := arrayutil.Zip([]int{1}, "not an array"); err == nil {
		t.Errorf("Zip([]int{1}, not an array) returned no error; want an error")
	}
}

func TestContains(t *testing.T) {
	contains, err := arrayutil.Contains([]int{1, 2, 3}, 2)
	if err != nil || !contains {
		t.Errorf("Contains([]int{1, 2, 3}, 2) = %v, %v; want true, nil", contains, err)
	}

	contains, err = arrayutil.Contains([]int{1, 2, 3}, 4)
	if err != nil || contains {
		t.Errorf("Contains([]int{1, 2, 3}, 4) = %v, %v; want false, nil", contains, err)
	}

	if _, err := arrayutil.Contains("not an array", 1); err == nil {
		t.Errorf("Contains(not an array, 1) returned no error; want an error")
	}
}

func TestWrap(t *testing.T) {
	wrapped, err := arrayutil.Wrap([]int{1, 2, 3})
	if err != nil || len(wrapped) != 3 {
		t.Errorf("Wrap([]int{1, 2, 3}) = %v, %v; want length 3, nil", wrapped, err)
	}

	if _, err := arrayutil.Wrap("not an array"); err == nil {
		t.Errorf("Wrap(not an array) returned no error; want an error")
	}
}

func TestUnwrap(t *testing.T) {
	unwrap, err := arrayutil.Unwrap([]interface{}{1, 2, 3}, reflect.TypeOf(0))
	if err != nil || !reflect.DeepEqual(unwrap, []int{1, 2, 3}) {
		t.Errorf("Unwrap([]interface{}{1, 2, 3}, int) = %v, %v; want []int{1, 2, 3}, nil", unwrap, err)
	}

	if _, err := arrayutil.Unwrap([]interface{}{1, 2, 3}, nil); err == nil {
		t.Errorf("Unwrap([], nil) returned no error; want an error")
	}
}

func TestIsArray(t *testing.T) {
	if !arrayutil.IsArray([]int{1, 2}) {
		t.Errorf("IsArray([]int{1, 2}) = false; want true")
	}

	if arrayutil.IsArray("not an array") {
		t.Errorf("IsArray(not an array) = true; want false")
	}
}

func TestToString(t *testing.T) {
	result, err := arrayutil.ToString([]int{1, 2, 3}, ", ")
	if err != nil || result != "1, 2, 3" {
		t.Errorf("ToString([]int{1, 2, 3}, \", \") = %v, %v; want \"1, 2, 3\", nil", result, err)
	}

	if _, err := arrayutil.ToString("not an array", ", "); err == nil {
		t.Errorf("ToString(not an array, \", \") returned no error; want an error")
	}
}
