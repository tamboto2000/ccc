# CCC

[![Go Reference](https://pkg.go.dev/badge/github.com/tamboto2000/ccc.svg)](https://pkg.go.dev/github.com/tamboto2000/ccc)

Card BIN Checker.
When you know, you know ;)

### Install

```sh
$ GO111MODULE=on go get github.com/tamboto2000/ccc
```

### Example

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/tamboto2000/ccc"
)

func main() {
	card, err := ccc.Check("537941")
	if err != nil {
		panic(err.Error())
	}

	js, err := json.Marshal(card)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(js))
}
```

### Example With Proxy

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/tamboto2000/ccc"
)

func main() {
	card, err := ccc.CheckWithProx("537941", "http://185.198.188.55:8080")
	if err != nil {
		panic(err.Error())
	}

	js, err := json.Marshal(card)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(js))
}
```

License
-------

MIT