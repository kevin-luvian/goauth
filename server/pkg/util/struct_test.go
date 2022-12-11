package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtil_ReadSruct(t *testing.T) {
	type AStruct2 struct {
		Name string `json:"name-2"`
		Keg  string `json:"keg"`
	}

	type AStruct struct {
		Name string   `json:"name"`
		AS2  AStruct2 `json:"a-struct2"`
	}

	var (
		err      error
		basePath = "consul"
		origin   = AStruct{
			Name: "abc",
		}
		want = AStruct{
			Name: "consul/name",
			AS2: AStruct2{
				Name: "consul/a-struct2/name-2",
				Keg:  "consul/a-struct2/keg",
			},
		}
	)

	data, err := StructToMap(origin)
	assert.NoError(t, err)

	m := FillIn(data, basePath, func(path string) interface{} {
		return path
	}).(map[string]interface{})

	err = MapToStruct(m, &origin)
	assert.NoError(t, err)

	assert.Equal(t, origin, want)
}
