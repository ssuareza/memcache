# About
This is a memcache client library for Go.

# Installing
```sh
$ go get github.com/ssuareza/memcache
```

# Example
```go
import (
    "github.com/ssuareza/memcache"
)

func main() {
    client, _ := memcache.New("localhost:11211")

    // set
    item := &memcache.Item{Key: "dog", Value: []byte("dago")}
	_ = client.Set(item)
	
    // get
    value, _ := client.Get("dog")
}
```