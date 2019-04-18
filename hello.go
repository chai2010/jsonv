// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"

	"github.com/chai2010/jsonv"
)

func main() {
	v0 := jsonv.Get(`{"value":"abc"}`, "value")
	fmt.Printf("%v\n", v0)

	v1, err1 := jsonv.GetValue(`{"value":"abc"}`, "value")
	fmt.Printf("%[1]T: %[1]v, %v\n", v1, err1)
	assert(err1 == nil)

	v2, err2 := jsonv.GetValue(`{"value":{"abc":123}}`, "value", "abc")
	fmt.Printf("%[1]T: %[1]v, %v\n", v2, err2)
	assert(err2 == nil)

	v3, err3 := jsonv.GetValue(`{"value":{"abc":[11,22]}}`, "value", "abc", 1)
	fmt.Printf("%[1]T: %[1]v, %v\n", v3, err3)
	assert(err3 == nil)

	v4, err4 := jsonv.GetValue(`[{"value":{"abc":[11,22]}}]`, 0, "value", "abc", 0)
	fmt.Printf("%[1]T: %[1]v, %v\n", v4, err4)
	assert(err4 == nil)

	v5, err5 := jsonv.GetValue(`[{"value":{"abc":[11,22]}}]`, 0, "value", "abc", 0)
	fmt.Printf("%[1]T: %[1]v, %v\n", v5, err5)
	assert(err5 == nil)

	// not found
	v6, err6 := jsonv.GetValue(`[{"value":{"abc":[11,22]}}]`, 0, "value", "abc", 100)
	fmt.Printf("%[1]T: %[1]v, %v\n", v6, err6)
	assert(err6 == jsonv.ErrNotFound)
}

func assert(ok bool, a ...interface{}) {
	if !ok {
		panic(fmt.Sprint(a...))
	}
}
