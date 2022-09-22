
## Installation
```go

go get github.com/DigitalUnion/du-pco-go-sdk/dupco

```

## Quickstart

```go

import (
	"github.com/DigitalUnion/du-pco-go-sdk/dupco"
	"github.com/goccy/go-json"
	"log"
)

func ExampleClient() {
	api := dupco.NewClient("aaaaaa", "test", "aaaaaaaa")
	for i := 0; i < 10; i++ {
		r := api.Call("/geofence/v1/list_fence", nil)
		j, _ := json.Marshal(r)
		log.Println(string(j))
	}
}

```
