# operator for rtctunnel

## usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/nobonobo/operator"
)

func main() {
	op := operator.New()

	// Publish data
	err := op.Pub("test", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to data
	data, err := op.Sub("test")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
```
