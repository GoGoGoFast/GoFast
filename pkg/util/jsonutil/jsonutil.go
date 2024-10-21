package jsonutil

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Marshal marshals a Go value to a JSON byte slice.
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal unmarshals a JSON byte slice to a Go value.
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// MarshalIndent marshals a Go value to a pretty-printed JSON byte slice.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// DeepCopy performs a deep copy of a JSON-compatible Go value.
func DeepCopy(src, dest interface{}) error {
	data, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, dest)
}

// Merge merges two JSON-compatible Go values. The values must be maps or structs.
func Merge(dst, src interface{}) error {
	dstData, err := json.Marshal(dst)
	if err != nil {
		return err
	}

	srcData, err := json.Marshal(src)
	if err != nil {
		return err
	}

	var dstMap map[string]interface{}
	var srcMap map[string]interface{}

	if err := json.Unmarshal(dstData, &dstMap); err != nil {
		return err
	}

	if err := json.Unmarshal(srcData, &srcMap); err != nil {
		return err
	}

	for k, v := range srcMap {
		dstMap[k] = v
	}

	mergedData, err := json.Marshal(dstMap)
	if err != nil {
		return err
	}

	return json.Unmarshal(mergedData, dst)
}

// PrettyPrint prints a pretty-printed JSON string of a Go value.
func PrettyPrint(v interface{}) (string, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Query allows querying JSON data using a simple dot notation.
func Query(data []byte, query string) (interface{}, error) {
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	parts := splitQuery(query)
	for _, part := range parts {
		switch v := jsonData.(type) {
		case map[string]interface{}:
			var ok bool
			jsonData, ok = v[part]
			if !ok {
				return nil, fmt.Errorf("key %s not found", part)
			}
		case []interface{}:
			index, err := strconv.Atoi(part)
			if err != nil || index < 0 || index >= len(v) {
				return nil, fmt.Errorf("invalid array index %s", part)
			}
			jsonData = v[index]
		default:
			return nil, fmt.Errorf("unexpected type %T for key %s", jsonData, part)
		}
	}

	return jsonData, nil
}

func splitQuery(query string) []string {
	return strings.Split(query, ".")
}
