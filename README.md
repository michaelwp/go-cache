# go-cache
simple go library to manage number of data in cache

installation:
```
go get -d github.com/michaelwp/go-cache
```

example on how to use it:
```go
package main

import (
	"fmt"
	gocache "github.com/michaelwp/go-cache"
)

func main() {
	capacity := int32(3)
	cache := gocache.NewCache(capacity)

	// add new data to cache
	cache.Add("A", "testA")
	cache.Add("B", "testB")
	cache.Add("C", "testC")
	cache.Add("D", "testD")
	cache.Add("E", "testE")

	// get data from cache
	valA, isAExist := cache.Get("A")
	if !isAExist {
		fmt.Println("A key doesn't exist")
	} else {
		fmt.Println("A:", valA)
	}

	valD, isDExist := cache.Get("D")
	if !isDExist {
		fmt.Println("D key doesn't exist")
	} else {
		fmt.Println("D:", valD)
	}

	cache.Add("F", "testF")
	cache.Add("G", "testG")

	fmt.Println("cache:", cache.View())
}

```
