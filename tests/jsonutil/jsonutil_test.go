package jsonutil

import (
	"VeloCore/pkg/util/jsonutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalUnmarshal(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "Alice", Age: 30}
	data, err := jsonutil.Marshal(p)
	assert.NoError(t, err)

	var p2 Person
	err = jsonutil.Unmarshal(data, &p2)
	assert.NoError(t, err)
	assert.Equal(t, p, p2)
}

func TestDeepCopy(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p1 := Person{Name: "Alice", Age: 30}
	var p2 Person

	err := jsonutil.DeepCopy(&p1, &p2)
	assert.NoError(t, err)
	assert.Equal(t, p1, p2)
}

func TestMerge(t *testing.T) {
	type Person struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Gender string `json:"gender"`
	}

	p1 := Person{Name: "Alice", Age: 30}
	p2 := Person{Gender: "female"}

	err := jsonutil.Merge(&p1, &p2)
	assert.NoError(t, err)
	assert.Equal(t, "female", p1.Gender)
}

func TestPrettyPrint(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	p := Person{Name: "Alice", Age: 30}
	output, err := jsonutil.PrettyPrint(p)
	assert.NoError(t, err)

	expected := `{
  "name": "Alice",
  "age": 30
}`
	assert.Equal(t, expected, output)
}

func TestQuery(t *testing.T) {
	data := []byte(`{
        "person": {
            "name": "Alice",
            "age": 30,
            "contacts": [
                {"type": "email", "value": "alice@example.com"},
                {"type": "phone", "value": "123-456-7890"}
            ]
        }
    }`)

	result, err := jsonutil.Query(data, "person.name")
	assert.NoError(t, err)
	assert.Equal(t, "Alice", result)

	result, err = jsonutil.Query(data, "person.contacts.0.type")
	assert.NoError(t, err)
	assert.Equal(t, "email", result)

	result, err = jsonutil.Query(data, "person.contacts.1.value")
	assert.NoError(t, err)
	assert.Equal(t, "123-456-7890", result)
}
