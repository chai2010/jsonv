// Copyright 2019 chaishushan{AT}gmail.com. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"strings"

	"github.com/chai2010/jsonv"
)

func main() {
	var cfg Config = `{
		"port": 8080,
		"db": { "user": "root", "password": "123456" }
	}`

	fmt.Println("port:", cfg.Get("port"))
	fmt.Println("db.user:", cfg.Get("db.user"))
	fmt.Println("db.password:", cfg.Get("db.password"))

}

type Config string

func (cfg Config) Get(path string) string {
	var keys []interface{}
	for _, s := range strings.Split(path, ".") {
		keys = append(keys, s)
	}
	if len(keys) < 1 {
		return ""
	}

	v := jsonv.Get(string(cfg), keys[0], keys[1:]...)
	if v == nil {
		return ""
	}

	return fmt.Sprint(v)
}
