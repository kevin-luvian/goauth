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
