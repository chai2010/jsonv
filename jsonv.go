// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Get json value.
package jsonv

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func Get(data, key interface{}, subKeys ...interface{}) interface{} {
	v, _ := GetValue(data, key, subKeys...)
	return v
}

func GetValue(data, key interface{}, subKeys ...interface{}) (value interface{}, ok bool) {
	switch data := data.(type) {
	case string:
		if err := json.Unmarshal([]byte(data), &value); err != nil {
			return nil, false
		}
	case []byte:
		if err := json.Unmarshal([]byte(data), &value); err != nil {
			return nil, false
		}

	case map[string]interface{}:
		value = data
	case []interface{}:
		value = data

	default:
		return nil, false
	}

	allKeys := []interface{}{key}
	allKeys = append(allKeys, subKeys[:len(subKeys)]...)

	for _, curKey := range allKeys {
		switch x := value.(type) {
		case map[string]interface{}:
			var sKey string
			if s, ok := curKey.(string); ok {
				sKey = s
			} else {
				sKey = fmt.Sprint(curKey)
			}

			if subMap, _ := x[sKey]; subMap == nil {
				return nil, false
			} else {
				value = subMap
			}

		case []interface{}:
			var iKey int
			if i, ok := curKey.(int); ok {
				iKey = i
			} else {
				i, err := strconv.Atoi(fmt.Sprint(curKey))
				if err != nil {
					return nil, false
				}
				iKey = i
			}

			if iKey < 0 || iKey >= len(x) {
				return nil, false
			}

			value = x[iKey]

		default:
			return nil, false
		}
	}

	return value, true
}
