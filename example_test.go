// Copyright 2019 chaishushan{AT}gmail.com. All rights reserved.
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

	v1, err1 := jsonv.GetValue(`{"value":"abc"}`, "value")
	fmt.Printf("%[1]T: %[1]v, %v\n", v1, err1)

	v2, err2 := jsonv.GetValue(`{"value":{"abc":123}}`, "value", "abc")
	fmt.Printf("%[1]T: %[1]v, %v\n", v2, err2)

	v3, err3 := jsonv.GetValue(`{"value":{"abc":[11,22]}}`, "value", "abc", 1)
	fmt.Printf("%[1]T: %[1]v, %v\n", v3, err3)

	// failed
	v4, err4 := jsonv.GetValue(`{"value":{"abc":[11,22]}}`, "aaa")
	fmt.Printf("%[1]T: %[1]v, %v\n", v4, err4)

	// failed
	v5, err5 := jsonv.GetValue(`[{"value":{"abc":[11,22]}}]`, 0, "value", "abc", 10)
	fmt.Printf("%[1]T: %[1]v, %v\n", v5, err5)
	// err5 == jsonv.ErrNotFound

	// Output:
	// abc
	// string: abc, <nil>
	// float64: 123, <nil>
	// float64: 22, <nil>
	// <nil>: <nil>, jsonv: not found
	// <nil>: <nil>, jsonv: not found
}
