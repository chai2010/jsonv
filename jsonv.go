// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Get json value.
package jsonv

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type strError string

func (e strError) Error() string { return string(e) }

const ErrNotFound = strError("jsonv: not found")

func Get(data, key interface{}, subKeys ...interface{}) interface{} {
	v, _ := GetValue(data, key, subKeys...)
	return v
}

func GetValue(data, key interface{}, subKeys ...interface{}) (value interface{}, err error) {
	switch data := data.(type) {
	case string:
		if err := json.Unmarshal([]byte(data), &value); err != nil {
			return nil, err
		}
	case []byte:
		if err := json.Unmarshal([]byte(data), &value); err != nil {
			return nil, err
		}

	case map[string]interface{}:
		value = data
	case []interface{}:
		value = data

	default:
		if _type := reflect.TypeOf(data); _type.Kind() == reflect.String {
			// type xxx string
			if err := json.Unmarshal([]byte(reflect.ValueOf(data).String()), &value); err != nil {
				return nil, err
			}
		} else if _type.Kind() == reflect.Slice {
			if _type.Elem().Kind() == reflect.Uint8 {
				// type xxx []byte
				if err := json.Unmarshal([]byte(reflect.ValueOf(data).Bytes()), &value); err != nil {
					return nil, err
				}
			} else if _type.Elem().Kind() == reflect.Interface {
				// type xxx []interface{}
				value = data
			} else {
				return nil, fmt.Errorf("jsonv: unsupport data type: %T", data)
			}
		} else if _type.Kind() == reflect.Map {
			// type xxx map[string]interface{}
			if _type.Key().Kind() == reflect.String && _type.Elem().Kind() == reflect.Interface {
				value = data
			} else {
				return nil, fmt.Errorf("jsonv: unsupport data type: %T", data)
			}
		} else {
			return nil, fmt.Errorf("jsonv: unsupport data type: %T", data)
		}
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
				return nil, ErrNotFound
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
					return nil, ErrNotFound
				}
				iKey = i
			}

			if iKey < 0 || iKey >= len(x) {
				return nil, ErrNotFound
			}

			value = x[iKey]

		default:
			return nil, ErrNotFound
		}
	}

	// OK
	return value, nil
}
