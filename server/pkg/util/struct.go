package util

import (
	"encoding/json"
)

func MapToStruct(m map[string]interface{}, t interface{}) error {
	bs, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bs, t); err != nil {
		return err
	}

	return nil
}

func StructToMap(t interface{}) (map[string]interface{}, error) {
	bs, err := json.Marshal(t)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var m map[string]interface{}
	if err := json.Unmarshal(bs, &m); err != nil {
		return map[string]interface{}{}, err
	}

	return m, nil
}

func FillIn(i interface{}, path string, getVal func(path string) interface{}) interface{} {
	switch v := i.(type) {
	case map[string]interface{}:
		for key, val := range v {
			v[key] = FillIn(val, path+"/"+key, getVal)
		}
		return v
	default:
		return getVal(path)
	}
}
