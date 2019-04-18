// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package jsonv_test

import (
	"testing"

	"github.com/chai2010/jsonv"
)

func TestGet(t *testing.T) {
	v0 := jsonv.Get(`{"value":"abc"}`, "value")
	tAssert(t, v0 != nil)
	tAssert(t, v0.(string) == "abc")

	v1 := jsonv.Get(`{"value":123}`, "value")
	tAssert(t, v1 != nil)
	tAssert(t, v1.(float64) == 123)

	v2 := jsonv.Get(`["abc"]`, 0)
	tAssert(t, v2 != nil)
	tAssert(t, v2.(string) == "abc")

	v3 := jsonv.Get(`["abc", 123]`, 1)
	tAssert(t, v3 != nil)
	tAssert(t, v3.(float64) == 123)

	v4 := jsonv.Get(`[["abc", 123]]`, 0, 1)
	tAssert(t, v4 != nil)
	tAssert(t, v4.(float64) == 123)

	v5 := jsonv.Get(`[{"value":"abc"}]`, 0, "value")
	tAssert(t, v5 != nil)
	tAssert(t, v5.(string) == "abc")

	v6 := jsonv.Get(`[{"value":["abc",123]}]`, 0, "value", 1)
	tAssert(t, v6 != nil)
	tAssert(t, v6.(float64) == 123)
}

func TestGet_notFound(t *testing.T) {
	v0 := jsonv.Get(`{}`, "aaa")
	tAssert(t, v0 == nil)

	v1 := jsonv.Get(`[]`, 0)
	tAssert(t, v1 == nil)

	v2 := jsonv.Get(`{"value":"abc"}`, "aaa")
	tAssert(t, v2 == nil)

	v3 := jsonv.Get(`[123]`, 100)
	tAssert(t, v3 == nil)
}

func TestGetValue_notFound(t *testing.T) {
	v0, err1 := jsonv.GetValue(`{}`, "aaa")
	tAssert(t, v0 == nil)
	tAssert(t, err1 == jsonv.ErrNotFound)

	v1, err1 := jsonv.GetValue(`[]`, 0)
	tAssert(t, v1 == nil)
	tAssert(t, err1 == jsonv.ErrNotFound)

	v2, err2 := jsonv.GetValue(`{"value":"abc"}`, "aaa")
	tAssert(t, v2 == nil)
	tAssert(t, err2 == jsonv.ErrNotFound)

	v3, err3 := jsonv.GetValue(`[123]`, 100)
	tAssert(t, v3 == nil)
	tAssert(t, err3 == jsonv.ErrNotFound)
}

func TestGetValue_failed(t *testing.T) {
	var err error

	_, err = jsonv.GetValue(``, "aaa")
	tAssert(t, err != nil)
	tAssert(t, err != jsonv.ErrNotFound)

	_, err = jsonv.GetValue(`{,}`, "aaa")
	tAssert(t, err != nil)
	tAssert(t, err != jsonv.ErrNotFound)
}

func tAssert(tb testing.TB, ok bool, a ...interface{}) {
	if !ok {
		tb.Helper()
		tb.Fatal(a...)
	}
}
