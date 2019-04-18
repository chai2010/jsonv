# jsonv: Get Value from JSON

## Example1 ([hello.go](hello.go))

```go
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
```

Output:

```
$ go run hello.go
abc
string: abc, <nil>
float64: 123, <nil>
float64: 22, <nil>
float64: 11, <nil>
float64: 11, <nil>
<nil>: <nil>, jsonv: not found
```

## Example2 ([hello2.go](hello2.go))

```go
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
```

Output:

```
$ go run hello2.go
port: 8080
db.user: root
db.password: 123456
```
