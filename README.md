# go-etcdcacher

golang in process cache for etcd using go-etcd

## Install

```bash
go get github.com/coreos/go-etcd/etcd
go get github.com/jvdberg/go-etcdcacher/etcdcacher
```

## Examples

```go
package main

import (
	"github.com/coreos/go-etcd/etcd"
	"github.com/jvdberg/go-etcdcacher/etcdcacher"
	"time"
	"fmt"
)

func main() {
	c := etcd.NewClient([]string{"http://127.0.0.1:4001"})
	c.Set("/test/testkey","test",0)

	cache := etcdcacher.NewCache(c,"/test/")

	values, _ := cache.Get("/test/testkey")
    for i, res := range values {
        fmt.Printf("[%d] get response: %+v\n", i, res)
    }

	c.Set("/test/testkey","test2",0)
	time.Sleep(1* time.Second)

	values,_ = cache.Get("/test/testkey")
    for i, res := range values {
        fmt.Printf("[%d] get response: %+v\n", i, res)
    }
}
```