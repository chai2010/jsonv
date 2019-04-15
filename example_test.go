// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package jsonv_test

import (
	"fmt"

	"github.com/chai2010/jsonv"
)

func Example() {
	v0 := jsonv.Get(`{"value":"abc"}`, "value")
	fmt.Printf("%v\n", v0)

	v1, ok1 := jsonv.GetValue(`{"value":"abc"}`, "value")
	fmt.Printf("%[1]T: %[1]v, %v\n", v1, ok1)

	v2, ok2 := jsonv.GetValue(`{"value":{"abc":123}}`, "value", "abc")
	fmt.Printf("%[1]T: %[1]v, %v\n", v2, ok2)

	v3, ok3 := jsonv.GetValue(`{"value":{"abc":[11,22]}}`, "value", "abc", 1)
	fmt.Printf("%[1]T: %[1]v, %v\n", v3, ok3)

	// failed
	v4, ok4 := jsonv.GetValue(`{"value":{"abc":[11,22]}}`, "aaa")
	fmt.Printf("%[1]T: %[1]v, %v\n", v4, ok4)

	// failed
	v5, ok5 := jsonv.GetValue(`[{"value":{"abc":[11,22]}}]`, 0, "value", "abc", 10)
	fmt.Printf("%[1]T: %[1]v, %v\n", v5, ok5)

	// Output:
	// abc
	// string: abc, true
	// float64: 123, true
	// float64: 22, true
	// <nil>: <nil>, false
	// <nil>: <nil>, false
}
